package services

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/repositories"
)

type RoleService struct {
	repository *repositories.RoleRepository
	userRoleRepository *repositories.UserRoleRepository
	permissionRepository *repositories.PermissionRepository
	permissionRoleRepository *repositories.PermissionRoleRepository
}

func NewRoleService(db *database.Database) *RoleService {
	return &RoleService{
		repository:               repositories.NewRoleRepository(db),
		userRoleRepository:       repositories.NewUserRoleRepository(db),
		permissionRepository:     repositories.NewPermissionRepository(db),
		permissionRoleRepository: repositories.NewPermissionRoleRepository(db),
	}
}

func (s *RoleService) CreateRole(ctx context.Context, request entities.CreateRoleRequest) error {
	// Validate the request
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	// Check if the role name already exists
	if exists, err := s.repository.ExistsByName(ctx, request.RoleName); err != nil {
		return err
	} else if exists {
		return apperrors.NewConflictError("Role with name '"+request.RoleName+"' already exists")
	}
	// Check if the role code already exists
	if exists, err := s.repository.ExistsByCode(ctx, request.Code); err != nil {
		return err
	} else if exists {
		return apperrors.NewConflictError("Role with code '"+request.Code+"' already exists")
	}
	// Create the new role entity
	role := entities.Role{
		RoleId:      0,
		RoleName:    request.RoleName,
		Code:        request.Code,
		Description: request.Description,
		Status:      request.Status,
		UniqueId:    encryption.GenerateUUID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	// Persist the role entity
	if err := s.repository.Create(ctx, role); err != nil {
		return apperrors.NewInternalServerError("Error while creating role: " + err.Error())
	}
	return nil
}

func (s *RoleService) UpdateRole(ctx context.Context, roleId int, request entities.UpdateRoleRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	role, err := s.repository.FindById(ctx, roleId)
	if err != nil {
		return apperrors.NewNotFoundError("role not found")
	}
	exists, err := s.userRoleRepository.ExistsByRoleId(ctx, roleId)
    if err != nil {
        return err
    }
    if exists {
        return apperrors.NewConflictError("Role '" + role.RoleName + "' is currently assigned to users and cannot be updated")
    }
	//--------------------------------
	role.RoleName = request.RoleName
	role.Code = request.Code
	role.Description = request.Description
	role.Status = request.Status
	role.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, role)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating role: "+ err.Error())
	}
	return nil
}

func (s *RoleService) AssignRolePermission(ctx context.Context, roleId int, request entities.AssignRolePermissionRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	role, err := s.repository.FindById(ctx, roleId)
	if err != nil {
		return apperrors.NewNotFoundError("role not found")
	}
	permission, err := s.permissionRepository.FindById(ctx, request.PermissionId)
	if err != nil {
		return apperrors.NewNotFoundError("permission not found")
	}
	exists, err := s.permissionRoleRepository.ExistsByRoleAndPermission(ctx, roleId, request.PermissionId)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewConflictError(fmt.Sprintf("permission '%s' already assigned to role '%s'", permission.PermissionName, role.RoleName))
	}
	permissionRole := entities.PermissionRole{
		PermissionId:     int64(request.PermissionId),
		RoleId:           roleId,
		UniqueId:         encryption.GenerateUUID(),
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	}
	err = s.permissionRoleRepository.Create(ctx, permissionRole)
	if err != nil {
		return apperrors.NewInternalServerError("error while assigning permission: "+ err.Error())
	}
	return nil
}

func (s *RoleService) DeleteRole(ctx context.Context, roleId int) error {
    role, err := s.repository.FindById(ctx, roleId)
    if err != nil {
        return apperrors.NewNotFoundError("role not found")
    }
    exists, err := s.userRoleRepository.ExistsByRoleId(ctx, roleId)
    if err != nil {
        return err
    }
    if exists {
        return apperrors.NewConflictError("Role '" + role.RoleName + "' is currently assigned to users and cannot be deleted")
    }
    err = s.repository.Delete(ctx, role)
    if err != nil {
        return apperrors.NewInternalServerError("error while deleting role: " + err.Error())
    }
    return nil
}

func (s *RoleService) GetAllRolesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.Role], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No roles found")
	}
	roles, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, roles, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *RoleService) GetAllRoles(ctx context.Context) ([]entities.Role, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No roles found")
	}
	roles, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	return roles, nil
}

func (s *RoleService) GetAllEnaledRoles(ctx context.Context) ([]entities.Role, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No roles found")
	}
	roles, err := s.repository.FindAllEnabled(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	return roles, nil
}

func (s *RoleService) GetUnassignedRolesByUser(ctx context.Context, userId int64) ([]entities.Role, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No roles found")
	}
	roles, err := s.repository.FindUnassignedRolesByUser(ctx, userId)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	return roles, nil
}

func (s *RoleService) SearchRoles(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchRoleRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.RoleData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No roles found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	roles, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, roles, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *RoleService) GetRoleByUniqueId(ctx context.Context, uniqueId string) (entities.RoleData, error) {
	role, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.RoleData{}, apperrors.NewNotFoundError("role not found")
	}
	return role, nil
}

func (s *RoleService) CountRoles(ctx context.Context) (int64, error) {
	count, err := s.repository.Count(ctx)
	if err != nil {
		return 0, apperrors.NewNotFoundError("No roles found")
	}
	return count, nil
}

func (s *RoleService) GetUserRole(ctx context.Context, uniqueId string) (entities.UserRoleData, error) {
	userRole, err := s.userRoleRepository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.UserRoleData{}, apperrors.NewInternalServerError("user role not found: "+err.Error())
	}
	return userRole, nil
}

func (s *RoleService) GetAssignedRolesByUser(ctx context.Context, userId int64) ([]entities.UserRoleData, error) {
	userRoles, err := s.userRoleRepository.FindAllByUserId(ctx, userId)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	return userRoles, nil
}