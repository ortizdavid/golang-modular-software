package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

func (ctrl *UserApi) getAllUsers(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	users, err := ctrl.service.GetAllUsers(c.Context(), c, params)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(users)
}

func (ctrl *UserApi) getActiveUsers(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	users, err := ctrl.service.GetActiveUsers(c.Context(), c, params)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(users)
}

func (ctrl *UserApi) getInactiveUsers(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	users, err := ctrl.service.GetInactiveUsers(c.Context(), c, params)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(users)
}

func (ctrl *UserApi) getOfflineUsers(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	users, err := ctrl.service.GetOnlineUsers(c.Context(), c, params)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(users)
}

func (ctrl *UserApi) getOnlineUsers(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	users, err := ctrl.service.GetOnlineUsers(c.Context(), c, params)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(users)
}


func (ctrl *UserApi) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchUserRequest{SearchParam: searcParam}
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.SearchUsers(c.Context(), c, request, params)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	msg := fmt.Sprintf("Searched for '%v' and found %d results", request.SearchParam, pagination.MetaData.TotalItems)
	ctrl.infoLogger.Info(c, msg)
	return c.JSON(msg)
}
