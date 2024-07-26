package services

import (
	"context"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/repositories"
	"gorm.io/gorm"
)

type RoleService struct {
	repository *repositories.RoleRepository
}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{
		repository: repositories.NewRoleRepository(db),
	}
}

func (s *RoleService) CreateRole(ctx context.Context, request entities.CreateRoleRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	role := entities.Role{
		RoleName:    request.RoleName,
		Code:        request.Code,
		Description: request.Description,
		UniqueId:    encryption.GenerateUUID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	err := s.repository.Create(ctx, role)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating role: "+ err.Error())
	}
	return nil
}

func (s *RoleService) UpdateRole(ctx context.Context, roleId int, request entities.CreateRoleRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	role, err := s.repository.FindById(ctx, roleId)
	if err != nil {
		return apperrors.NewNotFoundError("role not found")
	}
	role.RoleName = request.RoleName
	role.Description = request.Description
	role.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, role)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating role: "+ err.Error())
	}
	return nil
}

func (s *RoleService) DeleteRole(ctx context.Context, roleId int) error {
	role, err := s.repository.FindById(ctx, roleId)
	if err != nil {
		return apperrors.NewNotFoundError("role not found")
	}
	err = s.repository.Delete(ctx, role)
	if err != nil {
		return apperrors.NewInternalServerError("error while deleting role: "+ err.Error())
	}
	return nil
}

func (s *RoleService) GetAllRolesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.Role], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count()
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
	_, err := s.repository.Count()
	if err != nil {
		return nil, apperrors.NewNotFoundError("No roles found")
	}
	roles, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	
	return roles, nil
}