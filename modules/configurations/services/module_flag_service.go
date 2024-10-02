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

type ModuleFlagService struct {
	repository        *repositories.ModuleFlagRepository
	moduleRepository *repositories.ModuleRepository
}

func NewModuleFlagService(db *database.Database) *ModuleFlagService {
	return &ModuleFlagService{
		repository:        repositories.NewModuleFlagRepository(db),
		moduleRepository: repositories.NewModuleRepository(db),
	}
}

func (s *ModuleFlagService) ManageModuleFlags(ctx context.Context, requests []entities.ManageModuleFlagRequest) error {
    // Validate each request
    for _, req := range requests {
        if err := req.Validate(); err != nil {
            return apperrors.NewBadRequestError(err.Error())
        }
    }
    // Retrieve and update each module flag
    var moduleFlags []entities.ModuleFlag
    for _, req := range requests {
        flag, err := s.repository.FindById(ctx, req.FlagId)
        if err != nil {
            return apperrors.NewNotFoundError("module flag not found")
        }
        flag.Status = req.Status
        flag.UpdatedAt = time.Now().UTC() // Ensure time is in UTC
        moduleFlags = append(moduleFlags, flag)
    }
    // Perform bulk update
    if err := s.repository.UpdateBatch(ctx, moduleFlags); err != nil {
        return apperrors.NewInternalServerError("error while updating module flags: " + err.Error())
    }
    return nil
}


func (s *ModuleFlagService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.ModuleFlagData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No module flags found")
	}
	modules, err := s.repository.FindAllDataLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, modules, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *ModuleFlagService) GetAll(ctx context.Context) ([]entities.ModuleFlagData, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No modules found")
	}
	modules, err := s.repository.FindAllData(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return modules, nil
}

func (s *ModuleFlagService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.ModuleFlagData, error) {
	module, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.ModuleFlagData{}, apperrors.NewNotFoundError("module not found")
	}
	return module, nil
}

func (s *ModuleFlagService) GetByCode(ctx context.Context, code string) (entities.ModuleFlagData, error) {
	moduleFlag, err := s.repository.FindByModuleCode(ctx, code)
	if err != nil {
		return entities.ModuleFlagData{}, apperrors.NewNotFoundError("module flag not found")
	}
	return moduleFlag, nil
}


