package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type EmployeeReportController struct {
	moduleFlagStatusService     *configurations.ModuleFlagStatusService
	coreEntityFlagStatusService *configurations.CoreEntityFlagStatusService
	authService                 *authentication.AuthService
	configService               *configurations.AppConfigurationService
}

func NewEmployeeReportController(db *database.Database) *EmployeeReportController {
	return &EmployeeReportController{
		authService:                 authentication.NewAuthService(db),
		moduleFlagStatusService:     configurations.NewModuleFlagStatusService(db),
		coreEntityFlagStatusService: configurations.NewCoreEntityFlagStatusService(db),
		configService:               configurations.NewAppConfigurationService(db),
	}
}

func (ctrl *EmployeeReportController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/reports/employees")
	group.Get("/", ctrl.index)
}

func (ctrl *EmployeeReportController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	coreEntityFlagStatus, _ := ctrl.coreEntityFlagStatusService.LoadCoreEntityFlagStatus(c.Context())
	return c.Render("reports/employee/index", fiber.Map{
		"Title":                "Employee Reports",
		"LoggedUser":           loggedUser,
		"AppConfig":            ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus":     moduleFlagStatus,
		"CoreEntityFlagStatus": coreEntityFlagStatus,
	})
}
