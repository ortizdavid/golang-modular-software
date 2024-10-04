package api

import (
	"fmt"
	"slices"

	"github.com/gofiber/fiber/v2"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

func (api *EmployeeApi) getUserAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	account, err := api.accountService.GetByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(account)
}

func (api *EmployeeApi) addUserAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request authentication.CreateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	if !slices.Contains(api.accountService.AllowedRolesIdList(), request.RoleId) {
		return api.HandleErrorsApi(c, fmt.Errorf("role id '%d' not allowed", request.RoleId))
	}
	err = api.userService.CreateUser(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	userId := api.userService.GetUserInsertedId()
	assRequest := authentication.AssociateUserRequest{
		UserId: userId,
		EntityId: employee.EmployeeId,
		EntityName: "employee",
	}
	err = api.userService.AssociateUserToRole(c.Context(), assRequest)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.service.UpdateUserId(c.Context(), employee.EmployeeId, userId)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	msg := "Added account to employee '"+employee.IdentificationNumber+"' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(msg)
}

func (api *EmployeeApi) associateUserAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request authentication.AssociateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	if !slices.Contains(api.accountService.AllowedRolesList(), request.EntityName) {
		return api.HandleErrorsApi(c, fmt.Errorf("role code '%s' not allowed", request.EntityName))
	}
	err = api.userService.AssociateUserToRole(c.Context(), request)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.service.UpdateUserId(c.Context(), employee.EmployeeId, request.UserId)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	msg := "Employee '"+employee.IdentificationNumber+"' associated to an account successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(msg)
}

