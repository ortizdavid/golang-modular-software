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
	permissionService *services.PermissionService
	appConfig *configurations.AppConfiguration
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewRoleController(db *database.Database) *RoleController {
	return &RoleController{
		service:           services.NewRoleService(db),
		authService:       services.NewAuthService(db),
		permissionService: services.NewPermissionService(db),
		appConfig:         configurations.LoadAppConfigurations(db),
		infoLogger:        helpers.NewInfoLogger("users-info.log"),
		errorLogger:       helpers.NewInfoLogger("users-error.log"),
	}
}

func (ctrl *RoleController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
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
	group.Get("/:id/assign-permission", ctrl.assignPermissionForm)
	group.Post("/:id/assign-permission", ctrl.assignPermission)
	group.Get("/:roleId/remove-permission/:permissionRoleId", ctrl.removePermissionForm)
	group.Post("/:roleId/remove-permission/:permissionRoleId", ctrl.removePermission)
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

func (ctrl *RoleController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	permissionRoles, _ := ctrl.permissionService.GetAssignedPermissionsByRole(c.Context(), role.RoleId)
	return c.Render("authentication/role/details", fiber.Map{
		"Title":       "Details",
		"LoggedUser":  loggedUser,
		"AppConfig": ctrl.appConfig,
		"Role": role,
		"PermissionRoles": permissionRoles,
		"CountPermissions": len(permissionRoles),
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

func (ctrl *RoleController) assignPermissionForm(c *fiber.Ctx) error {
	id := c.Params("id")
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	permissions, err := ctrl.permissionService.GetUnassignedPermissionsByRole(c.Context(), role.RoleId)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/role/assign-permission", fiber.Map{
		"Title": "Assign Role",
		"Permissions": permissions,
		"Role": role,
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *RoleController) assignPermission(c *fiber.Ctx) error {
	id := c.Params("id")
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.AssignRolePermissionRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.AssignRolePermission(c.Context(), role.RoleId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("Role '%s' assigned permission %d", role.RoleName, request.PermissionId))
	return c.Redirect("/roles/"+id+"/details")
}

func (ctrl *RoleController) removePermissionForm(c *fiber.Ctx) error {
	roleId := c.Params("roleId")
	permissionRoleId := c.Params("permissionRoleId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), roleId)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	permissionRole, err := ctrl.permissionService.GetPermissionRole(c.Context(), permissionRoleId)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/role/remove-permission", fiber.Map{
		"Title": "Remove Permission",
		"PermissionRole": permissionRole,
		"Role": role,
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *RoleController) removePermission(c *fiber.Ctx) error {
	roleId := c.Params("roleId")
	permissionRoleId := c.Params("permissionRoleId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), roleId)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	permissionRole, err := ctrl.permissionService.GetPermissionRole(c.Context(), permissionRoleId)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.permissionService.RemovePermissionRole(c.Context(), permissionRole.UniqueId)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' removed permission '%s' from role '%s'", loggedUser.UserName, permissionRole.PermissionName, role.RoleName))
	return c.Redirect("/roles/"+roleId+"/details")
}