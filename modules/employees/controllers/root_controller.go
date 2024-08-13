package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type RootController struct {
	authService *authentication.AuthService
	appConfig *configurations.AppConfiguration
}

func NewRootController(db *database.Database) *RootController {
	return &RootController{
		authService: authentication.NewAuthService(db),
		appConfig:   configurations.LoadAppConfigurations(db),
	}
}

func (ctrl *RootController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/employees", authMiddleware.CheckLoggedUser)
	group.Get("/", ctrl.index)
}

func (ctrl *RootController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("employee/_root/index", fiber.Map{
		"Title": "Employees Management",
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}
