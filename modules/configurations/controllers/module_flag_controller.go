package controllers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"

	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type ModuleFlagController struct {
	service *services.ModuleFlagService
	flagStatusService *services.ModuleFlagStatusService
	authService *authentication.AuthService
	configService *services.AppConfigurationService
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewModuleFlagController(db *database.Database) *ModuleFlagController {
	return &ModuleFlagController{
		service:       services.NewModuleFlagService(db),
		flagStatusService:    services.NewModuleFlagStatusService(db),
		authService:   authentication.NewAuthService(db),
		configService: services.NewAppConfigurationService(db),
		infoLogger:    helpers.NewInfoLogger("configurations-info.log"),
		errorLogger:   helpers.NewErrorLogger("configurations-error.log"),
	}
}

func (ctrl *ModuleFlagController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/configurations/module-flags")
	group.Get("", ctrl.index)
	group.Get("/manage", ctrl.manageForm)
	group.Post("/manage", ctrl.manage)
}

func (ctrl *ModuleFlagController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	
	moduleFlags, err := ctrl.service.GetAllModuleFlags(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("configuration/module-flag/index", fiber.Map{
		"Title": "Module Flags",
		"AppConfig": ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"ModuleFlagStatus": flagStatus,
		"ModuleFlags": moduleFlags,
	})
}

func (ctrl *ModuleFlagController) manageForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	moduleFlags, err := ctrl.service.GetAllModuleFlags(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("configuration/module-flag/manage", fiber.Map{
		"Title": "Module Flags",
		"AppConfig": ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":loggedUser,
		"ModuleFlagStatus": flagStatus,
		"ModuleFlags": moduleFlags,
	})
}

func (ctrl *ModuleFlagController) manage(c *fiber.Ctx) error {
	var requests []entities.ManageModuleFlagRequest
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	// Iterate over the form data
	c.Request().PostArgs().VisitAll(func(key, value []byte) {
		keyStr := string(key)
		if strings.HasPrefix(keyStr, "flag_") {
			flagIdStr := strings.TrimPrefix(keyStr, "flag_")
			flagId := conversion.StringToInt(flagIdStr)
			// Create a new ManageModuleFlagRequest for each flag
			req := entities.ManageModuleFlagRequest{
				FlagId: flagId,
				Status: string(value),
			}
			requests = append(requests, req)
		}
	})
	// Call the service method to process the module flags
    err := ctrl.service.ManageModuleFlags(c.Context(), requests); 
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
        return helpers.HandleHttpErrors(c, err)
    }
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' updated module flags!", loggedUser.UserName))
	return c.Redirect("/configurations/module-flags")
}

