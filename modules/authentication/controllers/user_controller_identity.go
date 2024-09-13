package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

func (ctrl *UserController) resetPasswordForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	if loggedUser.UserId == user.UserId {
		return helpers.HandleHttpErrors(c, apperrors.NewConflictError("You cannot reset your own password"))
	}
	return c.Render("authentication/user/reset-password", fiber.Map{
		"Title":            "Reset Password",
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"User":             user,
	})
}

func (ctrl *UserController) resetPassword(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.ResetPasswordRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	if loggedUser.UserId == user.UserId {
		return helpers.HandleHttpErrors(c, apperrors.NewConflictError("You cannot reset your own password"))
	}
	err = ctrl.service.ResetUserPassword(c.Context(), user.UserId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' password reseted", loggedUser.UserName))
	return c.Redirect("/user-management/users/" + user.UniqueId + "/details")
}
