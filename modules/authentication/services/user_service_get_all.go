package services

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	entities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

func (s *UserService) GetAllUsers(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.UserData], error) {
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No users found")
	}
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.FindAllDataLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, users, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *UserService) GetAllActiveUsers(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.UserData], error) {
	count, err := s.CountUsersByStatus(ctx, true)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No active users found")
	}
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.FindAllByStatus(ctx, "Yes", params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, users, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *UserService) GetAllInactiveUsers(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.UserData], error) {
	count, err := s.CountUsersByStatus(ctx, false)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No inactive users found")
	}
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.FindAllByStatus(ctx, "No", params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, users, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *UserService) GetAllOnlineUsers(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.UserData], error) {
	count, err := s.CountUsersByActivityStatus(ctx, entities.ActivityStatusOnline)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No online users found")
	}
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.FindAllByActivityStatus(ctx, entities.ActivityStatusOnline, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, users, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *UserService) GetAllOfflineUsers(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.UserData], error) {
	count, err := s.CountUsersByActivityStatus(ctx, entities.ActivityStatusOffline)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No offline users found")
	}
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.FindAllByActivityStatus(ctx, entities.ActivityStatusOffline, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, users, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *UserService) SearchUsers(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchUserRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.UserData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No users found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	users, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, users, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *UserService) GetUsersWithoutAssociation(ctx context.Context, allowedRoles []string) ([]entities.User, error) {
	users, err := s.repository.FindUsersWithoutAssociation(ctx, allowedRoles)
	if err != nil {
		return nil, apperrors.NewNotFoundError("no users found")
	}
	return users, nil
}
