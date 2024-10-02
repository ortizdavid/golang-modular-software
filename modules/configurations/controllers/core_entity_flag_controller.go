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
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type CoreEntityFlagController struct {
	service                     *services.CoreEntityFlagService
	moduleFlagStatusService     *services.ModuleFlagStatusService
	coreEntityFlagStatusService *services.CoreEntityFlagStatusService
	authService                 *authentication.AuthService
	configService               *services.AppConfigurationService
	infoLogger                  *helpers.Logger
	errorLogger                 *helpers.Logger
	shared.BaseController
}

func NewCoreEntityFlagController(db *database.Database) *CoreEntityFlagController {
	return &CoreEntityFlagController{
		service:                     services.NewCoreEntityFlagService(db),
		moduleFlagStatusService:     services.NewModuleFlagStatusService(db),
		coreEntityFlagStatusService: services.NewCoreEntityFlagStatusService(db),
		authService:                 authentication.NewAuthService(db),
		configService:               services.NewAppConfigurationService(db),
		infoLogger:                  helpers.NewInfoLogger(infoLogFile),
		errorLogger:                 helpers.NewErrorLogger(erroLogFile),
	}
}

func (ctrl *CoreEntityFlagController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/configurations/core-entity-flags")
	group.Get("", ctrl.index)
	group.Get("/manage", ctrl.manageForm)
	group.Post("/manage", ctrl.manage)
}

func (ctrl *CoreEntityFlagController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	coreEntityFlagStatus, _ := ctrl.coreEntityFlagStatusService.LoadCoreEntityFlagStatus(c.Context())
	coreEntityFlags, err := ctrl.service.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("configuration/core-entity-flag/index", fiber.Map{
		"Title":                "Core Entities Flags",
		"AppConfig":            ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":           loggedUser,
		"ModuleFlagStatus":     moduleFlagStatus,
		"CoreEntityFlagStatus": coreEntityFlagStatus,
		"CoreEntityFlags":      coreEntityFlags,
	})
}

func (ctrl *CoreEntityFlagController) manageForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	coreEntityFlags, err := ctrl.service.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("configuration/core-entity-flag/manage", fiber.Map{
		"Title":            "Manage Core Entity Flags",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"CoreEntityFlags":  coreEntityFlags,
	})
}

func (ctrl *CoreEntityFlagController) manage(c *fiber.Ctx) error {
	var requests []entities.ManageCoreEntityFlagRequest
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	// Iterate over the form data
	c.Request().PostArgs().VisitAll(func(key, value []byte) {
		keyStr := string(key)
		if strings.HasPrefix(keyStr, "flag_") {
			flagIdStr := strings.TrimPrefix(keyStr, "flag_")
			flagId := conversion.StringToInt(flagIdStr)
			// Create a new ManageCoreEntityFlagRequest for each flag
			req := entities.ManageCoreEntityFlagRequest{
				FlagId: flagId,
				Status: string(value),
			}
			requests = append(requests, req)
		}
	})
	// Call the service method to process the core entity flags
	err := ctrl.service.ManageCoreEntityFlags(c.Context(), requests)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' updated core entity flags!", loggedUser.UserName))
	return c.Redirect("/configurations/core-entity-flags")
}
