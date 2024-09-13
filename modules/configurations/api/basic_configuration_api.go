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

type BasicConfigurationApi struct {
	service     *services.BasicConfigurationService
	authService *authentication.AuthService
	infoLogger  *helpers.Logger
	errorLogger *helpers.Logger
	shared.BaseController
}

func NewBasicConfigurationApi(db *database.Database) *BasicConfigurationApi {
	return &BasicConfigurationApi{
		service:     services.NewBasicConfigurationService(db),
		authService: authentication.NewAuthService(db),
		infoLogger:  helpers.NewInfoLogger("configurations-info.log"),
		errorLogger: helpers.NewErrorLogger("configurations-error.log"),
	}
}

func (api *BasicConfigurationApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/configurations/basic-configurations")
	group.Get("", api.getBasicConfiguration)
	group.Put("", api.edit)
}

func (api *BasicConfigurationApi) getBasicConfiguration(c *fiber.Ctx) error {
	_, err := api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	configuration, err := api.service.GetBasicConfiguration(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(configuration)
}

func (api *BasicConfigurationApi) edit(c *fiber.Ctx) error {
	var request entities.UpdateBasicConfigurationRequest
	loggedUser, err := api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.service.UpdateBasicConfiguration(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	message := fmt.Sprintf("User '%s' updated basic configurations!", loggedUser.UserName)
	api.infoLogger.Info(c, message)
	return c.JSON(message)
}
