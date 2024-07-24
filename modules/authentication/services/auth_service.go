package services

import (
	"context"
	"github.com/gofiber/fiber/v2"
	`github.com/ortizdavid/go-nopain/conversion`
	"github.com/ortizdavid/go-nopain/encryption"
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

func (s *AuthService) Authenticate(ctx context.Context, fiberCtx *fiber.Ctx, userName string, password string) error {
	user, err := s.repository.FindByUserName(ctx, userName)
	if err != nil {
		return err
	}
	// check if exists
	exists, err := s.repository.ExistsActiveUser(ctx, userName)
	if err != nil {
		return err
	}
	if !exists {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}
	hashedPassword := user.Password
	// check password
	if !encryption.CheckPassword(hashedPassword, password) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}
	// create session
	store := config.GetSessionStore()
	session, err := store.Get(fiberCtx)
	if err != nil {
		return err
	}
	session.Set("user_name", userName)
	session.Set("password", hashedPassword)
	if err := session.Save(); err != nil {
		return err
	}
	//Update Token
	user.Token = encryption.GenerateRandomToken()
	if err := s.repository.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) Logout(ctx context.Context, fiberCtx *fiber.Ctx) error {
	store := config.GetSessionStore()
	session, err := store.Get(fiberCtx)
	if err != nil {
		return err
	}
	session.Destroy()
	if err := session.Save(); err != nil {
		return err
	} 
	return nil
}

func (s *AuthService) GetLoggedUser(ctx context.Context, fiberCtx *fiber.Ctx) (entities.UserData, error) {
	store := config.GetSessionStore()
	session, err := store.Get(fiberCtx)
	if err != nil {
		return entities.UserData{}, err
	}
	userName := conversion.AnyToString(session.Get("user_name"))
	password := conversion.AnyToString(session.Get("password"))
	user, err := s.repository.GetByUserNameAndPassword(ctx, userName, password)
	return user, err
}

func (s *AuthService) IsUserAuthenticated(ctx context.Context, fiberCtx *fiber.Ctx) bool {
	loggedUser, err := s.GetLoggedUser(ctx, fiberCtx)
	if err != nil {
		return false
	}
	return loggedUser.UserId == 0 && loggedUser.RoleId == 0 
}

func (s *AuthService) IsUserAdmin(ctx context.Context,fiberCtx *fiber.Ctx) bool {
	loggedUser, err := s.GetLoggedUser(ctx, fiberCtx)
	if err != nil {
		return false
	}
	return loggedUser.RoleCode == "admin"
}
