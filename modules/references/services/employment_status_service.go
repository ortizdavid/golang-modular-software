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

type EmploymentStatusService struct {
	repository *repositories.EmploymentStatusRepository
}

func NewEmploymentStatusService(db *database.Database) *EmploymentStatusService {
	return &EmploymentStatusService{
		repository: repositories.NewEmploymentStatusRepository(db),
	}
}

func (s *EmploymentStatusService) Create(ctx context.Context, request entities.CreateStatusRequest) error {
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
	employmentStatus := entities.EmploymentStatus{
		StatusName:  request.StatusName,
		Code:        request.Code,
		Description: request.Description,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, employmentStatus)
	if err != nil {
		return apperrors.InternalServerError("error while creating status: " + err.Error())
	}
	return nil
}

func (s *EmploymentStatusService) Update(ctx context.Context, uniqueId string, request entities.UpdateStatusRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	employmentStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NotFoundError("status not found")
	}
	employmentStatus.StatusName = request.StatusName
	employmentStatus.Code = request.Code
	employmentStatus.Description = request.Description
	employmentStatus.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, employmentStatus)
	if err != nil {
		return apperrors.InternalServerError("error while updating status: " + err.Error())
	}
	return nil
}

func (s *EmploymentStatusService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.EmploymentStatus], error) {
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

func (s *EmploymentStatusService) GetAll(ctx context.Context) ([]entities.EmploymentStatus, error) {
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

func (s *EmploymentStatusService) SearchStatuses(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchStatusRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.EmploymentStatus], error) {
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

func (s *EmploymentStatusService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.EmploymentStatus, error) {
	maritalStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.EmploymentStatus{}, apperrors.NotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *EmploymentStatusService) GetByName(ctx context.Context, name string) (entities.EmploymentStatus, error) {
	maritalStatus, err := s.repository.FindByField(ctx, "status_name", name)
	if err != nil {
		return entities.EmploymentStatus{}, apperrors.NotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *EmploymentStatusService) GetByCode(ctx context.Context, code string) (entities.EmploymentStatus, error) {
	maritalStatus, err := s.repository.FindByField(ctx, "code", code)
	if err != nil {
		return entities.EmploymentStatus{}, apperrors.NotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *EmploymentStatusService) Remove(ctx context.Context, uniqueId string) error {
	employmentStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NotFoundError("status not found")
	}
	err = s.repository.Delete(ctx, employmentStatus)
	if err != nil {
		return apperrors.InternalServerError("error while removing employmentStatus: " + err.Error())
	}
	return nil
}
