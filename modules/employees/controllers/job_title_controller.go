package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	"github.com/ortizdavid/golang-modular-software/modules/employees/services"
)

type JobTitleController struct {
	service                 *services.JobTitleService
	authService             *authentication.AuthService
	moduleFlagStatusService *configurations.ModuleFlagStatusService
	configService           *configurations.AppConfigurationService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
}

func NewJobTitleController(db *database.Database) *JobTitleController {
	return &JobTitleController{
		service:                 services.NewJobTitleService(db),
		authService:             authentication.NewAuthService(db),
		moduleFlagStatusService: configurations.NewModuleFlagStatusService(db),
		configService:           configurations.NewAppConfigurationService(db),
		infoLogger:              helpers.NewInfoLogger("company-info.log"),
		errorLogger:             helpers.NewErrorLogger("company-error.log"),
	}
}

func (ctrl *JobTitleController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/employees/job-titles")
	group.Get("", ctrl.index)
}

func (ctrl *JobTitleController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("employee/job-title/index", fiber.Map{
		"Title":            "Job Titles",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
	})
}