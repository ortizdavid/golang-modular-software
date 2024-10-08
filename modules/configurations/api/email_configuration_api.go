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

type EmailConfigurationApi struct {
	service     *services.EmailConfigurationService
	authService *authentication.AuthService
	infoLogger  *helpers.Logger
	errorLogger *helpers.Logger
	shared.BaseController
}

func NewEmailConfigurationApi(db *database.Database) *EmailConfigurationApi {
	return &EmailConfigurationApi{
		service:     services.NewEmailConfigurationService(db),
		authService: authentication.NewAuthService(db),
		infoLogger:  helpers.NewInfoLogger("configurations-info.log"),
		errorLogger: helpers.NewErrorLogger("configurations-error.log"),
	}
}

func (api *EmailConfigurationApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/configurations/email-configurations")
	group.Get("", api.getEmailConfiguration)
	group.Put("", api.edit)
}

func (api *EmailConfigurationApi) getEmailConfiguration(c *fiber.Ctx) error {
	_, err := api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	emailConfig, err := api.service.GetCurrent(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(emailConfig)
}

func (api *EmailConfigurationApi) edit(c *fiber.Ctx) error {
	var request entities.UpdateEmailConfigurationRequest
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
	message := fmt.Sprintf("User '%s' updated email configurations!", loggedUser.UserName)
	api.infoLogger.Info(c, message)
	return c.JSON(message)
}
