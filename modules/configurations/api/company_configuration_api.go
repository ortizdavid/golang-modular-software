package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type CompanyConfigurationApi struct {
	service     *services.CompanyConfigurationService
	basicConfigService *services.BasicConfigurationService
	authService *authentication.AuthService
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewCompanyConfigurationApi(db *database.Database) *CompanyConfigurationApi {
	return &CompanyConfigurationApi{
		service:     services.NewCompanyConfigurationService(db),
		basicConfigService: services.NewBasicConfigurationService(db),
		authService: authentication.NewAuthService(db),
		infoLogger:  helpers.NewInfoLogger("configurations-info.log"),
		errorLogger: helpers.NewErrorLogger("configurations-error.log"),
	}
}

func (api *CompanyConfigurationApi) Routes(router *fiber.App) {
	group := router.Group("/api/configurations/company-configurations")
	group.Get("", api.getCompanyConfiguration)
	group.Put("", api.edit)
}

func (api *CompanyConfigurationApi) getCompanyConfiguration(c *fiber.Ctx) error {
	basicConfig, err := api.basicConfigService.GetBasicConfiguration(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	_, err = api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(basicConfig)
}

func (api *CompanyConfigurationApi) edit(c *fiber.Ctx) error {
	var request entities.UpdateCompanyConfigurationRequest
	loggedUser, err := api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid update company request")
	}
	err = api.service.UpdateCompanyConfiguration(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	message := fmt.Sprintf("User '%s' updated company configurations!", loggedUser.UserName)
	api.infoLogger.Info(c, message)
	return c.JSON(message)
}
