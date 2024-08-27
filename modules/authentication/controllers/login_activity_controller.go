package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type LoginActivityController struct {
	service     *services.LoginActivityService
	authService *services.AuthService
	configService *configurations.AppConfigurationService
	flagStatusService *configurations.ModuleFlagStatusService
	infoLogger  *helpers.Logger
	errorLogger *helpers.Logger
}

func NewLoginActivityController(db *database.Database) *LoginActivityController {
	return &LoginActivityController{
		service:           services.NewLoginActivityService(db),
		authService:       services.NewAuthService(db),
		configService:     configurations.NewAppConfigurationService(db),
		flagStatusService: configurations.NewModuleFlagStatusService(db),
		infoLogger:        helpers.NewInfoLogger("auth-info.log"),
		errorLogger:       helpers.NewInfoLogger("auth-error.log"),
	}
}

func (ctrl *LoginActivityController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/user-management/login-activities", authMiddleware.CheckLoggedUser)
	group.Get("", ctrl.index)
	group.Get("/:id/details", ctrl.details)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
}

func (ctrl *LoginActivityController) index(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	pagination, err := ctrl.service.GetAllLoginActivities(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/login-activity/index", fiber.Map{
		"Title":       "Login Activities",
		"LoggedUser":  loggedUser,
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"Pagination":  pagination,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages":  pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *LoginActivityController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	loginActivity, err := ctrl.service.GetLoginActivityByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/login-activity/details", fiber.Map{
		"Title":         "Details",
		"LoggedUser":    loggedUser,
		"AppConfig":     ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"LoginActivity": loginActivity,
	})
}

func (ctrl *LoginActivityController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("authentication/login-activity/search", fiber.Map{
		"Title":      "Search Login Activities",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
	})
}

func (ctrl *LoginActivityController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchLoginActivityRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.SearchLoginActivities(c.Context(), c, request, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("authentication/login-activity/search-results", fiber.Map{
		"Title":       "Search Results",
		"Pagination":  pagination,
		"Param":       request.SearchParam,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages":  pagination.MetaData.TotalPages + 1,
		"LoggedUser":  loggedUser,
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
	})
}
