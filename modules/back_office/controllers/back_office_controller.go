package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type BackOfficeController struct {
	authService *authentication.AuthService
	configService *configurations.AppConfigurationService
    flagStatusService *configurations.ModuleFlagStatusService
}

func NewBackOfficeController(db *database.Database) *BackOfficeController {
	return &BackOfficeController{
		authService:       authentication.NewAuthService(db),
		configService:     configurations.NewAppConfigurationService(db),
		flagStatusService: configurations.NewModuleFlagStatusService(db),
	}
}

func (ctrl *BackOfficeController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/account")
	group.Get("/home",  ctrl.home)
	group.Get("/notifications", ctrl.notifications)
}

func (ctrl *BackOfficeController) home(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("_back_office/home", fiber.Map{
		"Title":      "Home",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *BackOfficeController) notifications(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("_back_office/notifications", fiber.Map{
		"Title":      "Notifications",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"LoggedUser": loggedUser,
	})
}
