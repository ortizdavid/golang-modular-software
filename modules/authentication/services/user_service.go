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
)

type UserService struct {
	repository *repositories.UserRepository
	roleRepository *repositories.RoleRepository
	userRoleRepository *repositories.UserRoleRepository
	loginActRepository *repositories.LoginActivityRepository
}

func NewUserService(db *database.Database) *UserService {
	return &UserService{
		repository: repositories.NewUserRepository(db),
		roleRepository: repositories.NewRoleRepository(db),
		userRoleRepository: repositories.NewUserRoleRepository(db),
		loginActRepository: repositories.NewLoginActivityRepository(db),
	}
}

func (s *UserService) CreateUser(ctx context.Context, request entities.CreateUserRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByName(ctx, request.UserName)
	if err != nil {
		return apperrors.NewNotFoundError(err.Error())
	}
	if exists {
		return apperrors.NewConflictError("user "+request.UserName+" already exists")
	}
	user := entities.User{
		UserId:    0,
		UserName:  request.UserName,
		Email:     request.Email,
		Password:  encryption.HashPassword(request.Password),
		IsActive:  true,
		UserImage:     "",
		Token:     encryption.GenerateRandomToken(),
		UniqueId:  encryption.GenerateUUID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	err = s.repository.Create(ctx, user)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating user: "+ err.Error())
	}
	userId := s.repository.LastInsertId
	_, err = s.roleRepository.FindById(ctx, request.RoleId)
	if err != nil {
		return apperrors.NewNotFoundError("role not found")
	}
	userRole := entities.UserRole{
		UserId:     userId,
		RoleId:     request.RoleId,
		UniqueId:  encryption.GenerateUUID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	err = s.userRoleRepository.Create(ctx, userRole)
	if err != nil {
		return apperrors.NewInternalServerError("error while adding role: "+ err.Error())
	}
	loginAct := entities.LoginActivity{
		UserId:   userId,
		Status: entities.ActivityStatusOffline,
		UniqueId:  encryption.GenerateUUID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	err = s.loginActRepository.Create(ctx, loginAct)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) EditUser(ctx context.Context, userId int64, request entities.EditUserRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found")
	}
	user.UserName = request.UserName
	err = s.repository.Update(ctx, user)
	if err != nil {
		return apperrors.NewInternalServerError("error while assign role: "+ err.Error())
	}
	return nil
}

func (s *UserService) AssignUserRole(ctx context.Context, userId int64, request entities.AssignUserRoleRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found")
	}
	role, err := s.roleRepository.FindById(ctx, request.RoleId)
	if err != nil {
		return apperrors.NewNotFoundError("role not found")
	}
	exists, err := s.userRoleRepository.ExistsByUserId(ctx, userId, request.RoleId)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewConflictError(fmt.Sprintf("role '%s' already assigned to user '%s'", role.RoleName, user.UserName))
	}
	userRole := entities.UserRole{
		UserId:     userId,
		RoleId:     request.RoleId,
		UniqueId:  encryption.GenerateUUID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	err = s.userRoleRepository.Create(ctx, userRole)
	if err != nil {
		return apperrors.NewInternalServerError("error while assigning role: "+ err.Error())
	}
	return nil
}

func (s *UserService) ChangeUserPassword(ctx context.Context, userId int64, request entities.ChangePasswordRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found")
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
		return apperrors.NewNotFoundError("user not found")
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
		return apperrors.NewNotFoundError("user not found")
	}
	// remove current image if exists //TODO
	uploadPath := config.UploadImagePath()
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
		return apperrors.NewNotFoundError("user not found")
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
		return apperrors.NewNotFoundError("user not found")
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
		return apperrors.NewNotFoundError("user not found")
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
	users, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
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
		return nil, apperrors.NewNotFoundError("No users found")
	}
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.FindAllByStatus(ctx, true, params.Limit, params.CurrentPage)
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
		return nil, apperrors.NewNotFoundError("No users found")
	}
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.FindAllByStatus(ctx, false, params.Limit, params.CurrentPage)
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
		return entities.UserData{}, apperrors.NewNotFoundError("user not found")
	}
	return user, nil
}

func (s *UserService) GetUserByUniqueId(ctx context.Context, uniqueId string) (entities.UserData, error) {
	user, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found")
	}
	return user, nil
}

func (s *UserService) GetUserByToken(ctx context.Context, token string) (entities.User, error) {
	user, err := s.repository.FindByToken(ctx, token)
	if err != nil {
		return entities.User{}, apperrors.NewNotFoundError("user not found")
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

func (s *UserService) UserHasRoles(ctx context.Context, userId int64, comparedRoles ...string) (bool, error) {
	for _, user := range comparedRoles {
		exists, err := s.roleRepository.ExistsByCode(ctx, user)
		if err != nil {
			return false, fmt.Errorf("error validating user '%s': %w", user, err)
		}
		if !exists {
			return false, fmt.Errorf("user '%s' does not exist", user)
		}
	}
	hasRole, err := s.repository.HasRoles(ctx, userId, comparedRoles...)
	if err != nil {
		return false, fmt.Errorf("error checking users for user ID %d: %w", userId, err)
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

