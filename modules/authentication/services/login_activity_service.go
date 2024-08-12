package services

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/repositories"
)

type LoginActivityService struct {
	repository *repositories.LoginActivityRepository
}

func NewLoginActivityService(db *database.Database) *LoginActivityService {
	return &LoginActivityService{
		repository: repositories.NewLoginActivityRepository(db),
	}
}

func (s *LoginActivityService) GetAllLoginActivities(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.LoginActivityData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No login activities found")
	}
	loginActivities, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, loginActivities, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *LoginActivityService) SearchLoginActivities(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchLoginActivityRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.LoginActivityData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No activities found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	loginActivities, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, loginActivities, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *LoginActivityService) GetLoginActivityByUniqueId(ctx context.Context, uniqueId string) (entities.LoginActivityData, error) {
	loginActivity, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.LoginActivityData{}, apperrors.NewNotFoundError("login activity not found")
	}
	return loginActivity, nil
}
