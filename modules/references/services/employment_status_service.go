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

type EmploymentStatusService struct {
	repository *repositories.EmploymentStatusRepository
}

func NewEmploymentStatusService(db *database.Database) *EmploymentStatusService {
	return &EmploymentStatusService{
		repository: repositories.NewEmploymentStatusRepository(db),
	}
}

func (s *EmploymentStatusService) CreateEmploymentStatus(ctx context.Context, request entities.CreateStatusRequest) error {
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
	employmentStatus := entities.EmploymentStatus{
		StatusName:  request.StatusName,
		Code:        request.Code,
		Description: request.Description,
		UniqueId:    encryption.GenerateUUID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	err = s.repository.Create(ctx, employmentStatus)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating status: " + err.Error())
	}
	return nil
}

func (s *EmploymentStatusService) UpdateEmploymentStatus(ctx context.Context, employmentStatusId int, request entities.UpdateStatusRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	employmentStatus, err := s.repository.FindById(ctx, employmentStatusId)
	if err != nil {
		return apperrors.NewNotFoundError("status not found")
	}
	employmentStatus.StatusName = request.StatusName
	employmentStatus.Code = request.Code
	employmentStatus.Description = request.Description
	employmentStatus.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, employmentStatus)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating status: " + err.Error())
	}
	return nil
}

func (s *EmploymentStatusService) GetAllStatusesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.EmploymentStatus], error) {
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

func (s *EmploymentStatusService) GetAllStatuses(ctx context.Context) ([]entities.EmploymentStatus, error) {
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

func (s *EmploymentStatusService) SearchStatuses(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchStatusRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.EmploymentStatus], error) {
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

func (s *EmploymentStatusService) GetEmploymentStatusByUniqueId(ctx context.Context, uniqueId string) (entities.EmploymentStatus, error) {
	employmentStatus, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.EmploymentStatus{}, apperrors.NewNotFoundError("status not found")
	}
	return employmentStatus, nil
}

func (s *EmploymentStatusService) RemoveEmploymentStatus(ctx context.Context, uniqueId string) error {
	employmentStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("status not found")
	}
	err = s.repository.Delete(ctx, employmentStatus)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing employmentStatus: "+ err.Error())
	}
	return nil
}