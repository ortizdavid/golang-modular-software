package api

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

type CompanyConfigurationApi struct {
	service     *services.CompanyConfigurationService
	authService *authentication.AuthService
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewCompanyConfigurationApi(db *database.Database) *CompanyConfigurationApi {
	return &CompanyConfigurationApi{
		service:     services.NewCompanyConfigurationService(db),
		authService: authentication.NewAuthService(db),
		infoLogger:  helpers.NewInfoLogger("configurations-info.log"),
		errorLogger: helpers.NewErrorLogger("configurations-error.log"),
	}
}

func (api *CompanyConfigurationApi) Routes(router *fiber.App, db *database.Database) {
	apiKeyMiddleware := middlewares.NewApiKeyMiddleware(db)
	group := router.Group("/api/configurations/company-configurations", apiKeyMiddleware.AllowRoles("super-admin"))
	group.Get("", api.getCompanyConfiguration)
	group.Put("", api.edit)
}

func (api *CompanyConfigurationApi) getCompanyConfiguration(c *fiber.Ctx) error {
	_, err := api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	companyConfig, err := api.service.GetCompanyConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	return c.JSON(companyConfig)
}

func (api *CompanyConfigurationApi) edit(c *fiber.Ctx) error {
	var request entities.UpdateCompanyConfigurationRequest
	loggedUser, err := api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	err = api.service.UpdateCompanyConfiguration(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrorsApi(c, err)
	}
	message := fmt.Sprintf("User '%s' updated company configurations!", loggedUser.UserName)
	api.infoLogger.Info(c, message)
	return c.JSON(message)
}
