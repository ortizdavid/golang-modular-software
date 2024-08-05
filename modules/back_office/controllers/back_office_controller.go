package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type BackOfficeController struct {
	appConfig *configurations.AppConfiguration
	authService *authentication.AuthService
}

func NewBackOfficeController(db *database.Database) *BackOfficeController {
	return &BackOfficeController{
		appConfig: configurations.LoadAppConfigurations(db),
		authService:   authentication.NewAuthService(db),
	}
}

func (ctrl *BackOfficeController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewAuthenticationMiddleware(db)
	router.Get("/home", authMiddleware.CheckLoggedUser, ctrl.home)
	router.Get("/dashboard", authMiddleware.CheckLoggedUser, ctrl.dashboard)
	router.Get("/notifications", authMiddleware.CheckLoggedUser, ctrl.notifications)
	router.Get("/user-data", authMiddleware.CheckLoggedUser, ctrl.userData)
	router.Get("/edit-user-data", authMiddleware.CheckLoggedUser, ctrl.editUserData)
}

func (ctrl *BackOfficeController) home(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("_back_office/home", fiber.Map{
		"Title": "Home",
		"AppConfig": ctrl.appConfig,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *BackOfficeController) dashboard(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("_back_office/dashboard", fiber.Map{
		"Title": "Dashboard",
		"AppConfig": ctrl.appConfig,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *BackOfficeController) notifications(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("_back_office/notifications", fiber.Map{
		"Title": "Notifications",
		"AppConfig": ctrl.appConfig,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *BackOfficeController) userData(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("_back_office/user-data", fiber.Map{
		"Title": "User Data",
		"AppConfig": ctrl.appConfig,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *BackOfficeController) editUserData(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("_back_office/edit-user-data", fiber.Map{
		"Title": "Edit User Data",
		"AppConfig": ctrl.appConfig,
		"LoggedUser": loggedUser,
	})
}