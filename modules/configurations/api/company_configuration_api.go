package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type CompanyConfigurationApi struct {
	service     *services.CompanyConfigurationService
	authService *authentication.AuthService
	infoLogger  *helpers.Logger
	errorLogger *helpers.Logger
	shared.BaseController
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
	group := router.Group("/api/configurations/company-configurations")
	group.Get("", api.getCompanyConfiguration)
	group.Put("", api.edit)
}

func (api *CompanyConfigurationApi) getCompanyConfiguration(c *fiber.Ctx) error {
	_, err := api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	companyConfig, err := api.service.GetCurrent(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(companyConfig)
}

func (api *CompanyConfigurationApi) edit(c *fiber.Ctx) error {
	var request entities.UpdateCompanyConfigurationRequest
	loggedUser, err := api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.service.Update(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	message := fmt.Sprintf("User '%s' updated company configurations!", loggedUser.UserName)
	api.infoLogger.Info(c, message)
	return c.JSON(message)
}
