package services

import (
	"context"
	"time"

	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	entities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

func (s *UserService) ChangeUserPassword(ctx context.Context, userId int64, request entities.UpdatePasswordRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	return s.updatePassword(ctx, userId, request)
}

func (s *UserService) ResetUserPassword(ctx context.Context, userId int64, request entities.UpdatePasswordRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	return s.updatePassword(ctx, userId, request)
}

func (s *UserService) updatePassword(ctx context.Context, userId int64, request entities.UpdatePasswordRequest) error {
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NotFoundError("user not found. invalid id")
	}
	user.Password = encryption.HashPassword(request.NewPassword)
	user.Token = encryption.GenerateRandomToken()
	user.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, user)
	if err != nil {
		return apperrors.InternalServerError("error while updating password: " + err.Error())
	}
	return nil
}
