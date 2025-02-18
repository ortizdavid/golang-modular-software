package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type PermissionService struct {
	repository               *repositories.PermissionRepository
	permissionRoleRepository *repositories.PermissionRoleRepository
}

func NewPermissionService(db *database.Database) *PermissionService {
	return &PermissionService{
		repository:               repositories.NewPermissionRepository(db),
		permissionRoleRepository: repositories.NewPermissionRoleRepository(db),
	}
}

func (s *PermissionService) CreatePermission(ctx context.Context, request entities.CreatePermissionRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	permission := entities.Permission{
		PermissionName: request.PermissionName,
		Code:           request.Code,
		Description:    request.Description,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
	err := s.repository.Create(ctx, permission)
	if err != nil {
		return apperrors.InternalServerError("error while creating permission: " + err.Error())
	}
	return nil
}

func (s *PermissionService) UpdatePermission(ctx context.Context, uniqueId string, request entities.UpdatePermissionRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	permission, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NotFoundError("permission not found")
	}
	exists, err := s.permissionRoleRepository.ExistsByPermissionId(ctx, permission.PermissionId)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.ConflictError("Permission '" + permission.PermissionName + "' is currently assigned to users and cannot be updated")
	}
	permission.PermissionName = request.PermissionName
	permission.Description = request.Description
	permission.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, permission)
	if err != nil {
		return apperrors.InternalServerError("error while updating permission: " + err.Error())
	}
	return nil
}

func (s *PermissionService) RemovePermission(ctx context.Context, permissionId int) error {
	permission, err := s.repository.FindById(ctx, permissionId)
	if err != nil {
		return apperrors.NotFoundError("permission not found")
	}
	exists, err := s.permissionRoleRepository.ExistsByPermissionId(ctx, permissionId)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.ConflictError("Permission '" + permission.PermissionName + "' is currently assigned to roles and cannot be deleted")
	}
	err = s.repository.Delete(ctx, permission)
	if err != nil {
		return apperrors.InternalServerError("error while deleting permission: " + err.Error())
	}
	return nil
}

func (s *PermissionService) GetAllPermissionsPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.PermissionData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No permissions found")
	}
	permissions, err := s.repository.FindAllDataLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, permissions, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *PermissionService) GetAllPermissions(ctx context.Context) ([]entities.Permission, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No permissions found")
	}
	permissions, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	return permissions, nil
}

func (s *PermissionService) SearchPermissions(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchPermissionRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.PermissionData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NotFoundError("No permissions found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	permissions, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, permissions, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *PermissionService) GetPermissionByUniqueId(ctx context.Context, uniqueId string) (entities.PermissionData, error) {
	permission, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.PermissionData{}, apperrors.NotFoundError("permission not found")
	}
	return permission, nil
}

func (s *PermissionService) CountPermissions(ctx context.Context) (int64, error) {
	count, err := s.repository.Count(ctx)
	if err != nil {
		return 0, apperrors.NotFoundError("No permissions found")
	}
	return count, nil
}

func (s *PermissionService) GetPermissionRole(ctx context.Context, uniqueId string) (entities.PermissionRoleData, error) {
	userRole, err := s.permissionRoleRepository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.PermissionRoleData{}, apperrors.InternalServerError("permission role not found: " + err.Error())
	}
	return userRole, nil
}

func (s *PermissionService) RemovePermissionRole(ctx context.Context, uniqueId string) error {
	permissionRole, err := s.permissionRoleRepository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NotFoundError("permission role not found")
	}
	err = s.permissionRoleRepository.Delete(ctx, permissionRole)
	if err != nil {
		return apperrors.InternalServerError("error while removing permission: " + err.Error())
	}
	return nil
}

func (s *PermissionService) GetAssignedPermissionsByRole(ctx context.Context, roleId int) ([]entities.PermissionRoleData, error) {
	permissionRoles, err := s.permissionRoleRepository.FindAllDataByRoleId(ctx, roleId)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	return permissionRoles, nil
}

func (s *PermissionService) GetUnassignedPermissionsByRole(ctx context.Context, roleId int) ([]entities.Permission, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No permissions found")
	}
	roles, err := s.repository.FindUnassignedPermissionsByRole(ctx, roleId)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	return roles, nil
}
