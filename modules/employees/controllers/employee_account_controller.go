package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

func (ctrl *EmployeeController) addUserAccountForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	allowedRoles, err := ctrl.accountService.GetEmployeAllowedRoles(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employeeAccount, _ := ctrl.accountService.GetEmployeeAccountById(c.Context(), employee.EmployeeId)
	if employeeAccount.UserName != "" {
		return ctrl.HandleErrorsWeb(c, apperrors.NewConflictError("Employee '"+employee.FirstName+"' already have account"))
	}
	return c.Render("employee/employee-info/add-user-account", fiber.Map{
		"Title":            "Add Employee User Account",
		"Employee":         employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"AllowedRoles":     allowedRoles,
	})
}

func (ctrl *EmployeeController) addUserAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request authentication.CreateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.userService.CreateUser(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	userId := ctrl.userService.GetUserInsertedId()
	assRequest := authentication.AssociateUserRequest{
		UserId: userId,
		EntityId: employee.EmployeeId,
		EntityName: "employee",
	}
	err = ctrl.userService.AssociateUserToRole(c.Context(), assRequest)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.UpdateEmployeeUserId(c.Context(), employee.EmployeeId, userId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+loggedUser.UserName+"' added user account to employee '"+employee.IdentificationNumber+"' successfully")
	return c.Redirect("/employees/employee-info/" +id+ "/details")
}

func (ctrl *EmployeeController) associateUserAccountForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	availableUsers, err := ctrl.userService.GetUsersWithoutAssociation(c.Context(), ctrl.accountService.GetAllowedRoles())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/associate-user-account", fiber.Map{
		"Title":            "Associate to Existing Account",
		"Employee":         employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"AvailableUsers":   availableUsers,
	})
}

func (ctrl *EmployeeController) associateUserAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request authentication.AssociateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.userService.AssociateUserToRole(c.Context(), request)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.UpdateEmployeeUserId(c.Context(), employee.EmployeeId, request.UserId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+loggedUser.UserName+"' associated user account to employee '"+employee.IdentificationNumber+"' successfully")
	return c.Redirect("/employees/employee-info/" +id+ "/details")
}