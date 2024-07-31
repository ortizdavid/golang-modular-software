package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type BackOfficeController struct {
	configService *configurations.BasicConfigurationService
	authService *authentication.AuthService
}

func NewBackOfficeController(db *database.Database) *BackOfficeController {
	return &BackOfficeController{
		configService: configurations.NewBasicConfigurationService(db),
		authService: authentication.NewAuthService(db),
	}
}

func (ctrl *BackOfficeController) Routes(router *fiber.App) {
	router.Get("/dashboard", ctrl.dashboard)
	router.Get("/home", ctrl.home)
}

func (ctrl *BackOfficeController) home(c *fiber.Ctx) error {
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("_back_office/home", fiber.Map{
		"Title": "Home",
		"BasicConfiguration": basicConfig,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *BackOfficeController) dashboard(c *fiber.Ctx) error {
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("_back_office/dashboard", fiber.Map{
		"Title": "Dashboard",
		"BasicConfiguration": basicConfig,
		"LoggedUser": loggedUser,
	})
}