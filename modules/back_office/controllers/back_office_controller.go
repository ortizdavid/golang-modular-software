package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type BackOfficeController struct {
	authService             *authentication.AuthService
	configService           *configurations.AppConfigurationService
	moduleFlagStatusService *configurations.ModuleFlagStatusService
	coreEntityFlagStatusService *configurations.CoreEntityFlagStatusService
}

func NewBackOfficeController(db *database.Database) *BackOfficeController {
	return &BackOfficeController{
		authService:                 authentication.NewAuthService(db),
		configService:               configurations.NewAppConfigurationService(db),
		moduleFlagStatusService:     configurations.NewModuleFlagStatusService(db),
		coreEntityFlagStatusService: configurations.NewCoreEntityFlagStatusService(db),
	}
}

func (ctrl *BackOfficeController) Routes(router *fiber.App, db *database.Database) {
	router.Get("/", ctrl.index)
	
	group := router.Group("/account")
	group.Get("/home", ctrl.home)
	group.Get("/notifications", ctrl.notifications)
}

func (ctrl *BackOfficeController) index(c *fiber.Ctx) error {
	return c.Redirect("/account/home")
}

func (ctrl *BackOfficeController) home(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	coreEntityFlagStatus, _ := ctrl.coreEntityFlagStatusService.LoadCoreEntityFlagStatus(c.Context())
	fmt.Println(coreEntityFlagStatus)
	return c.Render("_back_office/home", fiber.Map{
		"Title":            "Home",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"CoreEntityFlagStatus": coreEntityFlagStatus,
	})
}

func (ctrl *BackOfficeController) notifications(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	coreEntityFlagStatus, _ := ctrl.coreEntityFlagStatusService.LoadCoreEntityFlagStatus(c.Context())
	return c.Render("_back_office/notifications", fiber.Map{
		"Title":            "Notifications",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"CoreEntityFlagStatus": coreEntityFlagStatus,
		"LoggedUser":       loggedUser,
	})
}
