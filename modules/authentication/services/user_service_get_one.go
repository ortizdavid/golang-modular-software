package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	entities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

func (s *UserService) GetUserById(ctx context.Context, userId int64) (entities.UserData, error) {
	user, err := s.repository.GetDataById(ctx, userId)
	if err != nil {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found. invalid id")
	}
	return user, nil
}

func (s *UserService) GetUserByUniqueId(ctx context.Context, uniqueId string) (entities.UserData, error) {
	user, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found. invalid unique id")
	}
	return user, nil
}

func (s *UserService) GetUserApiKey(ctx context.Context, xUserId string) (entities.UserApiKey, error) {
	user, err := s.userApiKeyRepository.FindByXUserId(ctx, xUserId)
	if err != nil {
		return entities.UserApiKey{}, apperrors.NewNotFoundError("user api key not found. invalid x-user-id")
	}
	return user, nil
}

func (s *UserService) GetUserApiKeyById(ctx context.Context, userId int64) (entities.UserApiKey, error) {
	user, err := s.userApiKeyRepository.FindByUserId(ctx, userId)
	if err != nil {
		return entities.UserApiKey{}, apperrors.NewNotFoundError("user api key not found")
	}
	return user, nil
}

func (s *UserService) GetUserByToken(ctx context.Context, token string) (entities.User, error) {
	user, err := s.repository.FindByToken(ctx, token)
	if err != nil {
		return entities.User{}, apperrors.NewNotFoundError("user not found. invalid token")
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, token string) (entities.User, error) {
	user, err := s.repository.FindByEmail(ctx, token)
	if err != nil {
		return entities.User{}, apperrors.NewNotFoundError("user not found")
	}
	return user, nil
}