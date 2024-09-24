package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
)

func (ctrl *UserController) deactivateForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	if loggedUser.UserId == user.UserId {
		return ctrl.HandleErrorsWeb(c, apperrors.NewConflictError("You cannot deactivate your own account"))
	}
	return c.Render("authentication/user/deactivate", fiber.Map{
		"Title":            "Deactivate User",
		"User":             user,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *UserController) deactivate(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	if loggedUser.UserId == user.UserId {
		return ctrl.HandleErrorsWeb(c, apperrors.NewConflictError("You cannot deactivate your own account"))
	}
	err = ctrl.service.DeactivateUser(c.Context(), user.UserId)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' deactivated successfully!", user.UserName))
	return c.Redirect("/user-management/users/" + id + "/details")
}

func (ctrl *UserController) activateForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	if loggedUser.UserId == user.UserId {
		return ctrl.HandleErrorsWeb(c, apperrors.NewConflictError("You cannot activate your own account"))
	}
	return c.Render("authentication/user/activate", fiber.Map{
		"Title":            "Activate User",
		"User":             user,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *UserController) activate(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	if loggedUser.UserId == user.UserId {
		return ctrl.HandleErrorsWeb(c, apperrors.NewConflictError("You cannot activate your own account"))
	}
	err = ctrl.service.ActivateUser(c.Context(), user.UserId)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' activated successfully!", user.UserName))
	return c.Redirect("/user-management/users/" + id + "/details")
}
