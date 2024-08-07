package services

import (
	"context"
	"database/sql"
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
)

type AuthService struct {
	repository *repositories.UserRepository
	emailService *configurations.EmailConfigurationService
	loginActRepository *repositories.LoginActivityRepository
}

func NewAuthService(db *database.Database) *AuthService {
	return &AuthService{
		repository: repositories.NewUserRepository(db),
		emailService: configurations.NewEmailConfigurationService(db),
		loginActRepository: repositories.NewLoginActivityRepository(db),
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
		// Update or insert login activity
	loginAct, err := s.loginActRepository.FindByUserId(ctx, user.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			// Create a new login activity record if not found
			loginAct = entities.LoginActivity{
				UserId:     user.UserId,
				Status:     entities.ActivityStatusOnline,
				Host:       fiberCtx.Hostname(),
				Browser:    string(fiberCtx.Context().UserAgent()),
				IPAddress:  fiberCtx.IP(),
				Device:     fiberCtx.Get("Device"),
				Location:   fiberCtx.Get("Location"),
				UniqueId:   encryption.GenerateUUID(),
				LastLogin:  time.Now().UTC(),
				CreatedAt:  time.Now().UTC(),
				UpdatedAt:  time.Now().UTC(),
			}
			if err := s.loginActRepository.Create(ctx, loginAct); err != nil {
				return apperrors.NewInternalServerError("Failed to create login activity: " + err.Error())
			}
		} else {
			return apperrors.NewInternalServerError("Failed to find login activity: " + err.Error())
		}
	} else {
		// Update the existing login activity record
		loginAct.Status = entities.ActivityStatusOnline
		loginAct.Host = fiberCtx.Hostname()
		loginAct.Browser = string(fiberCtx.Context().UserAgent())
		loginAct.LastLogin = time.Now().UTC()
		loginAct.IPAddress = fiberCtx.IP()
		loginAct.Device = fiberCtx.Get("Device")
		loginAct.Location = fiberCtx.Get("Location")
		loginAct.UpdatedAt = time.Now().UTC()
		if err := s.loginActRepository.Update(ctx, loginAct); err != nil {
			return apperrors.NewInternalServerError("Failed to update login activity: " + err.Error())
		}
	}
	return nil
}

func (s *AuthService) Logout(ctx context.Context, fiberCtx *fiber.Ctx) error {
	store := config.GetSessionStore()
	session, err := store.Get(fiberCtx)
	if err != nil {
		return apperrors.NewInternalServerError("Error while get session store: " +err.Error())
	}
	loggedUser, err := s.GetLoggedUser(ctx, fiberCtx)
	if err != nil {
		return err
	}
	// Get user
	user, _ := s.repository.FindById(ctx, loggedUser.UserId)
	//Update Login Activity--------------------------------------------------------
	loginAct, err := s.loginActRepository.FindByUserId(ctx, user.UserId)
	if err != nil {
		return apperrors.NewInternalServerError("User activity not found: "+err.Error())
	}
	loginAct.Status = entities.ActivityStatusOffline
	loginAct.Host = fiberCtx.Hostname()
	loginAct.Browser = string(fiberCtx.Context().UserAgent())
	loginAct.LastLogout = time.Now().UTC()
	loginAct.IPAddress = fiberCtx.IP() 
	loginAct.Device = fiberCtx.Get("Device") 
	loginAct.Location = fiberCtx.Get("Location") 
	if err := s.loginActRepository.Update(ctx, loginAct); err != nil {
		return apperrors.NewInternalServerError(err.Error())
	}
	//----- Destroy session
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
	htmlBody := helpers.RecoverPasswordTmpl(user.UserName, request.Password)
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
	htmlBody := helpers.RecoverLinkTmpl(user.UserName, recoverLink)
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
