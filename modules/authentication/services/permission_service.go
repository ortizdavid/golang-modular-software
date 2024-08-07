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
)

type PermissionService struct {
	repository *repositories.PermissionRepository
}

func NewPermissionService(db *database.Database) *PermissionService {
	return &PermissionService{
		repository: repositories.NewPermissionRepository(db),
	}
}

func (s *PermissionService) CreatePermission(ctx context.Context, request entities.CreatePermissionRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	permission := entities.Permission{
		PermissionName:    request.PermissionName,
		Code:        request.Code,
		Description: request.Description,
		UniqueId:    encryption.GenerateUUID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	err := s.repository.Create(ctx, permission)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating permission: "+ err.Error())
	}
	return nil
}

func (s *PermissionService) UpdatePermission(ctx context.Context, permissionId int, request entities.UpdatePermissionRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	permission, err := s.repository.FindById(ctx, permissionId)
	if err != nil {
		return apperrors.NewNotFoundError("permission not found")
	}
	permission.PermissionName = request.PermissionName
	permission.Description = request.Description
	permission.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, permission)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating permission: "+ err.Error())
	}
	return nil
}

func (s *PermissionService) DeletePermission(ctx context.Context, permissionId int) error {
	permission, err := s.repository.FindById(ctx, permissionId)
	if err != nil {
		return apperrors.NewNotFoundError("permission not found")
	}
	err = s.repository.Delete(ctx, permission)
	if err != nil {
		return apperrors.NewInternalServerError("error while deleting permission: "+ err.Error())
	}
	return nil
}

func (s *PermissionService) GetAllPermissionsPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.Permission], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No permissions found")
	}
	permissions, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, permissions, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *PermissionService) GetAllPermissions(ctx context.Context) ([]entities.Permission, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No permissions found")
	}
	permissions, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	return permissions, nil
}

func (s *PermissionService) CountPermissions(ctx context.Context) (int64, error) {
	count, err := s.repository.Count(ctx)
	if err != nil {
		return 0, apperrors.NewNotFoundError("No users found")
	}
	return count, nil
}