package api

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

type EmailConfigurationApi struct {
	service     *services.EmailConfigurationService
	basicConfigService *services.BasicConfigurationService
	authService *authentication.AuthService
	logger      *zap.Logger
}

func NewEmailConfigurationApi(db *gorm.DB) *EmailConfigurationApi {
	return &EmailConfigurationApi{
		service:     services.NewEmailConfigurationService(db),
		basicConfigService: services.NewBasicConfigurationService(db),
		authService: authentication.NewAuthService(db),
		logger: configurationLogger,
	}
}

func (api *EmailConfigurationApi) Routes(router *fiber.App) {
	group := router.Group("/configurations/email-configurations")
	group.Put("", api.edit)
}

func (api *EmailConfigurationApi) index(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	message := fmt.Sprintf("User '%s' updated email configurations!", loggedUser.UserName)
	api.logger.Info(message, config.LogRequestPath(c))
	return c.JSON(message)
}
