package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

func (ctrl *UserController) assignRoleForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	roles, err := ctrl.roleService.GetUnassignedRolesByUser(c.Context(), user.UserId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("authentication/user/assign-role", fiber.Map{
		"Title":            "Assign Role",
		"Roles":            roles,
		"User":             user,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *UserController) assignRole(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.AssignUserRoleRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.AssignUserRole(c.Context(), user.UserId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' assigned role %d", user.UserName, request.RoleId))
	return c.Redirect("/user-management/users/" + id + "/details")
}

func (ctrl *UserController) removeRoleForm(c *fiber.Ctx) error {
	userId := c.Params("userId")
	userRoleId := c.Params("userRoleId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), userId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	userRole, err := ctrl.roleService.GetUserRole(c.Context(), userRoleId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("authentication/user/remove-role", fiber.Map{
		"Title":            "Remove Role",
		"UserRole":         userRole,
		"User":             user,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *UserController) removeRole(c *fiber.Ctx) error {
	userId := c.Params("userId")
	userRoleId := c.Params("userRoleId")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), userId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	userRole, err := ctrl.roleService.GetUserRole(c.Context(), userRoleId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.RemoveUserRole(c.Context(), userRole.UniqueId)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' removed role %s", user.UserName, userRole.RoleName))
	return c.Redirect("/user-management/users/" + userId + "/details")
}
