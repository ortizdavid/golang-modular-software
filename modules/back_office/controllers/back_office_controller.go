package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
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

func (ctrl *BackOfficeController) Routes(router *fiber.App, db *database.Database) {
	router.Get("/home", ctrl.home)
	router.Get("/dashboard", ctrl.dashboard)
}

func (ctrl *BackOfficeController) home(c *fiber.Ctx) error {

	return nil
}

func (ctrl *BackOfficeController) dashboard(c *fiber.Ctx) error {
	return nil
}