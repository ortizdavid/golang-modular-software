package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type PermissionController struct {
	service       *services.PermissionService
	authService   *services.AuthService
	configService *configurations.AppConfigurationService
	flagStatusService *configurations.ModuleFlagStatusService
	infoLogger    *helpers.Logger
	errorLogger   *helpers.Logger
}

func NewPermissionController(db *database.Database) *PermissionController {
	return &PermissionController{
		service:           services.NewPermissionService(db),
		authService:       services.NewAuthService(db),
		configService:     configurations.NewAppConfigurationService(db),
		flagStatusService: configurations.NewModuleFlagStatusService(db),
		infoLogger:        helpers.NewInfoLogger("users-info.log"),
		errorLogger:       helpers.NewInfoLogger("users-error.log"),
	}
}

func (ctrl *PermissionController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/user-management/permissions")
	group.Get("/", ctrl.index)
	group.Get("/:id/details", ctrl.details)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/:id/delete", ctrl.deleteForm)
	group.Post("/:id/delete", ctrl.delete)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
}

func (ctrl *PermissionController) index(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	pagination, err := ctrl.service.GetAllPermissionsPaginated(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/permission/index", fiber.Map{
		"Title":      "Permissions",
		"Pagination": pagination,
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
	})
}

func (ctrl *PermissionController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	permission, err := ctrl.service.GetPermissionByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/permission/details", fiber.Map{
		"Title":      "Details",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"Permission": permission,
	})
}

func (ctrl *PermissionController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("authentication/permission/create", fiber.Map{
		"Title":      "Create Permission",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
	})
}

func (ctrl *PermissionController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreatePermissionRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.CreatePermission(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+loggedUser.UserName+"' created permission "+request.PermissionName)
	return c.Redirect("/user-management/permissions")
}

func (ctrl *PermissionController) editForm(c *fiber.Ctx) error {
	id := c.Params(("id"))
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	permission, err := ctrl.service.GetPermissionByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/permission/edit", fiber.Map{
		"Title":      "Edit Permission",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"Permission": permission,
	})
}

func (ctrl *PermissionController) edit(c *fiber.Ctx) error {
	id := c.Params(("id"))
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.UpdatePermissionRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	permission, err := ctrl.service.GetPermissionByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.UpdatePermission(c.Context(), permission.PermissionId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+loggedUser.UserName+"' updated permission "+request.PermissionName)
	return c.Redirect("/user-management/permissions")
}

func (ctrl *PermissionController) deleteForm(c *fiber.Ctx) error {
	id := c.Params(("id"))
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	permission, err := ctrl.service.GetPermissionByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/permission/delete", fiber.Map{
		"Title":      "Delete Permission",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"Permission": permission,
	})
}

func (ctrl *PermissionController) delete(c *fiber.Ctx) error {
	id := c.Params(("id"))
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	permission, err := ctrl.service.GetPermissionByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.DeletePermission(c.Context(), permission.PermissionId)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+loggedUser.UserName+"' deleted permission "+permission.PermissionName)
	return c.Redirect("/user-management/permissions")
}

func (ctrl *PermissionController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("authentication/permission/search", fiber.Map{
		"Title":      "Search Permissions",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
	})
}

func (ctrl *PermissionController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchPermissionRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.SearchPermissions(c.Context(), c, request, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("authentication/permission/search-results", fiber.Map{
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
