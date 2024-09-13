package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	services "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type AuthController struct {
	service            *services.AuthService
	userService        *services.UserService
	configService      *configurations.BasicConfigurationService
	emailConfigService *configurations.EmailConfigurationService
	infoLogger         *helpers.Logger
	errorLogger        *helpers.Logger
	shared.BaseController
}

func NewAuthController(db *database.Database) *AuthController {
	return &AuthController{
		service:            services.NewAuthService(db),
		userService:        services.NewUserService(db),
		configService:      configurations.NewBasicConfigurationService(db),
		emailConfigService: configurations.NewEmailConfigurationService(db),
		infoLogger:         helpers.NewInfoLogger("auth-info.log"),
		errorLogger:        helpers.NewInfoLogger("auth-error.log"),
	}
}

func (auth AuthController) Routes(router *fiber.App) {
	group := router.Group("/auth")
	group.Get("/login", auth.loginForm)
	group.Post("/login", auth.login)
	group.Get("/logout", auth.logout)
	group.Get("/recover-password/:token", auth.recoverPasswordForm)
	group.Post("/recover-password/:token", auth.recoverPassword)
	group.Get("/get-recover-link", auth.getRecoverLinkForm)
	group.Post("/get-recover-link", auth.getRecoverLink)
}

func (ctrl *AuthController) loginForm(c *fiber.Ctx) error {
	return c.Render("authentication/auth/login", fiber.Map{
		"Title": "Login",
	})
}

func (ctrl *AuthController) login(c *fiber.Ctx) error {
	var request entities.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err := ctrl.service.Authenticate(c.Context(), c, request)
	if err != nil {
		ctrl.errorLogger.Error(c, fmt.Sprintf("User '%s' failed to login", request.UserName))
		return c.Status(fiber.StatusUnauthorized).Redirect("/auth/login")
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' authenticated sucessful!", request.UserName))
	return c.Status(fiber.StatusOK).Redirect("/account/home")
}

func (ctrl *AuthController) logout(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.service.GetLoggedUser(c.Context(), c)
	err := ctrl.service.Logout(c.Context(), c)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' logged out", loggedUser.UserName))
	return c.Redirect("/auth/login")
}

func (ctrl *AuthController) recoverPasswordForm(c *fiber.Ctx) error {
	token := c.Params("token")
	user, err := ctrl.userService.GetUserByToken(c.Context(), token)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("authentication/auth/recover-password", fiber.Map{
		"Title": "Password Recover",
		"User":  user,
	})
}

func (ctrl *AuthController) recoverPassword(c *fiber.Ctx) error {
	token := c.Params("token")
	var request entities.RecoverPasswordRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	user, err := ctrl.userService.GetUserByToken(c.Context(), token)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.RecoverPassword(c.Context(), c, user, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' recovered password", user.UserName))
	return c.Redirect("/auth/login")
}

func (ctrl *AuthController) getRecoverLinkForm(c *fiber.Ctx) error {
	return c.Render("authentication/auth/get-recover-link", fiber.Map{
		"Title": "Get Recovery Link",
	})
}

func (ctrl *AuthController) getRecoverLink(c *fiber.Ctx) error {
	var request entities.GetRecoverLinkRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	user, err := ctrl.userService.GetUserByEmail(c.Context(), request.Email)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.GetRecoverLink(c.Context(), c, user, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' recovered password", user.UserName))
	return c.Redirect("/auth/get-recover-link")
}
