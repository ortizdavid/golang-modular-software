package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/repositories"
)

type CoreEntityService struct {
	repository        *repositories.CoreEntityRepository
}

func NewCoreEntityService(db *database.Database) *CoreEntityService {
	return &CoreEntityService{
		repository:        repositories.NewCoreEntityRepository(db),
	}
}

func (s *CoreEntityService) CreateCoreEntity(ctx context.Context, request entities.CreateCoreEntityRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByName(ctx, request.EntityName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("CoreEntity already exists " + request.EntityName)
	}
	coreEntity := entities.CoreEntity{
		ModuleId:    request.ModuleId,
		EntityName:  request.EntityName,
		Code:        request.Code,
		Description: request.Description,
		UniqueId:    encryption.GenerateUUID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	err = s.repository.Create(ctx, coreEntity)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating coreEntity: " + err.Error())
	}
	return nil
}

func (s *CoreEntityService) UpdateCoreEntity(ctx context.Context, entityId int, request entities.UpdateCoreEntityRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	coreEntity, err := s.repository.FindById(ctx, entityId)
	if err != nil {
		return apperrors.NewNotFoundError("coreEntity not found")
	}
	coreEntity.ModuleId = request.ModuleId
	coreEntity.Code = request.Code
	coreEntity.EntityName = request.EntityName
	coreEntity.Description = request.Description
	coreEntity.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, coreEntity)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating core entity: " + err.Error())
	}
	return nil
}

func (s *CoreEntityService) GetAllCoreEntitiesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.CoreEntityData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No core entities found")
	}
	coreEntities, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, coreEntities, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *CoreEntityService) GetAllCoreEntities(ctx context.Context) ([]entities.CoreEntity, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No core entities found")
	}
	coreEntities, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return coreEntities, nil
}

func (s *CoreEntityService) SearchCompanies(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchCoreEntityRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.CoreEntityData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No core entities found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	coreEntities, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, coreEntities, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *CoreEntityService) GetCoreEntityByUniqueId(ctx context.Context, uniqueId string) (entities.CoreEntityData, error) {
	coreEntity, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.CoreEntityData{}, apperrors.NewNotFoundError("core entity not found")
	}
	return coreEntity, nil
}

func (s *CoreEntityService) SearchCoreEntities(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchCoreEntityRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.CoreEntityData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No core entities found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	coreEntities, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, coreEntities, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}