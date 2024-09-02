package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type RootController struct {
	authService       *authentication.AuthService
	flagStatusService *services.ModuleFlagStatusService
	configService     *services.AppConfigurationService
}

func NewRootController(db *database.Database) *RootController {
	return &RootController{
		authService:       authentication.NewAuthService(db),
		flagStatusService: services.NewModuleFlagStatusService(db),
		configService:     services.NewAppConfigurationService(db),
	}
}

func (ctrl *RootController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/configurations", authMiddleware.CheckLoggedUser)
	group.Get("/", ctrl.index)
}

func (ctrl *RootController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("configuration/_root/index", fiber.Map{
		"Title":            "Configurations",
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}
