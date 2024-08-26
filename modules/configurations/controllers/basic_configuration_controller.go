package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type BasicConfigurationController struct {
	service *services.BasicConfigurationService
	flagStatusService *services.ModuleFlagStatusService
	authService *authentication.AuthService
	configService *services.AppConfigurationService
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewBasicConfigurationController(db *database.Database) *BasicConfigurationController {
	return &BasicConfigurationController{
		service:           services.NewBasicConfigurationService(db),
		flagStatusService: services.NewModuleFlagStatusService(db),
		authService:       authentication.NewAuthService(db),
		configService:     services.NewAppConfigurationService(db),
		infoLogger:        helpers.NewInfoLogger("configurations-info.log"),
		errorLogger:       helpers.NewErrorLogger("configurations-error.log"),
	}
}

func (ctrl *BasicConfigurationController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/configurations/basic-configurations", authMiddleware.CheckLoggedUser)
	group.Get("", ctrl.index)
	group.Get("/edit", ctrl.editForm)
	group.Post("/edit", ctrl.edit)
}

func (ctrl *BasicConfigurationController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("configuration/basic/index", fiber.Map{
		"Title": "Basic Configurations",
		"AppConfig": ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"ModuleFlagStatus": flagStatus,
	})
}

func (ctrl *BasicConfigurationController) editForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("configuration/basic/edit", fiber.Map{
		"Title": "Edit Basic Configuration",
		"AppConfig": ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":loggedUser,
		"ModuleFlagStatus": flagStatus,
	})
}

func (ctrl *BasicConfigurationController) edit(c *fiber.Ctx) error {
	var request entities.UpdateBasicConfigurationRequest
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.UpdateBasicConfiguration(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' updated basic configurations!", loggedUser.UserName))
	return c.Redirect("/configurations/basic-configurations")
}