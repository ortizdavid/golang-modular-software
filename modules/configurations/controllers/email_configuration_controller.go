package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authEntities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type EmailConfigurationController struct {
	service *services.EmailConfigurationService
	basicConfigService *services.BasicConfigurationService
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewEmailConfigurationController(db *database.Database) *EmailConfigurationController {
	return &EmailConfigurationController{
		service:            services.NewEmailConfigurationService(db),
		basicConfigService: services.NewBasicConfigurationService(db),
		infoLogger:  helpers.NewInfoLogger("configurations-info.log"),
		errorLogger: helpers.NewErrorLogger("configurations-error.log"),
	}
}

func (ctrl *EmailConfigurationController) Routes(router *fiber.App) {
	group := router.Group("/configurations/email-configurations")
	group.Get("", ctrl.index)
	group.Get("/edit", ctrl.editForm)
	group.Post("/edit", ctrl.edit)
}


func (ctrl *EmailConfigurationController) index(c *fiber.Ctx) error {
	loggedUser := c.Locals("loggedUser").(*authEntities.UserData)
	basicConfig, err := ctrl.basicConfigService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("configurations/email/index", fiber.Map{
		"Title": "Email Configurations",
		"BasicConfiguration": basicConfig,
		"LoggedUser":loggedUser,
	})
}

func (ctrl *EmailConfigurationController) editForm(c *fiber.Ctx) error {
	loggedUser := c.Locals("loggedUser").(*authEntities.UserData)
	basicConfig, err := ctrl.basicConfigService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	emailConfig, err := ctrl.service.GetEmailConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("configurations/email/edit", fiber.Map{
		"Title": "Edita EmailConfig de Email",
		"EmailConfiguration": emailConfig,
		"BasicConfiguration": basicConfig,
		"LoggedUser": loggedUser,
	})
	
}

func (ctrl *EmailConfigurationController) edit(c *fiber.Ctx) error {
	var request entities.UpdateEmailConfigurationRequest
	loggedUser := c.Locals("loggedUser").(*authEntities.UserData)
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.UpdateEmailConfiguration(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' updated email configurations!", loggedUser.UserName))
	return c.Redirect("/email-configurations")
}
