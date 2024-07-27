package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type BasicConfigurationController struct {
	service *services.BasicConfigurationService
	authService *authentication.AuthService
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewBasicConfigurationController(db *database.Database) *BasicConfigurationController {
	return &BasicConfigurationController{
		service:     services.NewBasicConfigurationService(db),
		authService: authentication.NewAuthService(db),
		infoLogger:  helpers.NewInfoLogger("configurations-info.log"),
		errorLogger: helpers.NewErrorLogger("configurations-error.log"),
	}
}


func (ctrl *BasicConfigurationController) Routes(router *fiber.App) {
	group := router.Group("/configurations/basic-configurations")
	group.Get("", ctrl.index)
	group.Get("/edit", ctrl.editForm)
	group.Post("/edit", ctrl.edit)
}

func (ctrl *BasicConfigurationController) index(c *fiber.Ctx) error {
	configuration, err := ctrl.service.GetBasicConfiguration(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Render("configurations/email/index", fiber.Map{
		"Title": "Email Configurations",
		"BasicConfiguration": configuration,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *BasicConfigurationController) editForm(c *fiber.Ctx) error {
	configuration, err := ctrl.service.GetBasicConfiguration(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Render("configurations/email/edit", fiber.Map{
		"Title": "Edit Basic Configuration",
		"BasicConfiguaration": configuration,
		"LoggedUser":loggedUser,
	})
}

func (ctrl *BasicConfigurationController) edit(c *fiber.Ctx) error {
	var request entities.UpdateBasicConfigurationRequest
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid update basic config request")
	}
	err = ctrl.service.UpdateBasicConfiguration(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update basic configurations")
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' updated basic configurations!", loggedUser.UserName))
	return c.Redirect("/basic-configurations")
}