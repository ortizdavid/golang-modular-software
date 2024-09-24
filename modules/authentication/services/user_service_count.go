package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	entities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

func (s *UserService) CountUsers(ctx context.Context) (int64, error) {
	count, err := s.repository.Count(ctx)
	if err != nil {
		return 0, apperrors.NewNotFoundError("No users found")
	}
	return count, nil
}

func (s *UserService) CountUsersByStatus(ctx context.Context, status bool) (int64, error) {
	count, err := s.repository.CountByStatus(ctx, status)
	if err != nil {
		return 0, apperrors.NewNotFoundError("No users found")
	}
	return count, nil
}

func (s *UserService) CountUsersByActivityStatus(ctx context.Context, status entities.LoginActivityStatus) (int64, error) {
	count, err := s.repository.CountByActivityStatus(ctx, status)
	if err != nil {
		return 0, apperrors.NewNotFoundError("No users found")
	}
	return count, nil
}
