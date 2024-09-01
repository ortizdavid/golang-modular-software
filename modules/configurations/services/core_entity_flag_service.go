package services

import (
	"context"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/repositories"
)

type CoreEntityFlagService struct {
	repository        *repositories.CoreEntityFlagRepository
	moduleRepository *repositories.ModuleRepository
}

func NewCoreEntityFlagService(db *database.Database) *CoreEntityFlagService {
	return &CoreEntityFlagService{
		repository:        repositories.NewCoreEntityFlagRepository(db),
		moduleRepository: repositories.NewModuleRepository(db),
	}
}

func (s *CoreEntityFlagService) ManageCoreEntityFlags(ctx context.Context, requests []entities.ManageCoreEntityFlagRequest) error {
    // Validate each request
    for _, req := range requests {
        if err := req.Validate(); err != nil {
            return apperrors.NewBadRequestError(err.Error())
        }
    }
    // Retrieve and update each core entity flag
    var moduleFlags []entities.CoreEntityFlag
    for _, req := range requests {
        flag, err := s.repository.FindById(ctx, req.FlagId)
        if err != nil {
            return apperrors.NewNotFoundError("core entity flag not found")
        }
        flag.Status = req.Status
        flag.UpdatedAt = time.Now().UTC() // Ensure time is in UTC
        moduleFlags = append(moduleFlags, flag)
    }
    // Perform bulk update
    if err := s.repository.UpdateBatch(ctx, moduleFlags); err != nil {
        return apperrors.NewInternalServerError("error while updating core entity flags: " + err.Error())
    }
    return nil
}


func (s *CoreEntityFlagService) GetAllCoreEntityFlagsPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.CoreEntityFlagData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No core entity flags found")
	}
	modules, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, modules, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *CoreEntityFlagService) GetAllCoreEntityFlags(ctx context.Context) ([]entities.CoreEntityFlagData, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No modules found")
	}
	modules, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return modules, nil
}

func (s *CoreEntityFlagService) GetCoreEntityFlagByUniqueId(ctx context.Context, uniqueId string) (entities.CoreEntityFlagData, error) {
	module, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.CoreEntityFlagData{}, apperrors.NewNotFoundError("module not found")
	}
	return module, nil
}

