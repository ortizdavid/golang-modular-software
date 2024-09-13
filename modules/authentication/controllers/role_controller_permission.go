package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

func (ctrl *RoleController) assignPermissionForm(c *fiber.Ctx) error {
	id := c.Params("id")
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	permissions, err := ctrl.permissionService.GetUnassignedPermissionsByRole(c.Context(), role.RoleId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("authentication/role/assign-permission", fiber.Map{
		"Title":            "Assign Role",
		"Permissions":      permissions,
		"Role":             role,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *RoleController) assignPermission(c *fiber.Ctx) error {
	id := c.Params("id")
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.AssignRolePermissionRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.AssignRolePermission(c.Context(), role.RoleId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("Role '%s' assigned permission %d", role.RoleName, request.PermissionId))
	return c.Redirect("/user-management/roles/" + id + "/details")
}

func (ctrl *RoleController) removePermissionForm(c *fiber.Ctx) error {
	roleId := c.Params("roleId")
	permissionRoleId := c.Params("permissionRoleId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), roleId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	permissionRole, err := ctrl.permissionService.GetPermissionRole(c.Context(), permissionRoleId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("authentication/role/remove-permission", fiber.Map{
		"Title":            "Remove Permission",
		"PermissionRole":   permissionRole,
		"Role":             role,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *RoleController) removePermission(c *fiber.Ctx) error {
	roleId := c.Params("roleId")
	permissionRoleId := c.Params("permissionRoleId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	role, err := ctrl.service.GetRoleByUniqueId(c.Context(), roleId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	permissionRole, err := ctrl.permissionService.GetPermissionRole(c.Context(), permissionRoleId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.permissionService.RemovePermissionRole(c.Context(), permissionRole.UniqueId)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' removed permission '%s' from role '%s'", loggedUser.UserName, permissionRole.PermissionName, role.RoleName))
	return c.Redirect("/user-management/roles/" + roleId + "/details")
}
