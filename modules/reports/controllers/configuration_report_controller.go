package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type ConfigurationReportController struct {
	moduleFlagStatusService *configurations.ModuleFlagStatusService
	coreEntityFlagStatusService *configurations.CoreEntityFlagStatusService
	authService *authentication.AuthService
	configService *configurations.AppConfigurationService
}

func NewConfigurationReportController(db *database.Database) *ConfigurationReportController {
	return &ConfigurationReportController{
		authService: authentication.NewAuthService(db),
		moduleFlagStatusService:     configurations.NewModuleFlagStatusService(db),
		coreEntityFlagStatusService: configurations.NewCoreEntityFlagStatusService(db),
		configService: configurations.NewAppConfigurationService(db),
	}
}

func (ctrl *ConfigurationReportController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/reports/configurations")
	group.Get("/", ctrl.index)
}

func (ctrl *ConfigurationReportController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	coreEntityFlagStatus, _ := ctrl.coreEntityFlagStatusService.LoadCoreEntityFlagStatus(c.Context())
	return c.Render("reports/configuration/index", fiber.Map{
		"Title":      "Configuration Reports",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"CoreEntityFlagStatus": coreEntityFlagStatus,
	})
}
