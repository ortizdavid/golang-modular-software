package services

import (
	"context"
	//"database/sql"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	entities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/repositories"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	//shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type AuthService struct {
	repository         *repositories.UserRepository
	emailService       *configurations.EmailConfigurationService
	loginActRepository *repositories.LoginActivityRepository
	jwtService         *JwtService
}

func NewAuthService(db *database.Database) *AuthService {
	return &AuthService{
		repository:         repositories.NewUserRepository(db),
		emailService:       configurations.NewEmailConfigurationService(db),
		loginActRepository: repositories.NewLoginActivityRepository(db),
		jwtService:         NewJwtService(config.JwtSecretKey()),
	}
}

func (s *AuthService) Authenticate(ctx context.Context, fiberCtx *fiber.Ctx, request entities.LoginRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	userName := request.UserName
	password := request.Password
	// check if exists
	exists, err := s.repository.ExistsActiveUser(ctx, userName)
	if err != nil {
		return apperrors.InternalServerError("Error checking user existence")
	}
	if !exists {
		return apperrors.NotFoundError("User does not exists")
	}
	user, err := s.repository.FindByUserName(ctx, userName)
	if err != nil {
		return apperrors.NotFoundError(err.Error())
	}
	hashedPassword := user.Password
	// check password
	if !encryption.CheckPassword(hashedPassword, password) {
		return apperrors.UnauthorizedError("Invalid User Name or Password")
	}
	// create session
	store := config.GetSessionStore()
	session, err := store.Get(fiberCtx)
	if err != nil {
		return apperrors.InternalServerError(err.Error())
	}
	session.Set("user_name", userName)
	session.Set("password", hashedPassword)
	if err := session.Save(); err != nil {
		return apperrors.InternalServerError(err.Error())
	}
	//Update Token
	user.Token = encryption.GenerateRandomToken()
	user.Status = entities.ActivityStatusOnline
	if err := s.repository.Update(ctx, user); err != nil {
		return apperrors.InternalServerError(err.Error())
	}

	loginAct, err := s.loginActRepository.FindByUserId(ctx, user.UserId)
	// Update the existing login activity record
	loginAct.Status = entities.ActivityStatusOnline
	loginAct.Host = fiberCtx.Hostname()
	loginAct.Browser = string(fiberCtx.Context().UserAgent())
	loginAct.LastLogin = time.Now().UTC()
	loginAct.TotalLogin = loginAct.TotalLogin + 1
	loginAct.IPAddress = fiberCtx.IP()
	loginAct.Device = fiberCtx.Get("Device")
	loginAct.Location = fiberCtx.Get("Location")
	loginAct.UpdatedAt = time.Now().UTC()
	if err := s.loginActRepository.Update(ctx, loginAct); err != nil {
		return apperrors.InternalServerError("Failed to update login activity: " + err.Error())
	}
	return nil
}

// Authenticate API
func (s *AuthService) AuthenticateAPI(ctx context.Context, request entities.LoginRequest) (string, error) {
	if err := request.Validate(); err != nil {
		return "", apperrors.BadRequestError(err.Error())
	}
	userName := request.UserName
	password := request.Password
	// check if exists
	exists, err := s.repository.ExistsActiveUser(ctx, userName)
	if err != nil {
		return "", apperrors.InternalServerError("Error checking user existence")
	}
	if !exists {
		return "", apperrors.NotFoundError("User does not exists")
	}
	user, err := s.repository.FindByUserName(ctx, userName)
	if err != nil {
		return "", apperrors.NotFoundError(err.Error())
	}
	hashedPassword := user.Password
	// check password
	if !encryption.CheckPassword(hashedPassword, password) {
		return "", apperrors.UnauthorizedError("Invalid User Name or Password")
	}
	// Generate JWT token
	token, err := s.jwtService.GenerateJwtToken(user.UserId)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Error generating token")
	}
	return token, nil
}

