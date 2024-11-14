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
	if user.UserId == 0 {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found. invalid id")
	}
	return user, nil
}

func (s *UserService) GetUserByUniqueId(ctx context.Context, uniqueId string) (entities.UserData, error) {
	user, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found. invalid unique id")
	}
	if user.UserId == 0 {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found. invalid unique id")
	}
	return user, nil
}

func (s *UserService) GetUserByName(ctx context.Context, email string) (entities.UserData, error) {
	user, err := s.repository.GetDataByUserName(ctx, email)
	if err != nil {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found")
	}
	if user.UserId == 0 {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found")
	}
	return user, nil
}

func (s *UserService) GetUserApiKey(ctx context.Context, xUserId string) (entities.UserApiKey, error) {
	user, err := s.userApiKeyRepository.FindByXUserId(ctx, xUserId)
	if err != nil {
		return entities.UserApiKey{}, apperrors.NewNotFoundError("user api key not found. invalid x-user-id")
	}
	if user.UserId == 0 {
		return entities.UserApiKey{}, apperrors.NewNotFoundError("user api key not found. invalid x-user-id")
	}
	return user, nil
}

func (s *UserService) GetUserApiKeyById(ctx context.Context, userId int64) (entities.UserApiKey, error) {
	user, err := s.userApiKeyRepository.FindByUserId(ctx, userId)
	if err != nil {
		return entities.UserApiKey{}, apperrors.NewNotFoundError("user api key not found")
	}
	if user.UserId == 0 {
		return entities.UserApiKey{}, apperrors.NewNotFoundError("user api key not found")
	}
	return user, nil
}

func (s *UserService) GetUserByToken(ctx context.Context, token string) (entities.UserData, error) {
	user, err := s.repository.GetDataByToken(ctx, token)
	if err != nil {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found. invalid token")
	}
	if user.UserId == 0 {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found. invalid token")
	}
	return user, nil
}

func (s *UserService) FindUserByToken(ctx context.Context, token string) (entities.User, error) {
	user, err := s.repository.FindByToken(ctx, token)
	if err != nil {
		return entities.User{}, apperrors.NewNotFoundError("user not found. invalid token")
	}
	if user.UserId == 0 {
		return entities.User{}, apperrors.NewNotFoundError("user not found. invalid token")
	}
	return user, nil
}

func (s *UserService) FindUserByEmail(ctx context.Context, email string) (entities.User, error) {
	user, err := s.repository.FindByEmail(ctx, email)
	if err != nil {
		return entities.User{}, apperrors.NewNotFoundError("user not found. invalid email")
	}
	if user.UserId == 0 {
		return entities.User{}, apperrors.NewNotFoundError("user not found. invalid email")
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (entities.UserData, error) {
	user, err := s.repository.GetDataByEmail(ctx, email)
	if err != nil {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found")
	}
	if user.UserId == 0 {
		return entities.UserData{}, apperrors.NewNotFoundError("user not found")
	}
	return user, nil
}