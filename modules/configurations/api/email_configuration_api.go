package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	"gorm.io/gorm"
)

type EmailConfigurationApi struct {
	service     *services.EmailConfigurationService
	basicConfigService *services.BasicConfigurationService
	authService *authentication.AuthService
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewEmailConfigurationApi(db *gorm.DB) *EmailConfigurationApi {
	return &EmailConfigurationApi{
		service:     services.NewEmailConfigurationService(db),
		basicConfigService: services.NewBasicConfigurationService(db),
		authService: authentication.NewAuthService(db),
		infoLogger:  helpers.NewInfoLogger("configurations-info.log"),
		errorLogger: helpers.NewErrorLogger("configurations-error.log"),
	}
}

func (api *EmailConfigurationApi) Routes(router *fiber.App) {
	group := router.Group("/configurations/email-configurations")
	group.Put("", api.getEmailConfiguration)
	group.Put("", api.edit)
}

func (api *EmailConfigurationApi) getEmailConfiguration(c *fiber.Ctx) error {
	basicConfig, err := api.basicConfigService.GetBasicConfiguration(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggedUser, err := api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Render("configurations/email/index", fiber.Map{
		"Title": "Email Configurations",
		"BasicConfiguration": basicConfig,
		"LoggedUser":loggedUser,
	})
}

func (api *EmailConfigurationApi) edit(c *fiber.Ctx) error {
	var request entities.UpdateEmailConfigurationRequest
	loggedUser, err := api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid update email request")
	}
	err = api.service.UpdateEmailConfiguration(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	message := fmt.Sprintf("User '%s' updated email configurations!", loggedUser.UserName)
	api.infoLogger.Info(c, message)
	return c.JSON(message)
}
