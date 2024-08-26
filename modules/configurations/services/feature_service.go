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

type FeatureService struct {
	repository        *repositories.FeatureRepository
}

func NewFeatureService(db *database.Database) *FeatureService {
	return &FeatureService{
		repository:        repositories.NewFeatureRepository(db),
	}
}

func (s *FeatureService) CreateFeature(ctx context.Context, request entities.CreateFeatureRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByName(ctx, request.FeatureName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("Feature already exists " + request.FeatureName)
	}
	feature := entities.Feature{
		FeatureName:  request.FeatureName,
		Description: request.Description,
		UniqueId:    encryption.GenerateUUID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	err = s.repository.Create(ctx, feature)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating feature: " + err.Error())
	}
	return nil
}

func (s *FeatureService) UpdateFeature(ctx context.Context, featureId int, request entities.UpdateFeatureRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	feature, err := s.repository.FindById(ctx, featureId)
	if err != nil {
		return apperrors.NewNotFoundError("feature not found")
	}
	feature.FeatureName = request.FeatureName
	feature.Description = request.Description
	feature.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, feature)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating feature: " + err.Error())
	}
	return nil
}

func (s *FeatureService) GetAllFeaturesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.FeatureData], error) {
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

func (s *FeatureService) GetAllFeatures(ctx context.Context) ([]entities.Feature, error) {
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

func (s *FeatureService) SearchCompanies(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchFeatureRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.FeatureData], error) {
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

func (s *FeatureService) GetFeatureByUniqueId(ctx context.Context, uniqueId string) (entities.FeatureData, error) {
	feature, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.FeatureData{}, apperrors.NewNotFoundError("feature not found")
	}
	return feature, nil
}

func (s *FeatureService) SearchFeatures(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchFeatureRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.FeatureData], error) {
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
