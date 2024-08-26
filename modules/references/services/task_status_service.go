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

type TaskStatusService struct {
	repository *repositories.TaskStatusRepository
}

func NewTaskStatusService(db *database.Database) *TaskStatusService {
	return &TaskStatusService{
		repository: repositories.NewTaskStatusRepository(db),
	}
}

func (s *TaskStatusService) CreateTaskStatus(ctx context.Context, request entities.CreateStatusRequest) error {
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
	taskStatus := entities.TaskStatus{
		StatusName:  request.StatusName,
		Code:        request.Code,
		LblColor:    request.LblColor,
		BgColor:     request.BgColor,
		Description: request.Description,
		UniqueId:    encryption.GenerateUUID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	err = s.repository.Create(ctx, taskStatus)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating status: " + err.Error())
	}
	return nil
}

func (s *TaskStatusService) UpdateTaskStatus(ctx context.Context, taskStatusId int, request entities.UpdateStatusRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	taskStatus, err := s.repository.FindById(ctx, taskStatusId)
	if err != nil {
		return apperrors.NewNotFoundError("status not found")
	}
	taskStatus.StatusName = request.StatusName
	taskStatus.Code = request.Code
	taskStatus.LblColor = request.LblColor
	taskStatus.BgColor = request.BgColor
	taskStatus.Description = request.Description
	taskStatus.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, taskStatus)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating status: " + err.Error())
	}
	return nil
}

func (s *TaskStatusService) GetAllStatusesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.TaskStatus], error) {
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

func (s *TaskStatusService) GetAllStatuses(ctx context.Context) ([]entities.TaskStatus, error) {
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

func (s *TaskStatusService) SearchStatuses(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchStatusRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.TaskStatus], error) {
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

func (s *TaskStatusService) GetTaskStatusByUniqueId(ctx context.Context, uniqueId string) (entities.TaskStatus, error) {
	taskStatus, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.TaskStatus{}, apperrors.NewNotFoundError("status not found")
	}
	return taskStatus, nil
}

func (s *TaskStatusService) RemoveTaskStatus(ctx context.Context, uniqueId string) error {
	taskStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("status not found")
	}
	err = s.repository.Delete(ctx, taskStatus)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing taskStatus: "+ err.Error())
	}
	return nil
}