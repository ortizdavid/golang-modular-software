package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

func (ctrl *UserApi) create(c *fiber.Ctx) error {
	var request entities.CreateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	err := ctrl.service.CreateUser(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsApi(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+request.UserName+"' added successfully")
	return c.Redirect("/user-management/users")
}

func (ctrl *UserApi) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	var request entities.UpdateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	err := ctrl.service.UpdateUser(c.Context(), id, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsApi(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+request.UserName+"' edited successfuly")
	return c.Redirect("/user-management/users")
}
