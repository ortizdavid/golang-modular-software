package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type CompanyReportController struct {
	moduleFlagStatusService     *configurations.ModuleFlagStatusService
	coreEntityFlagStatusService *configurations.CoreEntityFlagStatusService
	authService                 *authentication.AuthService
	configService               *configurations.AppConfigurationService
	infoLogger					*helpers.Logger
	errorLogger					*helpers.Logger
}

func NewCompanyReportController(db *database.Database) *CompanyReportController {
	return &CompanyReportController{
		authService:                 authentication.NewAuthService(db),
		moduleFlagStatusService:     configurations.NewModuleFlagStatusService(db),
		coreEntityFlagStatusService: configurations.NewCoreEntityFlagStatusService(db),
		configService:               configurations.NewAppConfigurationService(db),
		infoLogger:                  helpers.NewInfoLogger(infoLogFile),
		errorLogger:                 helpers.NewErrorLogger(errorLogFile),
	}
}

func (ctrl *CompanyReportController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/reports/company")
	group.Get("/", ctrl.index)
}

func (ctrl *CompanyReportController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	coreEntityFlagStatus, _ := ctrl.coreEntityFlagStatusService.LoadCoreEntityFlagStatus(c.Context())
	return c.Render("reports/company/index", fiber.Map{
		"Title":                "Company Reports",
		"LoggedUser":           loggedUser,
		"AppConfig":            ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus":     moduleFlagStatus,
		"CoreEntityFlagStatus": coreEntityFlagStatus,
	})
}
