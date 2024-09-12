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
	coreEntityRepository *repositories.ModuleRepository
}

func NewCoreEntityFlagService(db *database.Database) *CoreEntityFlagService {
	return &CoreEntityFlagService{
		repository:        repositories.NewCoreEntityFlagRepository(db),
		coreEntityRepository: repositories.NewModuleRepository(db),
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
    var coreEntityFlags []entities.CoreEntityFlag
    for _, req := range requests {
        flag, err := s.repository.FindById(ctx, req.FlagId)
        if err != nil {
            return apperrors.NewNotFoundError("core entity flag not found")
        }
        flag.Status = req.Status
        flag.UpdatedAt = time.Now().UTC() // Ensure time is in UTC
        coreEntityFlags = append(coreEntityFlags, flag)
    }
    // Perform bulk update
    if err := s.repository.UpdateBatch(ctx, coreEntityFlags); err != nil {
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
	coreEntities, err := s.repository.FindAllDataLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, coreEntities, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *CoreEntityFlagService) GetAllCoreEntityFlags(ctx context.Context) ([]entities.CoreEntityFlagData, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No core entities found")
	}
	coreEntities, err := s.repository.FindAllData(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return coreEntities, nil
}

func (s *CoreEntityFlagService) GetCoreEntityFlagByUniqueId(ctx context.Context, uniqueId string) (entities.CoreEntityFlagData, error) {
	coreEntity, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.CoreEntityFlagData{}, apperrors.NewNotFoundError("core entity not found")
	}
	return coreEntity, nil
}

func (s *CoreEntityFlagService) GetCoreEntityFlagByCode(ctx context.Context, code string) (entities.CoreEntityFlagData, error) {
	coreEntity, err := s.repository.FindByEntityCode(ctx, code)
	if err != nil {
		return entities.CoreEntityFlagData{}, apperrors.NewNotFoundError("core entity not found")
	}
	return coreEntity, nil
}

