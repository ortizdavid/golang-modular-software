package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	"github.com/ortizdavid/golang-modular-software/modules/references/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"

)

type ContactTypeService struct {
	repository *repositories.ContactTypeRepository
}

func NewContactTypeService(db *database.Database) *ContactTypeService {
	return &ContactTypeService{
		repository: repositories.NewContactTypeRepository(db),
	}
}

func (s *ContactTypeService) Create(ctx context.Context, request entities.CreateTypeRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByName(ctx, request.TypeName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("type already exists")
	}
	contactType := entities.ContactType{
		TypeName:  request.TypeName,
		Code:        request.Code,
		BaseEntity: shared.BaseEntity{
			UniqueId:         encryption.GenerateUUID(),
			CreatedAt:        time.Now().UTC(),
			UpdatedAt:        time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, contactType)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating type: " + err.Error())
	}
	return nil
}

func (s *ContactTypeService) Update(ctx context.Context, uniqueId string, request entities.UpdateTypeRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	contactType, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("type not found")
	}
	contactType.TypeName = request.TypeName
	contactType.Code = request.Code
	contactType.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, contactType)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating type: " + err.Error())
	}
	return nil
}

func (s *ContactTypeService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.ContactType], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No types found")
	}
	types, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, types, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *ContactTypeService) GetAll(ctx context.Context) ([]entities.ContactType, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No types found")
	}
	types, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return types, nil
}

func (s *ContactTypeService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchTypeRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.ContactType], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No types found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	types, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, types, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *ContactTypeService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.ContactType, error) {
	contactType, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.ContactType{}, apperrors.NewNotFoundError("type not found")
	}
	return contactType, nil
}

func (s *ContactTypeService) Remove(ctx context.Context, uniqueId string) error {
	err := s.repository.DeleteByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing type: "+ err.Error())
	}
	return nil
}