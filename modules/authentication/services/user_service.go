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
		//--- Create User ----------------------------------------------------------------------
		user := entities.User{
			UserName:  request.UserName,
			Email:     request.Email,
			Password:  encryption.HashPassword(request.Password),
			IsActive:  true,
			UserImage:     "",
			Token:     encryption.GenerateRandomToken(),
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
		_, err = s.roleRepository.FindById(ctx, request.RoleId)
		if err != nil {
			return apperrors.NewNotFoundError("role not found. invalid role id")
		}
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

func (s *UserService) AssignUserRole(ctx context.Context, userId int64, request entities.AssignUserRoleRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. inavlid user id")
	}
	role, err := s.roleRepository.FindById(ctx, request.RoleId)
	if err != nil {
		return apperrors.NewNotFoundError("role not found. invalid role id")
	}
	exists, err := s.userRoleRepository.ExistsByUserAndRole(ctx, userId, request.RoleId)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewConflictError(fmt.Sprintf("role '%s' already assigned to user '%s'", role.RoleName, user.UserName))
	}
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
		return apperrors.NewInternalServerError("error while assigning role: "+ err.Error())
	}
	return nil
}

func (s *UserService) AssociateUserToRole(ctx context.Context, request entities.AssociateUserRequest) error{
	user, err := s.repository.FindById(ctx, request.UserId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. inavlid user id")
	}
	exists, err := s.userAssociationRepository.ExistsByUserId(ctx, user.UserId)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewConflictError(fmt.Sprintf("User '%s' already associated ", user.UserName))
	}
	userAssociation := entities.UserAssociation{
		UserId:     user.UserId,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err = s.userAssociationRepository.Create(ctx, userAssociation)
	if err != nil {
		return apperrors.NewInternalServerError("error while associating user to a role: "+ err.Error())
	}
	return nil
}

func (s *UserService) RemoveUserRole(ctx context.Context, uniqueId string) error {
	userRole, err := s.userRoleRepository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("user role not found. invalid unique id")
	}
	err = s.userRoleRepository.Delete(ctx, userRole)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing role: "+ err.Error())
	}
	return nil
}

func (s *UserService) ChangeUserPassword(ctx context.Context, userId int64, request entities.ChangePasswordRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. invalid id")
	}
	user.Password = encryption.HashPassword(request.NewPassword)
	user.Token = encryption.GenerateRandomToken()
	user.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, user)
	if err != nil {
		return apperrors.NewInternalServerError("error while changing password: "+ err.Error())
	}
	return nil
}

func (s *UserService) ResetUserPassword(ctx context.Context, userId int64, request entities.ResetPasswordRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. invalid id")
	}
	user.Password = encryption.HashPassword(request.NewPassword)
	user.Token = encryption.GenerateRandomToken()
	user.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, user)
	if err != nil {
		return apperrors.NewInternalServerError("error while reseting password: "+ err.Error())
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
	//currentImage := user.UserImage
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

func (s *UserService) ActivateUser(ctx context.Context, userId int64) error {
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. invalid user id")
	}
	user.IsActive = true
	user.Token = encryption.GenerateRandomToken()
	user.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, user)
	if err != nil {
		return apperrors.NewInternalServerError("error while activating user: "+ err.Error())
	}
	return nil
}

func (s *UserService) DeactivateUser(ctx context.Context, userId int64) error {
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. inavlid user id")
	}
	user.IsActive = false
	user.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, user)
	if err != nil {
		return apperrors.NewInternalServerError("error while deactivating user: "+ err.Error())
	}
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, userId int64) error {
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. invalid user id")
	}
	err = s.repository.Delete(ctx, user)
	if err != nil {
		return apperrors.NewInternalServerError("error while deleting user: "+ err.Error())
	}
	return nil
}

