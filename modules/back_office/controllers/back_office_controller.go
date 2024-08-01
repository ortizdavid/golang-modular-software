package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type BackOfficeController struct {
	configService *configurations.BasicConfigurationService
}

func NewBackOfficeController(db *database.Database) *BackOfficeController {
	return &BackOfficeController{
		configService: configurations.NewBasicConfigurationService(db),
	}
}

func (ctrl *BackOfficeController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewAuthenticationMiddleware(db)
	router.Get("/dashboard", authMiddleware.CheckLoggedUser, ctrl.dashboard)
	router.Get("/home", authMiddleware.CheckLoggedUser, ctrl.home)
}

func (ctrl *BackOfficeController) home(c *fiber.Ctx) error {
	loggedUser := c.Locals("loggedUser").(entities.UserData)
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
	loggedUser := c.Locals("loggedUser").(entities.UserData)
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