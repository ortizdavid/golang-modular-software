package api

import "github.com/gofiber/fiber/v2"

func (ctrl *UserApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(user)
}

func (ctrl *UserApi) getByName(c *fiber.Ctx) error {
	name := c.Params("name")
	user, err := ctrl.service.GetUserByName(c.Context(), name)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(user)
}

func (ctrl *UserApi) getByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := ctrl.service.GetUserByEmail(c.Context(), email)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(user)
}

func (ctrl *UserApi) getByToken(c *fiber.Ctx) error {
	token := c.Params("token")
	user, err := ctrl.service.GetUserByToken(c.Context(), token)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(user)
}

func (ctrl *UserApi) getUserRoles(c *fiber.Ctx) error {
	id := c.Params("id")
	roles, err := ctrl.roleService.GetAssignedRolesByUserUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(roles)
}