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

type DocumentStatusService struct {
	repository *repositories.DocumentStatusRepository
}

func NewDocumentStatusService(db *database.Database) *DocumentStatusService {
	return &DocumentStatusService{
		repository: repositories.NewDocumentStatusRepository(db),
	}
}

func (s *DocumentStatusService) Create(ctx context.Context, request entities.CreateStatusRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByName(ctx, request.StatusName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("status already exists")
	}
	documentStatus := entities.DocumentStatus{
		StatusName:  request.StatusName,
		Code:        request.Code,
		LblColor:    request.LblColor,
		BgColor:     request.BgColor,
		Description: request.Description,
		BaseEntity: shared.BaseEntity{
			UniqueId:         encryption.GenerateUUID(),
			CreatedAt:        time.Now().UTC(),
			UpdatedAt:        time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, documentStatus)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating status: " + err.Error())
	}
	return nil
}

func (s *DocumentStatusService) Update(ctx context.Context, uniqueId string, request entities.UpdateStatusRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	documentStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("status not found")
	}
	documentStatus.StatusName = request.StatusName
	documentStatus.Code = request.Code
	documentStatus.LblColor = request.LblColor
	documentStatus.BgColor = request.BgColor
	documentStatus.Description = request.Description
	documentStatus.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, documentStatus)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating status: " + err.Error())
	}
	return nil
}

func (s *DocumentStatusService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.DocumentStatus], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No statuses found")
	}
	statuses, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, statuses, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *DocumentStatusService) GetAll(ctx context.Context) ([]entities.DocumentStatus, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No statuses found")
	}
	statuses, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return statuses, nil
}

func (s *DocumentStatusService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchStatusRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.DocumentStatus], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No statuses found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	statuses, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, statuses, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *DocumentStatusService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.DocumentStatus, error) {
	maritalStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.DocumentStatus{}, apperrors.NewNotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *DocumentStatusService) GetByName(ctx context.Context, name string) (entities.DocumentStatus, error) {
	maritalStatus, err := s.repository.FindByField(ctx, "status_name", name)
	if err != nil {
		return entities.DocumentStatus{}, apperrors.NewNotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *DocumentStatusService) GetByCode(ctx context.Context, code string) (entities.DocumentStatus, error) {
	maritalStatus, err := s.repository.FindByField(ctx, "code", code)
	if err != nil {
		return entities.DocumentStatus{}, apperrors.NewNotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *DocumentStatusService) Remove(ctx context.Context, uniqueId string) error {
	documentStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("status not found")
	}
	err = s.repository.Delete(ctx, documentStatus)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing documentStatus: "+ err.Error())
	}
	return nil
}