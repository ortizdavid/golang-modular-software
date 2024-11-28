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
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type ModuleService struct {
	repository *repositories.ModuleRepository
}

func NewModuleService(db *database.Database) *ModuleService {
	return &ModuleService{
		repository: repositories.NewModuleRepository(db),
	}
}

func (s *ModuleService) Create(ctx context.Context, request entities.CreateModuleRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByName(ctx, request.ModuleName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.BadRequestError("Module already exists " + request.ModuleName)
	}
	module := entities.Module{
		ModuleName:  request.ModuleName,
		Code:        request.Code,
		Description: request.Description,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, module)
	if err != nil {
		return apperrors.InternalServerError("error while creating module: " + err.Error())
	}
	return nil
}

func (s *ModuleService) Update(ctx context.Context, uniqueId string, request entities.UpdateModuleRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	module, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NotFoundError("module not found")
	}
	module.ModuleName = request.ModuleName
	module.Code = request.Code
	module.Description = request.Description
	module.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, module)
	if err != nil {
		return apperrors.InternalServerError("error while updating module: " + err.Error())
	}
	return nil
}

func (s *ModuleService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.Module], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No modules found")
	}
	modules, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, modules, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *ModuleService) GetAll(ctx context.Context) ([]entities.Module, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No modules found")
	}
	modules, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	return modules, nil
}

/*func (s *ModuleService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchModuleRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.Module], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NotFoundError("No modules found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	modules, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, modules, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}*/

func (s *ModuleService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.Module, error) {
	module, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.Module{}, apperrors.NotFoundError("module not found")
	}
	return module, nil
}

func (s *ModuleService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchModuleRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.Module], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NotFoundError("No modules found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	modules, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, modules, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}
