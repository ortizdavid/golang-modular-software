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

type CompanyConfigurationController struct {
	service       *services.CompanyConfigurationService
	flagStatusService *services.ModuleFlagStatusService
	authService   *authentication.AuthService
	configService *services.AppConfigurationService
	infoLogger    *helpers.Logger
	errorLogger   *helpers.Logger
}

func NewCompanyConfigurationController(db *database.Database) *CompanyConfigurationController {
	return &CompanyConfigurationController{
		service:           services.NewCompanyConfigurationService(db),
		flagStatusService: services.NewModuleFlagStatusService(db),
		authService:       authentication.NewAuthService(db),
		configService:     services.NewAppConfigurationService(db),
		infoLogger:        helpers.NewInfoLogger("configurations-info.log"),
		errorLogger:       helpers.NewErrorLogger("configurations-error.log"),
	}
}

func (ctrl *CompanyConfigurationController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/configurations/company-configurations", authMiddleware.CheckLoggedUser)
	group.Get("", ctrl.index)
	group.Get("/edit", ctrl.editForm)
	group.Post("/edit", ctrl.edit)
}

func (ctrl *CompanyConfigurationController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("configuration/company/index", fiber.Map{
		"Title":      "Company Configurations",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"ModuleFlagStatus": flagStatus,
	})
}

func (ctrl *CompanyConfigurationController) editForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("configuration/company/edit", fiber.Map{
		"Title":      "Edit Company Configuarions",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
	})
}

func (ctrl *CompanyConfigurationController) edit(c *fiber.Ctx) error {
	var request entities.UpdateCompanyConfigurationRequest
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.UpdateCompanyConfiguration(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' updated company configurations!", loggedUser.UserName))
	return c.Redirect("/configurations/company-configurations")
}