func (s *AuthService) Logout(ctx context.Context, fiberCtx *fiber.Ctx) error {
	store := config.GetSessionStore()
	session, err := store.Get(fiberCtx)
	if err != nil {
		return apperrors.InternalServerError("Error while get session store: " + err.Error())
	}
	loggedUser, err := s.GetLoggedUser(ctx, fiberCtx)
	if err != nil {
		return err
	}
	// Get user
	user, _ := s.repository.FindById(ctx, loggedUser.UserId)
	user.Status = entities.ActivityStatusOffline
	if err := s.repository.Update(ctx, user); err != nil { //Update status
		return apperrors.InternalServerError(err.Error())
	}
	//Update Login Activity--------------------------------------------------------
	loginAct, err := s.loginActRepository.FindByUserId(ctx, user.UserId)
	if err != nil {
		return apperrors.InternalServerError("User activity not found: " + err.Error())
	}
	loginAct.Status = entities.ActivityStatusOffline
	loginAct.Host = fiberCtx.Hostname()
	loginAct.Browser = string(fiberCtx.Context().UserAgent())
	loginAct.LastLogout = time.Now().UTC()
	loginAct.TotalLogout = loginAct.TotalLogout + 1
	loginAct.IPAddress = fiberCtx.IP()
	loginAct.Device = fiberCtx.Get("Device")
	loginAct.Location = fiberCtx.Get("Location")
	if err := s.loginActRepository.Update(ctx, loginAct); err != nil {
		return apperrors.InternalServerError(err.Error())
	}
	//----- Destroy session
	session.Destroy()
	if err := session.Save(); err != nil {
		return apperrors.InternalServerError("Error while destroying session: " + err.Error())
	}
	return nil
}

func (s *AuthService) RecoverPassword(ctx context.Context, fiberCtx *fiber.Ctx, user entities.User, request entities.RecoverPasswordRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	user.Password = encryption.HashPassword(request.Password)
	err := s.repository.Update(ctx, user)
	if err != nil {
		return apperrors.InternalServerError("Error while recovering password: " + err.Error())
	}
	//send credentials
	defaultMail, err := s.emailService.GetDefaultEmailService(ctx)
	if err != nil {
		return apperrors.InternalServerError("Error while getting Email service: " + err.Error())
	}
	htmlBody := helpers.RecoverPasswordTmpl(user.UserName, request.Password)
	err = defaultMail.SendHTMLEmail(user.Email, "New Password", htmlBody)
	if err != nil {
		return apperrors.InternalServerError("Error while sending Email: " + err.Error())
	}
	return nil
}

func (s *AuthService) GetRecoverLink(ctx context.Context, fiberCtx *fiber.Ctx, user entities.User, request entities.GetRecoverLinkRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	//send credentials
	defaultMail, err := s.emailService.GetDefaultEmailService(ctx)
	if err != nil {
		return apperrors.InternalServerError("Error while getting Email service: " + err.Error())
	}
	recoverLink := fmt.Sprintf("%s/auth/recover-password/%s", fiberCtx.BaseURL(), user.Token)
	htmlBody := helpers.RecoverLinkTmpl(user.UserName, recoverLink)
	err = defaultMail.SendHTMLEmail(user.Email, "Password Recovery", htmlBody)
	if err != nil {
		return apperrors.InternalServerError("Error while sending Email: " + err.Error())
	}
	return nil
}

func (s *AuthService) GetLoggedUser(ctx context.Context, fiberCtx *fiber.Ctx) (entities.UserData, error) {
	store := config.GetSessionStore()
	session, err := store.Get(fiberCtx)
	if err != nil {
		return entities.UserData{}, apperrors.InternalServerError("Error while get session store: " + err.Error())
	}
	userName := conversion.AnyToString(session.Get("user_name"))
	password := conversion.AnyToString(session.Get("password"))
	user, err := s.repository.GetByUserNameAndPassword(ctx, userName, password)
	if err != nil {
		return entities.UserData{}, apperrors.InternalServerError("User not found")
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
