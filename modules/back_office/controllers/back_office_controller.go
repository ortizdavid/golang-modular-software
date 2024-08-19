package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type BackOfficeController struct {
	authService *authentication.AuthService
	configService *configurations.AppConfigurationService
	
}

func NewBackOfficeController(db *database.Database) *BackOfficeController {
	return &BackOfficeController{
		authService: authentication.NewAuthService(db),
		configService: configurations.NewAppConfigurationService(db),
	}
}

func (ctrl *BackOfficeController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	router.Get("/home", authMiddleware.CheckLoggedUser, ctrl.home)
	router.Get("/notifications", authMiddleware.CheckLoggedUser, ctrl.notifications)
}

func (ctrl *BackOfficeController) home(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("_back_office/home", fiber.Map{
		"Title":      "Home",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
	})
}

func (ctrl *BackOfficeController) notifications(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("_back_office/notifications", fiber.Map{
		"Title":      "Notifications",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
	})
}
