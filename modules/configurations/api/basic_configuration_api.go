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

type BasicConfigurationApi struct {
	service     *services.BasicConfigurationService
	authService *authentication.AuthService
	logger      *zap.Logger
}

func NewBasicConfigurationApi(db *gorm.DB) *BasicConfigurationApi {
	return &BasicConfigurationApi{
		service:     services.NewBasicConfigurationService(db),
		authService: authentication.NewAuthService(db),
		logger: configurationLogger,
	}
}

func (api *BasicConfigurationApi) Routes(router *fiber.App) {
	group := router.Group("/api/configurations/basic-configurations")
	group.Get("", api.getBasicConfiguration)
	group.Put("", api.edit)
}

func (api *BasicConfigurationApi) getBasicConfiguration(c *fiber.Ctx) error {
	configuration, err := api.service.GetBasicConfiguration(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	_, err = api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(configuration)
}

func (api *BasicConfigurationApi) edit(c *fiber.Ctx) error {
	var request entities.UpdateBasicConfigurationRequest
	loggedUser, err := api.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid update basic config request")
	}
	err = api.service.UpdateBasicConfiguration(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update basic configurations")
	}
	message := fmt.Sprintf("User '%s' updated basic configurations!", loggedUser.UserName)
	api.logger.Info(message, config.LogRequestPath(c))
	return c.JSON(message)
}