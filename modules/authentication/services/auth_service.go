package services

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/config"
	entities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/repositories"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	"gorm.io/gorm"
)

type AuthService struct {
	repository *repositories.UserRepository
	emailService *configurations.EmailConfigurationService
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		repository: repositories.NewUserRepository(db),
		emailService: configurations.NewEmailConfigurationService(db),
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

func (s *AuthService) RecoverPassword(ctx context.Context, fiberCtx *fiber.Ctx, user entities.User, request entities.RecoverPasswordRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	user.Password = encryption.HashPassword(request.Password)
	err := s.repository.Update(ctx, user)
	if err != nil {
		return apperrors.NewInternalServerError("Error while recovering password: "+err.Error())
	}
	//send credentials
	defaultMail, err := s.emailService.GetDefaultEmailService(ctx)
	if err != nil {
		return apperrors.NewInternalServerError("Error while getting Email service: "+err.Error())
	}
	htmlBody := `
		<html>
			<body>
				<h1>Password Changed!</h1>
				<p>Hello, `+user.UserName+`!</p>
				<p>Your new password: <b>`+request.Password+`</b></p>
			</body>
		</html>`
	err = defaultMail.SendHTMLEmail(user.Email, "New Password", htmlBody)
	if err != nil {
		return apperrors.NewInternalServerError("Error while sending Email: "+err.Error())
	}
	return nil
}

func (s *AuthService) GetRecoverLink(ctx context.Context, fiberCtx *fiber.Ctx, user entities.User, request entities.GetRecoverLinkRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	//send credentials
	defaultMail, err := s.emailService.GetDefaultEmailService(ctx)
	if err != nil {
		return apperrors.NewInternalServerError("Error while getting Email service: "+err.Error())
	}
	recoverLink := fmt.Sprintf("%s/auth/recover-password/%s", fiberCtx.BaseURL(), user.Token)
	htmlBody := `
		<html>
			<body>
				<h1>Password Recovery!</h1>
				<p>Hello, `+user.UserName+`!</p>
				<p>To recover password Click <a href="`+recoverLink+`">Here</a></p>
			</body>
		</html>`
	err = defaultMail.SendHTMLEmail(user.Email, "Password Recovery", htmlBody)
	if err != nil {
		return apperrors.NewInternalServerError("Error while sending Email: "+err.Error())
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
