package services

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/config"
	entities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/repositories"
	"gorm.io/gorm"
)

type AuthService struct {
	repository *repositories.UserRepository
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		repository: repositories.NewUserRepository(db),
	}
}

func (s *AuthService) Authenticate(ctx context.Context, fiberCtx *fiber.Ctx, request entities.LoginRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	userName := request.UserName
	password := request.Password
	// check if exists
	exists, err := s.repository.ExistsActiveUser(ctx, userName)
	if err != nil {
		return err
	}
	if !exists {
		return apperrors.NewNotFoundError("User does not exists")
	}
	user, err := s.repository.FindByUserName(ctx, userName)
	if err != nil {
		return apperrors.NewNotFoundError(err.Error())
	}
	hashedPassword := user.Password
	// check password
	if !encryption.CheckPassword(hashedPassword, password) {
		return apperrors.NewUnauthorizedError("Invalid credentials")
	}
	// create session
	store := config.GetSessionStore()
	session, err := store.Get(fiberCtx)
	if err != nil {
		return apperrors.NewInternalServerError(err.Error())
	}
	session.Set("user_name", userName)
	session.Set("password", hashedPassword)
	if err := session.Save(); err != nil {
		return apperrors.NewInternalServerError(err.Error())
	}
	//Update Token
	user.Token = encryption.GenerateRandomToken()
	if err := s.repository.Update(ctx, user); err != nil {
		return apperrors.NewInternalServerError(err.Error())
	}
	return nil
}

func (s *AuthService) Logout(ctx context.Context, fiberCtx *fiber.Ctx) error {
	store := config.GetSessionStore()
	session, err := store.Get(fiberCtx)
	if err != nil {
		return apperrors.NewInternalServerError("Error while get session store: " +err.Error())
	}
	session.Destroy()
	if err := session.Save(); err != nil {
		return apperrors.NewInternalServerError("Error while destroying session: " +err.Error())
	} 
	return nil
}

func (s *AuthService) GetLoggedUser(ctx context.Context, fiberCtx *fiber.Ctx) (entities.UserData, error) {
	store := config.GetSessionStore()
	session, err := store.Get(fiberCtx)
	if err != nil {
		return entities.UserData{}, apperrors.NewInternalServerError("Error while get session store: " +err.Error())
	}
	userName := conversion.AnyToString(session.Get("user_name"))
	password := conversion.AnyToString(session.Get("password"))
	user, err := s.repository.GetByUserNameAndPassword(ctx, userName, password)
	if err != nil {
		return entities.UserData{}, apperrors.NewInternalServerError("User not found")
	}
	return user, nil
}

func (s *AuthService) IsUserAuthenticated(ctx context.Context, fiberCtx *fiber.Ctx) bool {
	loggedUser, err := s.GetLoggedUser(ctx, fiberCtx)
	if err != nil {
		return false
	}
	return loggedUser.UserId == 0 
}

func (s *AuthService) IsUserAdmin(ctx context.Context, fiberCtx *fiber.Ctx) bool {
	loggedUser, err := s.GetLoggedUser(ctx, fiberCtx)
	if err != nil {
		return false
	}
	return loggedUser.UserId != 0
}
