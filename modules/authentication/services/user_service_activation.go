package services

import (
	"context"
	"time"

	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
)

func (s *UserService) ActivateUser(ctx context.Context, userId int64) error {
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. invalid user id")
	}
	user.IsActive = true
	user.Token = encryption.GenerateRandomToken()
	user.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, user)
	if err != nil {
		return apperrors.NewInternalServerError("error while activating user: " + err.Error())
	}
	return nil
}

func (s *UserService) DeactivateUser(ctx context.Context, userId int64) error {
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. inavlid user id")
	}
	user.IsActive = false
	user.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, user)
	if err != nil {
		return apperrors.NewInternalServerError("error while deactivating user: " + err.Error())
	}
	return nil
}