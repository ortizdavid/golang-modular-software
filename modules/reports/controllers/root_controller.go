package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type RootController struct {
	moduleFlagStatusService     *configurations.ModuleFlagStatusService
	coreEntityFlagStatusService *configurations.CoreEntityFlagStatusService
	authService                 *authentication.AuthService
	configService               *configurations.AppConfigurationService
}

func NewRootController(db *database.Database) *RootController {
	return &RootController{
		authService:                 authentication.NewAuthService(db),
		moduleFlagStatusService:     configurations.NewModuleFlagStatusService(db),
		coreEntityFlagStatusService: configurations.NewCoreEntityFlagStatusService(db),
		configService:               configurations.NewAppConfigurationService(db),
	}
}

func (ctrl *RootController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/reports")
	group.Get("/", ctrl.index)
}

func (ctrl *RootController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	coreEntityFlagStatus, _ := ctrl.coreEntityFlagStatusService.LoadCoreEntityFlagStatus(c.Context())
	return c.Render("reports/_root/index", fiber.Map{
		"Title":                "Reports and Analitics",
		"LoggedUser":           loggedUser,
		"AppConfig":            ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus":     moduleFlagStatus,
		"CoreEntityFlagStatus": coreEntityFlagStatus,
	})
}
