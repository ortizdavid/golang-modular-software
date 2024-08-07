package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type AccountController struct {
	service *services.UserService
	authService *services.AuthService
	roleService *services.RoleService
	appConfig *configurations.AppConfiguration
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewAccountController(db *database.Database) *AccountController {
	return &AccountController{
		service:       services.NewUserService(db),
		authService:   services.NewAuthService(db),
		roleService:   services.NewRoleService(db),
		appConfig: 		configurations.LoadAppConfigurations(db),
		infoLogger:    helpers.NewInfoLogger("users-info.log"),
		errorLogger:   helpers.NewInfoLogger("users-error.log"),
	}
}
func (ctrl *AccountController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewAuthenticationMiddleware(db)
	group := router.Group("account", authMiddleware.CheckLoggedUser)
	group.Get("/user-data", authMiddleware.CheckLoggedUser, ctrl.userData)
	group.Get("/change-data", authMiddleware.CheckLoggedUser, ctrl.changeUserDataForm)
	group.Get("/upload-image", authMiddleware.CheckLoggedUser, ctrl.uploadUserImageForm)
	group.Post("/upload-image", authMiddleware.CheckLoggedUser, ctrl.uploadUserImage)
	group.Get("/change-password", authMiddleware.CheckLoggedUser, ctrl.changePasswordForm)
	group.Post("/change-password", authMiddleware.CheckLoggedUser, ctrl.changePassword)
}

func (ctrl *AccountController) uploadUserImageForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("authentication/account/upload-image", fiber.Map{
		"Title":      "Upload Image",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.appConfig,
	})
}

func (ctrl *AccountController) uploadUserImage(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	err:= ctrl.service.UploadUserImage(c.Context(), c, loggedUser.UserId)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' uploaded image", loggedUser.UserName))
	return c.Redirect("/account/user-data")
}

func (ctrl *AccountController) changePasswordForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("authentication/account/change-password", fiber.Map{
		"Title":  "Change Password",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.appConfig,
	})
}

func (ctrl *AccountController) changePassword(c *fiber.Ctx) error {
	var request entities.ChangePasswordRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	err := ctrl.service.ChangeUserPassword(c.Context(), loggedUser.UserId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' updated password", loggedUser.UserName))
	return c.Redirect("/auth/login")
}

func (ctrl *AccountController) userData(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("authentication/account/user-data", fiber.Map{
		"Title":      "User Data",
		"AppConfig":  ctrl.appConfig,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *AccountController) changeUserDataForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("authentication/account/change-data", fiber.Map{
		"Title":      "Change Data",
		"AppConfig":  ctrl.appConfig,
		"LoggedUser": loggedUser,
	})
}