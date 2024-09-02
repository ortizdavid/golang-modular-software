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

type EmailConfigurationController struct {
	service           *services.EmailConfigurationService
	flagStatusService *services.ModuleFlagStatusService
	authService       *authentication.AuthService
	configService     *services.AppConfigurationService
	infoLogger        *helpers.Logger
	errorLogger       *helpers.Logger
}

func NewEmailConfigurationController(db *database.Database) *EmailConfigurationController {
	return &EmailConfigurationController{
		service:           services.NewEmailConfigurationService(db),
		flagStatusService: services.NewModuleFlagStatusService(db),
		authService:       authentication.NewAuthService(db),
		configService:     services.NewAppConfigurationService(db),
		infoLogger:        helpers.NewInfoLogger("configurations-info.log"),
		errorLogger:       helpers.NewErrorLogger("configurations-error.log"),
	}
}

func (ctrl *EmailConfigurationController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/configurations/email-configurations")
	group.Get("", ctrl.index)
	group.Get("/edit", ctrl.editForm)
	group.Post("/edit", ctrl.edit)
}

func (ctrl *EmailConfigurationController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("configuration/email/index", fiber.Map{
		"Title":            "Email Configurations",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *EmailConfigurationController) editForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("configuration/email/edit", fiber.Map{
		"Title":            "Edit Email Configuration",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
	})

}

func (ctrl *EmailConfigurationController) edit(c *fiber.Ctx) error {
	var request entities.UpdateEmailConfigurationRequest
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.UpdateEmailConfiguration(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' updated email configurations!", loggedUser.UserName))
	return c.Redirect("/configurations/email-configurations")
}
