package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type DocumentTypeService struct {
	repository *repositories.DocumentTypeRepository
}

func NewDocumentTypeService(db *database.Database) *DocumentTypeService {
	return &DocumentTypeService{
		repository: repositories.NewDocumentTypeRepository(db),
	}
}


func (s *DocumentTypeService) Create(ctx context.Context, request entities.CreateDocumentTypeRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByName(ctx, request.TypeName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("document type already exists")
	}
	documentType := entities.DocumentType{
		TypeName:    request.TypeName,
		Description: request.Description,
		BaseEntity:  shared.BaseEntity{
			UniqueId:       encryption.GenerateUUID(),
			CreatedAt:      time.Now().UTC(),
			UpdatedAt:      time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, documentType)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating document type: " + err.Error())
	}
	return nil
}

func (s *DocumentTypeService) Update(ctx context.Context, documentTypeId int, request entities.UpdateDocumentTypeRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	documentType, err := s.repository.FindById(ctx, documentTypeId)
	if err != nil {
		return apperrors.NewNotFoundError("document type not found")
	}
	documentType.TypeName = request.TypeName
	documentType.Description = request.Description
	documentType.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, documentType)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating document type: " + err.Error())
	}
	return nil
}

func (s *DocumentTypeService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.DocumentType], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No document types found")
	}
	documentTypes, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, documentTypes, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *DocumentTypeService) GetAll(ctx context.Context) ([]entities.DocumentType, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No document types found")
	}
	documentTypes, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return documentTypes, nil
}

func (s *DocumentTypeService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchDocumentTypeRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.DocumentType], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No document types found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	documentTypes, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, documentTypes, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *DocumentTypeService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.DocumentType, error) {
	documentType, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.DocumentType{}, apperrors.NewNotFoundError("document type not found")
	}
	return documentType, nil
}

func (s *DocumentTypeService) Remove(ctx context.Context, uniqueId string) error {
	err := s.repository.DeleteByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing document type: "+ err.Error())
	}
	return nil
}