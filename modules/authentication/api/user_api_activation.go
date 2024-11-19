package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
)

func (api *UserApi) deactivate(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := api.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	loggedUser, _ := api.authService.GetLoggedUser(c.Context(), c)
	if loggedUser.UserId == user.UserId {
		return api.HandleErrorsApi(c, apperrors.NewConflictError("You cannot deactivate your own account"))
	}
	err = api.service.DeactivateUser(c.Context(), user.UserId)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := fmt.Sprintf("User '%s' deactivated successfully!", user.UserName)
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}

func (api *UserApi) activate(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := api.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	loggedUser, _ := api.authService.GetLoggedUser(c.Context(), c)
	if loggedUser.UserId == user.UserId {
		return api.HandleErrorsApi(c, apperrors.NewConflictError("You cannot activate your own account"))
	}
	err = api.service.ActivateUser(c.Context(), user.UserId)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := fmt.Sprintf("User '%s' activated successfully!", user.UserName)
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}
