package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
)

type AuthApi struct {
	service            *services.AuthService
	userService        *services.UserService
	infoLogger         *helpers.Logger
	errorLogger        *helpers.Logger
}

func NewAuthApi(db *database.Database) *AuthApi {
	return &AuthApi{
		service:       services.NewAuthService(db),
		userService:   services.NewUserService(db),
		infoLogger:    helpers.NewInfoLogger("auth-info.log"),
		errorLogger:   helpers.NewInfoLogger("auth-error.log"),
	}
}

func (ctrl *AuthApi) Routes(router *fiber.App) {
	group := router.Group("/api/auth")
	group.Post("/login", ctrl.login)
	group.Get("/logout", ctrl.logout)
}

func (ctrl *AuthApi) login(c *fiber.Ctx) error {
	var request entities.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	err := ctrl.service.Authenticate(c.Context(), c, request)
	if err != nil {
		ctrl.errorLogger.Error(c, fmt.Sprintf("User '%s' failed to login", request.UserName))
		return helpers.HandleHttpErrorsApi(c, err)
	}
	message := fmt.Sprintf("User '%s' authenticated sucessful!", request.UserName)
	ctrl.infoLogger.Info(c, message)
	return c.Status(fiber.StatusOK).JSON(message)
}

func (ctrl *AuthApi) logout(c *fiber.Ctx) error {
	loggedUser := c.Locals("loggedUser").(entities.UserData)
	err := ctrl.service.Logout(c.Context(), c)
	if err != nil {
		ctrl.errorLogger.Error(c, fmt.Sprintf("User '%s' failed to logout", loggedUser.UserName))
		return helpers.HandleHttpErrors(c, err)
	}
	message := fmt.Sprintf("User '%s' logged out", loggedUser.UserName)
	ctrl.infoLogger.Info(c, message)
	return c.Status(fiber.StatusOK).JSON(message)
}