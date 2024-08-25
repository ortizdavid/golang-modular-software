package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
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

func (s *ModuleFlagService) CreateModuleFlag(ctx context.Context, request entities.ManageModuleFlagRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByModuleId(ctx, request.ModuleId)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("Module Flag already exists ")
	}
	module := entities.ModuleFlag{
		ModuleId: request.ModuleId,
		Status:    request.Status,
		UniqueId:  encryption.GenerateUUID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	err = s.repository.Create(ctx, module)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating module: " + err.Error())
	}
	return nil
}

func (s *ModuleFlagService) UpdateModuleFlag(ctx context.Context, moduleId int, request entities.ManageModuleFlagRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	module, err := s.repository.FindById(ctx, moduleId)
	if err != nil {
		return apperrors.NewNotFoundError("module not found")
	}
	module.ModuleId = request.ModuleId
	module.Status = request.Status
	module.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, module)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating module: " + err.Error())
	}
	return nil
}

func (s *ModuleFlagService) GetAllCompaniesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.ModuleFlagData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No modules found")
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

func (s *ModuleFlagService) GetAllModuleFlags(ctx context.Context) ([]entities.ModuleFlag, error) {
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

func (s *ModuleFlagService) SearchCompanies(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchModuleFlagRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.ModuleFlagData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No modules found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	modules, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, modules, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *ModuleFlagService) GetModuleFlagByUniqueId(ctx context.Context, uniqueId string) (entities.ModuleFlagData, error) {
	module, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.ModuleFlagData{}, apperrors.NewNotFoundError("module not found")
	}
	return module, nil
}

func (s *ModuleFlagService) SearchModuleFlags(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchModuleFlagRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.ModuleFlagData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No modules found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	modules, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, modules, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}
