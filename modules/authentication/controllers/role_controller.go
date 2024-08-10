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

type RoleController struct {
	service *services.RoleService
	authService *services.AuthService
	appConfig *configurations.AppConfiguration
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewRoleController(db *database.Database) *RoleController {
	return &RoleController{
		service:     services.NewRoleService(db),
		authService: services.NewAuthService(db),
		appConfig:   configurations.LoadAppConfigurations(db),
		infoLogger:  helpers.NewInfoLogger("users-info.log"),
		errorLogger: helpers.NewInfoLogger("users-error.log"),
	}
}

func (ctrl *RoleController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewAuthenticationMiddleware(db)
	group := router.Group("/roles", authMiddleware.CheckLoggedUser)
	group.Get("/", ctrl.index)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/details", ctrl.details)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/:id/delete", ctrl.deleteForm)
	group.Post("/:id/delete", ctrl.delete)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
}

func (ctrl *RoleController) index(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	pagination, err := ctrl.service.GetAllRolesPaginated(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/role/index", fiber.Map{
		"Title": "Roles",
		"Pagination": pagination,
		"LoggedUser":  loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *RoleController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("authentication/role/create", fiber.Map{
		"Title":       "Create Role",
		"LoggedUser":  loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *RoleController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateRoleRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.CreateRole(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+loggedUser.UserName+"' created role "+request.RoleName)
	return c.Redirect("/roles")
}

func (ctrl *RoleController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/role/details", fiber.Map{
		"Title":       "Details",
		"LoggedUser":  loggedUser,
		"AppConfig": ctrl.appConfig,
		"Role": role,
	})
}

func (ctrl *RoleController) editForm(c *fiber.Ctx) error {
	id := c.Params(("id"))
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/role/edit", fiber.Map{
		"Title":       "Edit Role",
		"LoggedUser":  loggedUser,
		"AppConfig": ctrl.appConfig,
		"Role": role,
	})
}

func (ctrl *RoleController) edit(c *fiber.Ctx) error {
	id := c.Params(("id"))
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.UpdateRoleRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.UpdateRole(c.Context(), role.RoleId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+loggedUser.UserName+"' updated role "+request.RoleName)
	return c.Redirect("/roles")
}

func (ctrl *RoleController) deleteForm(c *fiber.Ctx) error {
	id := c.Params(("id"))
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/role/delete", fiber.Map{
		"Title":       "Delete Role",
		"LoggedUser":  loggedUser,
		"AppConfig": ctrl.appConfig,
		"Role": role,
	})
}

func (ctrl *RoleController) delete(c *fiber.Ctx) error {
	id := c.Params(("id"))
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.DeleteRole(c.Context(), role.RoleId)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+loggedUser.UserName+"' deleted role "+role.RoleName)
	return c.Redirect("/roles")
}

func (ctrl *RoleController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("authentication/role/search", fiber.Map{
		"Title": "Search Roles",
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *RoleController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchRoleRequest{SearchParam: searcParam}
    loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
    params := helpers.GetPaginationParams(c)
    pagination, err := ctrl.service.SearchRoles(c.Context(), c, request, params)
    if err != nil {
        return helpers.HandleHttpErrors(c, err)
    }
    ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
    return c.Render("authentication/role/search-results", fiber.Map{
        "Title":        "Search Results",
        "Pagination":   pagination,
        "Param":        request.SearchParam,
        "CurrentPage":  pagination.MetaData.CurrentPage + 1,
        "TotalPages":   pagination.MetaData.TotalPages + 1,
        "LoggedUser":   loggedUser,
        "AppConfig":  ctrl.appConfig,
    })
}