func (s *UserService) GetAllUsers(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.UserData], error) {
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No users found")
	}
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.FindAllDataLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, users, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *UserService) GetAllActiveUsers(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.UserData], error) {
	count, err := s.CountUsersByStatus(ctx, true)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No active users found")
	}
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.FindAllByStatus(ctx, "Yes", params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, users, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *UserService) GetAllInactiveUsers(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.UserData], error) {
	count, err := s.CountUsersByStatus(ctx, false)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No inactive users found")
	}
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.FindAllByStatus(ctx, "No", params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, users, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *UserService) GetAllOnlineUsers(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.UserData], error) {
	count, err := s.CountUsersByActivityStatus(ctx, entities.ActivityStatusOnline)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No online users found")
	}
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.FindAllByActivityStatus(ctx, entities.ActivityStatusOnline, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, users, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *UserService) GetAllOfflineUsers(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.UserData], error) {
	count, err := s.CountUsersByActivityStatus(ctx, entities.ActivityStatusOffline)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No offline users found")
	}
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.FindAllByActivityStatus(ctx, entities.ActivityStatusOffline, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, users, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *UserService) SearchUsers(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchUserRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.UserData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No users found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, users, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *UserService) GetUserById(ctx context.Context, userId int64) (entities.UserData, error) {
	user, err := s.repository.GetDataById(ctx, userId)
	if err != nil {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found. invalid id")
	}
	return user, nil
}

func (s *UserService) GetUserByUniqueId(ctx context.Context, uniqueId string) (entities.UserData, error) {
	user, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found. invalid unique id")
	}
	return user, nil
}

func (s *UserService) GetUserApiKey(ctx context.Context, xUserId string) (entities.UserApiKey, error) {
	user, err := s.userApiKeyRepository.FindByXUserId(ctx, xUserId)
	if err != nil {
		return entities.UserApiKey{}, apperrors.NewNotFoundError("user api key not found. invalid x-user-id")
	}
	return user, nil
}

func (s *UserService) GetUserApiKeyById(ctx context.Context, userId int64) (entities.UserApiKey, error) {
	user, err := s.userApiKeyRepository.FindByUserId(ctx, userId)
	if err != nil {
		return entities.UserApiKey{}, apperrors.NewNotFoundError("user api key not found")
	}
	return user, nil
}

func (s *UserService) GetUserByToken(ctx context.Context, token string) (entities.User, error) {
	user, err := s.repository.FindByToken(ctx, token)
	if err != nil {
		return entities.User{}, apperrors.NewNotFoundError("user not found. invalid token")
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, token string) (entities.User, error) {
	user, err := s.repository.FindByEmail(ctx, token)
	if err != nil {
		return entities.User{}, apperrors.NewNotFoundError("user not found")
	}
	return user, nil
}

func (s *UserService) GetUsersWithoutAssociation(ctx context.Context, roles []string) ([]entities.User, error) {
	users, err := s.repository.FindUsersWithoutAssociation(ctx, roles)
	if err != nil {
		return nil, apperrors.NewNotFoundError("no users found")
	}
	return users, nil
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

func (s *UserService) CountUsers(ctx context.Context) (int64, error) {
	count, err := s.repository.Count(ctx)
	if err != nil {
		return 0, apperrors.NewNotFoundError("No users found")
	}
	return count, nil
}

func (s *UserService) CountUsersByStatus(ctx context.Context, status bool) (int64, error) {
	count, err := s.repository.CountByStatus(ctx, status)
	if err != nil {
		return 0, apperrors.NewNotFoundError("No users found")
	}
	return count, nil
}

func (s *UserService) CountUsersByActivityStatus(ctx context.Context, status entities.LoginActivityStatus) (int64, error) {
	count, err := s.repository.CountByActivityStatus(ctx, status)
	if err != nil {
		return 0, apperrors.NewNotFoundError("No users found")
	}
	return count, nil
}

func (s *UserService) GetUserInsertedId() int64 {
	return s.userInsertedId
}