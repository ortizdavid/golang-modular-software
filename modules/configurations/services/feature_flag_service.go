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

type FeatureFlagService struct {
	repository        *repositories.FeatureFlagRepository
	featureRepository *repositories.FeatureRepository
}

func NewFeatureFlagService(db *database.Database) *FeatureFlagService {
	return &FeatureFlagService{
		repository:        repositories.NewFeatureFlagRepository(db),
		featureRepository: repositories.NewFeatureRepository(db),
	}
}

func (s *FeatureFlagService) CreateFeatureFlag(ctx context.Context, request entities.ManageFeatureFlagRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByFeatureId(ctx, request.FeatureId)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("Feature Flag already exists ")
	}
	feature := entities.FeatureFlag{
		FeatureId: request.FeatureId,
		Status:    request.Status,
		UniqueId:  encryption.GenerateUUID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	err = s.repository.Create(ctx, feature)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating feature: " + err.Error())
	}
	return nil
}

func (s *FeatureFlagService) UpdateFeatureFlag(ctx context.Context, featureId int, request entities.ManageFeatureFlagRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	feature, err := s.repository.FindById(ctx, featureId)
	if err != nil {
		return apperrors.NewNotFoundError("feature not found")
	}
	feature.FeatureId = request.FeatureId
	feature.Status = request.Status
	feature.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, feature)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating feature: " + err.Error())
	}
	return nil
}

func (s *FeatureFlagService) GetAllFeatureFlagsPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.FeatureFlagData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No features found")
	}
	features, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, features, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *FeatureFlagService) GetAllFeatureFlags(ctx context.Context) ([]entities.FeatureFlag, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No features found")
	}
	features, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return features, nil
}

func (s *FeatureFlagService) SearchCompanies(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchFeatureFlagRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.FeatureFlagData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No features found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	features, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, features, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *FeatureFlagService) GetFeatureFlagByUniqueId(ctx context.Context, uniqueId string) (entities.FeatureFlagData, error) {
	feature, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.FeatureFlagData{}, apperrors.NewNotFoundError("feature not found")
	}
	return feature, nil
}

func (s *FeatureFlagService) SearchFeatureFlags(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchFeatureFlagRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.FeatureFlagData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No features found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	features, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, features, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}
