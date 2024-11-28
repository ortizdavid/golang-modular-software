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

type EvaluationStatusService struct {
	repository *repositories.EvaluationStatusRepository
}

func NewEvaluationStatusService(db *database.Database) *EvaluationStatusService {
	return &EvaluationStatusService{
		repository: repositories.NewEvaluationStatusRepository(db),
	}
}

func (s *EvaluationStatusService) Create(ctx context.Context, request entities.CreateStatusRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByName(ctx, request.StatusName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.BadRequestError("status already exists")
	}
	evaluationStatus := entities.EvaluationStatus{
		StatusName:  request.StatusName,
		Code:        request.Code,
		LblColor:    request.LblColor,
		BgColor:     request.BgColor,
		Description: request.Description,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, evaluationStatus)
	if err != nil {
		return apperrors.InternalServerError("error while creating status: " + err.Error())
	}
	return nil
}

func (s *EvaluationStatusService) Update(ctx context.Context, uniqueId string, request entities.UpdateStatusRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	evaluationStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NotFoundError("status not found")
	}
	evaluationStatus.StatusName = request.StatusName
	evaluationStatus.Code = request.Code
	evaluationStatus.LblColor = request.LblColor
	evaluationStatus.BgColor = request.BgColor
	evaluationStatus.Description = request.Description
	evaluationStatus.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, evaluationStatus)
	if err != nil {
		return apperrors.InternalServerError("error while updating status: " + err.Error())
	}
	return nil
}

func (s *EvaluationStatusService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.EvaluationStatus], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No statuses found")
	}
	statuses, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, statuses, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *EvaluationStatusService) GetAll(ctx context.Context) ([]entities.EvaluationStatus, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No statuses found")
	}
	statuses, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	return statuses, nil
}

func (s *EvaluationStatusService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchStatusRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.EvaluationStatus], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NotFoundError("No statuses found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	statuses, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, statuses, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *EvaluationStatusService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.EvaluationStatus, error) {
	maritalStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.EvaluationStatus{}, apperrors.NotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *EvaluationStatusService) GetByName(ctx context.Context, name string) (entities.EvaluationStatus, error) {
	maritalStatus, err := s.repository.FindByField(ctx, "status_name", name)
	if err != nil {
		return entities.EvaluationStatus{}, apperrors.NotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *EvaluationStatusService) GetByCode(ctx context.Context, code string) (entities.EvaluationStatus, error) {
	maritalStatus, err := s.repository.FindByField(ctx, "code", code)
	if err != nil {
		return entities.EvaluationStatus{}, apperrors.NotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *EvaluationStatusService) Remove(ctx context.Context, uniqueId string) error {
	evaluationStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NotFoundError("status not found")
	}
	err = s.repository.Delete(ctx, evaluationStatus)
	if err != nil {
		return apperrors.InternalServerError("error while removing evaluationStatus: " + err.Error())
	}
	return nil
}
