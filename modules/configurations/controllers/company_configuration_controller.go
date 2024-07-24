package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CompanyConfigurationController struct {
	service *services.CompanyConfigurationService
	basicConfigService *services.BasicConfigurationService
	authService *authentication.AuthService
	logger *zap.Logger
}

func NewCompanyConfigurationController(db *gorm.DB) *CompanyConfigurationController {
	return &CompanyConfigurationController{
		service:     services.NewCompanyConfigurationService(db),
		basicConfigService: services.NewBasicConfigurationService(db),
		authService: authentication.NewAuthService(db),
		logger: configurationLogger,
	}
}

func (ctrl *CompanyConfigurationController) Routes(router *fiber.App) {
	group := router.Group("/configurations/company-configurations")
	group.Get("", ctrl.index)
	group.Get("/edit", ctrl.editForm)
	group.Post("/edit", ctrl.edit)
}

func (ctrl *CompanyConfigurationController) index(c *fiber.Ctx) error {
	basicConfig, err := ctrl.basicConfigService.GetBasicConfiguration(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Render("configurations/company/index", fiber.Map{
		"Title": "Company Configurations",
		"BasicConfiguration": basicConfig,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *CompanyConfigurationController) editForm(c *fiber.Ctx) error {
	basicConfig, err := ctrl.basicConfigService.GetBasicConfiguration(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	companyConfig, err := ctrl.service.GetCompanyConfiguration(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
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
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid update company request")
	}
	err = ctrl.service.UpdateCompanyConfiguration(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	ctrl.logger.Info(fmt.Sprintf("User '%s' updated company configurations!", loggedUser.UserName), config.LogRequestPath(c))
	return c.Redirect("/email-configurations")
}
