package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	"github.com/ortizdavid/golang-modular-software/modules/references/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type MaritalStatusController struct {
	service                 *services.MaritalStatusService
	authService             *authentication.AuthService
	configService           *configurations.AppConfigurationService
	moduleFlagStatusService *configurations.ModuleFlagStatusService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
	shared.BaseController
}

func NewMaritalStatusController(db *database.Database) *MaritalStatusController {
	return &MaritalStatusController{
		service:                 services.NewMaritalStatusService(db),
		authService:             authentication.NewAuthService(db),
		configService:           configurations.NewAppConfigurationService(db),
		moduleFlagStatusService: configurations.NewModuleFlagStatusService(db),
		infoLogger:              helpers.NewInfoLogger(infoLogFile),
		errorLogger:             helpers.NewErrorLogger(errorLogFile),
	}
}

func (ctrl *MaritalStatusController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/references/marital-statuses")
	group.Get("", ctrl.index)
	group.Get("/:id/details", ctrl.details)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
	group.Get("/:id/delete", ctrl.removeForm)
	group.Post("/:id/delete", ctrl.remove)
}

func (ctrl *MaritalStatusController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("reference/marital-status/index", fiber.Map{
		"Title":            "Marital Statuses",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"Pagination":       pagination,
		"CurrentPage":      pagination.MetaData.CurrentPage + 1,
		"TotalPages":       pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *MaritalStatusController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	status, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("reference/marital-status/details", fiber.Map{
		"Title":            "Details",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"MaritalStatus":    status,
	})
}

func (ctrl *MaritalStatusController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("reference/marital-status/create", fiber.Map{
		"Title":            "Create Marital Status",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
	})
}

func (ctrl *MaritalStatusController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateStatusRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err := ctrl.service.Create(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Created marital status '"+request.StatusName+"' successfully")
	return c.Redirect("/references/marital-statuses")
}

func (ctrl *MaritalStatusController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	status, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("reference/marital-status/edit", fiber.Map{
		"Title":            "Edit Marital Status",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"MaritalStatus":    status,
	})
}

func (ctrl *MaritalStatusController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.UpdateStatusRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err := ctrl.service.Update(c.Context(), id, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Updated marital status '"+request.StatusName+"' successfully")
	return c.Redirect("/references/marital-statuses/" + id + "/details")
}

func (ctrl *MaritalStatusController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("reference/marital-status/search", fiber.Map{
		"Title":            "Search Statuses",
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *MaritalStatusController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchStatusRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.Search(c.Context(), c, request, params)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("reference/marital-status/search-results", fiber.Map{
		"Title":            "Search Results",
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"Pagination":       pagination,
		"Param":            request.SearchParam,
		"CurrentPage":      pagination.MetaData.CurrentPage + 1,
		"TotalPages":       pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *MaritalStatusController) removeForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	status, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("reference/marital-status/delete", fiber.Map{
		"Title":            "Remove Marital Status",
		"MaritalStatus":    status,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *MaritalStatusController) remove(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	status, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.Remove(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' removed marital status '%s'", loggedUser.UserName, status.StatusName))
	return c.Redirect("/references/marital-statuses")
}
