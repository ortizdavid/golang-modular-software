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
)

type MaritalStatusService struct {
	repository *repositories.MaritalStatusRepository
}

func NewMaritalStatusService(db *database.Database) *MaritalStatusService {
	return &MaritalStatusService{
		repository: repositories.NewMaritalStatusRepository(db),
	}
}

func (s *MaritalStatusService) CreateMaritalStatus(ctx context.Context, request entities.CreateStatusRequest) error {
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
	maritalStatus := entities.MaritalStatus{
		StatusName:  request.StatusName,
		Code:        request.Code,
		UniqueId:    encryption.GenerateUUID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	err = s.repository.Create(ctx, maritalStatus)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating status: " + err.Error())
	}
	return nil
}

func (s *MaritalStatusService) UpdateMaritalStatus(ctx context.Context, maritalStatusId int, request entities.UpdateStatusRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	maritalStatus, err := s.repository.FindById(ctx, maritalStatusId)
	if err != nil {
		return apperrors.NewNotFoundError("status not found")
	}
	maritalStatus.StatusName = request.StatusName
	maritalStatus.Code = request.Code
	maritalStatus.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, maritalStatus)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating status: " + err.Error())
	}
	return nil
}

func (s *MaritalStatusService) GetAllStatusesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.MaritalStatus], error) {
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

func (s *MaritalStatusService) GetAllStatuses(ctx context.Context) ([]entities.MaritalStatus, error) {
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

func (s *MaritalStatusService) SearchStatuses(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchStatusRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.MaritalStatus], error) {
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

func (s *MaritalStatusService) GetMaritalStatusByUniqueId(ctx context.Context, uniqueId string) (entities.MaritalStatus, error) {
	maritalStatus, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.MaritalStatus{}, apperrors.NewNotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *MaritalStatusService) RemoveMaritalStatus(ctx context.Context, uniqueId string) error {
	maritalStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("status not found")
	}
	err = s.repository.Delete(ctx, maritalStatus)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing maritalStatus: "+ err.Error())
	}
	return nil
}