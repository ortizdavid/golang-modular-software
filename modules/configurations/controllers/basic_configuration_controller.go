package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BasicConfigurationController struct {
	service *services.BasicConfigurationService
	authService *authentication.AuthService
	logger *zap.Logger
}

func NewBasicConfigurationController(db *gorm.DB) *BasicConfigurationController {
	return &BasicConfigurationController{
		service:     services.NewBasicConfigurationService(db),
		authService: authentication.NewAuthService(db),
		logger: configurationLogger,
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
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update basic configurations")
	}
	ctrl.logger.Info(fmt.Sprintf("User '%s' updated basic configurations!", loggedUser.UserName), config.LogRequestPath(c))
	return c.Redirect("/basic-configurations")
}