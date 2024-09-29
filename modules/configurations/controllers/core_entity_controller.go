package controllers

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

type CoreEntityController struct {
	service                 *services.CoreEntityService
	moduleService           *services.ModuleService
	moduleFlagStatusService *services.ModuleFlagStatusService
	authService             *authentication.AuthService
	configService           *services.AppConfigurationService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
	shared.BaseController
}

func NewCoreEntityController(db *database.Database) *CoreEntityController {
	return &CoreEntityController{
		service:                 services.NewCoreEntityService(db),
		moduleService:           services.NewModuleService(db),
		moduleFlagStatusService: services.NewModuleFlagStatusService(db),
		authService:             authentication.NewAuthService(db),
		configService:           services.NewAppConfigurationService(db),
		infoLogger:              helpers.NewInfoLogger(infoLogFile),
		errorLogger:             helpers.NewErrorLogger(erroLogFile),
	}
}

func (ctrl *CoreEntityController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/configurations/core-entities")
	group.Get("", ctrl.index)
	group.Get("/:id/details", ctrl.details)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
}

func (ctrl *CoreEntityController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllCoreEntitiesPaginated(c.Context(), c, params)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("configuration/core-entity/index", fiber.Map{
		"Title":            "Core Entities",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"Pagination":       pagination,
		"CurrentPage":      pagination.MetaData.CurrentPage + 1,
		"TotalPages":       pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *CoreEntityController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	coreEntity, err := ctrl.service.GetCoreEntityByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("configuration/core-entity/details", fiber.Map{
		"Title":            "Details",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"CoreEntity":       coreEntity,
	})
}

func (ctrl *CoreEntityController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	modules, err := ctrl.moduleService.GetAllModules(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("configuration/core-entity/create", fiber.Map{
		"Title":            "Create Core Entity",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"Modules":          modules,
	})
}

func (ctrl *CoreEntityController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateCoreEntityRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err := ctrl.service.CreateCoreEntity(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Created branch '"+request.EntityName+"' successfully")
	return c.Redirect("/configurations/core-entities")
}

func (ctrl *CoreEntityController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	coreEntity, err := ctrl.service.GetCoreEntityByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("configuration/core-entity/edit", fiber.Map{
		"Title":            "Edit Core Entity",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"CoreEntity":       coreEntity,
	})
}

func (ctrl *CoreEntityController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	coreEntity, err := ctrl.service.GetCoreEntityByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.UpdateCoreEntityRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.UpdateCoreEntity(c.Context(), coreEntity.EntityId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Updated CoreEntity '"+request.EntityName+"' successfully")
	return c.Redirect("/configurations/core-entities/" + id + "/details")
}

func (ctrl *CoreEntityController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("configuration/core-entity/search", fiber.Map{
		"Title":      "Search Core Entities",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"FlagStatus": moduleFlagStatus,
	})
}

func (ctrl *CoreEntityController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchCoreEntityRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.SearchCoreEntities(c.Context(), c, request, params)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("configuration/core-entity/search-results", fiber.Map{
		"Title":            "Search Results",
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"Pagination":       pagination,
		"Param":            request.SearchParam,
		"CurrentPage":      pagination.MetaData.CurrentPage + 1,
		"TotalPages":       pagination.MetaData.TotalPages + 1,
	})
}
