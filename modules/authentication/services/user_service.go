package services

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	entities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type UserService struct {
	repository *repositories.UserRepository
	roleRepository *repositories.RoleRepository
	userRoleRepository *repositories.UserRoleRepository
	userAssociationRepository *repositories.UserAssociationRepository
	loginActRepository *repositories.LoginActivityRepository
	userApiKeyRepository *repositories.UserApiKeyRepository
	userInsertedId	int64
}

func NewUserService(db *database.Database) *UserService {
	return &UserService{
		repository:                repositories.NewUserRepository(db),
		roleRepository:            repositories.NewRoleRepository(db),
		userRoleRepository:        repositories.NewUserRoleRepository(db),
		userAssociationRepository: repositories.NewUserAssociationRepository(db),
		loginActRepository:        repositories.NewLoginActivityRepository(db),
		userApiKeyRepository:      repositories.NewUserApiKeyRepository(db),
	}
}

func (s *UserService) CreateUser(ctx context.Context, request entities.CreateUserRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	// Apply a transaction to ensure all operations are atomic (either all succeed or all fail)
	return s.repository.WithTransaction(ctx, func(tx *database.Database) error {
		// --- Check if user with the given name or email already exists
		existsName, err := s.repository.ExistsByName(ctx, request.UserName)
		if err != nil {
			return apperrors.NewNotFoundError(err.Error())
		}
		if existsName {
			return apperrors.NewConflictError("user '"+request.UserName+"' already exists")
		}
		existsEmail, err := s.repository.ExistsByEmail(ctx, request.Email)
		if err != nil {
			return apperrors.NewNotFoundError(err.Error())
		}
		if existsEmail {
			return apperrors.NewConflictError("email '"+request.Email+"' already exists")
		}
		// Check Role ------------------
		role, err := s.roleRepository.FindById(ctx, request.RoleId)
		if err != nil {
			return apperrors.NewNotFoundError("role not found. invalid role id")
		}
		//--- Create User ----------------------------------------------------------------------
		user := entities.User{
			UserName:  request.UserName,
			Email:     request.Email,
			Password:  encryption.HashPassword(request.Password),
			IsActive:  true,
			UserImage: "",
			Token:     encryption.GenerateRandomToken(),
			InitialRole: role.Code,
			BaseEntity: shared.BaseEntity{
				UniqueId:  encryption.GenerateUUID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		err = s.repository.Create(ctx, user)
		if err != nil {
			return apperrors.NewInternalServerError("error while creating user: "+ err.Error())
		}
		userId := s.repository.LastInsertId
		s.userInsertedId = userId
		// --- Create a new UserRole association
		userRole := entities.UserRole{
			UserId:     userId,
			RoleId:     request.RoleId,
			BaseEntity: shared.BaseEntity{
				UniqueId:  encryption.GenerateUUID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		err = s.userRoleRepository.Create(ctx, userRole)
		if err != nil {
			return apperrors.NewInternalServerError("error while adding role: "+ err.Error())
		}
		// --- Create Login Activity entry
		loginAct := entities.LoginActivity{
			UserId:   userId,
			Status: entities.ActivityStatusOffline,
			BaseEntity: shared.BaseEntity{
				UniqueId:  encryption.GenerateUUID(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			},
		}
		err = s.loginActRepository.Create(ctx, loginAct)
		if err != nil {
			return apperrors.NewInternalServerError("error while creating login activity: "+err.Error())
		}
		// --- Generate and assign an API key for the user
		userApiKey := entities.UserApiKey{
			UserId:    userId,
			XUserId:   encryption.GenerateUUID(),
			XApiKey:   encryption.GenerateRandomToken(),
			IsActive:  true,
			ExpiresAt: time.Now().AddDate(0, 1, 0),  // Add 1 month to the current time
			BaseEntity: shared.BaseEntity{
				UniqueId:  encryption.GenerateUUID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		err = s.userApiKeyRepository.Create(ctx, userApiKey)
		if err != nil {
			return apperrors.NewInternalServerError("error while creating api key: "+err.Error())
		}
		return nil
	})
}

func (s *UserService) UpdateUser(ctx context.Context, userId int64, request entities.UpdateUserRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. invalid user id")
	}
	user.UserName = request.UserName
	err = s.repository.Update(ctx, user)
	if err != nil {
		return apperrors.NewInternalServerError("error while update user: "+ err.Error())
	}
	return nil
}


func (s *UserService) UploadUserImage(ctx context.Context, fiberCtx *fiber.Ctx, userId int64) error {
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. invalid user id")
	}
	// remove current image if exists //TODO
	uploadPath := config.UploadImagePath() + "/users"
	uploader := helpers.NewUploader(uploadPath, config.MaxUploadImageSize(), helpers.ExtImages)
	info, err := uploader.UploadSingleFile(fiberCtx, "user_image")
	if err != nil {
		return apperrors.NewNotFoundError("error while uploading image: "+err.Error())
	}
	user.UserImage = info.FinalName
	user.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, user)
	if err != nil {
		return apperrors.NewInternalServerError("error while changing password: "+ err.Error())
	}
	return nil
}

func (s *UserService) RemoveUser(ctx context.Context, uniqueId string) error {
	err := s.repository.DeleteByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewInternalServerError("error while deleting user: "+ err.Error())
	}
	return nil
}

func (s *UserService) UserHasRoles(ctx context.Context, userId int64, comparedRoles ...string) (bool, error) {
	for _, role := range comparedRoles {
		exists, err := s.roleRepository.ExistsByCode(ctx, role)
		if err != nil {
			return false, fmt.Errorf("error validating role '%s': %w", role, err)
		}
		if !exists {
			return false, fmt.Errorf("role '%s' does not exist", role)
		}
	}
	hasRole, err := s.repository.HasRoles(ctx, userId, comparedRoles...)
	if err != nil {
		return false, fmt.Errorf("error checking roles for user ID %d: %w", userId, err)
	}
	return hasRole, nil
}

func (s *UserService) GetUserInsertedId() int64 {
	return s.userInsertedId
}