package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type RootController struct {
	authService             *authentication.AuthService
	configService           *configurations.AppConfigurationService
	moduleFlagStatusService *configurations.ModuleFlagStatusService
}

func NewRootController(db *database.Database) *RootController {
	return &RootController{
		authService:             authentication.NewAuthService(db),
		configService:           configurations.NewAppConfigurationService(db),
		moduleFlagStatusService: configurations.NewModuleFlagStatusService(db),
	}
}

func (ctrl *RootController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/employees")
	group.Get("/", ctrl.index)
}

func (ctrl *RootController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("employee/_root/index", fiber.Map{
		"Title":            "Employees Management",
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
	})
}
