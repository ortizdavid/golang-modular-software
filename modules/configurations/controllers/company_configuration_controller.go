package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type CompanyConfigurationController struct {
	service *services.CompanyConfigurationService
	basicConfigService *services.BasicConfigurationService
	authService *authentication.AuthService
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewCompanyConfigurationController(db *database.Database) *CompanyConfigurationController {
	return &CompanyConfigurationController{
		service:     services.NewCompanyConfigurationService(db),
		basicConfigService: services.NewBasicConfigurationService(db),
		authService: authentication.NewAuthService(db),
		infoLogger:  helpers.NewInfoLogger("configurations-info.log"),
		errorLogger: helpers.NewErrorLogger("configurations-error.log"),
	}
}

func (ctrl *CompanyConfigurationController) Routes(router *fiber.App) {
	group := router.Group("/configurations/company-configurations")
	group.Get("", ctrl.index)
	group.Get("/edit", ctrl.editForm)
	group.Post("/edit", ctrl.edit)
}

func (ctrl *CompanyConfigurationController) index(c *fiber.Ctx) error {
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.basicConfigService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("configurations/company/index", fiber.Map{
		"Title": "Company Configurations",
		"BasicConfiguration": basicConfig,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *CompanyConfigurationController) editForm(c *fiber.Ctx) error {
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.basicConfigService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	companyConfig, err := ctrl.service.GetCompanyConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("configuration/company/edit", fiber.Map{
		"Title": "Edit Company Configuarions",
		"CompanyConfiguration": companyConfig,
		"BasicConfiguration": basicConfig,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *CompanyConfigurationController) edit(c *fiber.Ctx) error {
	var request entities.UpdateCompanyConfigurationRequest
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.UpdateCompanyConfiguration(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' updated company configurations!", loggedUser.UserName))
	return c.Redirect("/email-configurations")
}
