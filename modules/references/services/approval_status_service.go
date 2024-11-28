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

type ApprovalStatusService struct {
	repository *repositories.ApprovalStatusRepository
}

func NewApprovalStatusService(db *database.Database) *ApprovalStatusService {
	return &ApprovalStatusService{
		repository: repositories.NewApprovalStatusRepository(db),
	}
}

func (s *ApprovalStatusService) Create(ctx context.Context, request entities.CreateStatusRequest) error {
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
	approvalStatus := entities.ApprovalStatus{
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
	err = s.repository.Create(ctx, approvalStatus)
	if err != nil {
		return apperrors.InternalServerError("error while creating status: " + err.Error())
	}
	return nil
}

func (s *ApprovalStatusService) Update(ctx context.Context, uniqueId string, request entities.UpdateStatusRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	approvalStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NotFoundError("status not found")
	}
	approvalStatus.StatusName = request.StatusName
	approvalStatus.Code = request.Code
	approvalStatus.BgColor = request.BgColor
	approvalStatus.LblColor = request.LblColor
	approvalStatus.Description = request.Description
	approvalStatus.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, approvalStatus)
	if err != nil {
		return apperrors.InternalServerError("error while updating status: " + err.Error())
	}
	return nil
}

func (s *ApprovalStatusService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.ApprovalStatus], error) {
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

func (s *ApprovalStatusService) GetAll(ctx context.Context) ([]entities.ApprovalStatus, error) {
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

func (s *ApprovalStatusService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchStatusRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.ApprovalStatus], error) {
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

func (s *ApprovalStatusService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.ApprovalStatus, error) {
	maritalStatus, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.ApprovalStatus{}, apperrors.NotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *ApprovalStatusService) GetByName(ctx context.Context, name string) (entities.ApprovalStatus, error) {
	maritalStatus, err := s.repository.FindByField(ctx, "status_name", name)
	if err != nil {
		return entities.ApprovalStatus{}, apperrors.NotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *ApprovalStatusService) GetByCode(ctx context.Context, code string) (entities.ApprovalStatus, error) {
	maritalStatus, err := s.repository.FindByField(ctx, "code", code)
	if err != nil {
		return entities.ApprovalStatus{}, apperrors.NotFoundError("status not found")
	}
	return maritalStatus, nil
}

func (s *ApprovalStatusService) Remove(ctx context.Context, uniqueId string) error {
	err := s.repository.DeleteByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.InternalServerError("error while removing approvalStatus: " + err.Error())
	}
	return nil
}
