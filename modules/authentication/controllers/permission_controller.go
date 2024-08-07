package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type PermissionController struct {
	service *services.PermissionService
	authService *services.AuthService
	appConfig *configurations.AppConfiguration
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewPermissionController(db *database.Database) *PermissionController {
	return &PermissionController{
		service:     services.NewPermissionService(db),
		authService: services.NewAuthService(db),
		appConfig:   configurations.LoadAppConfigurations(db),
		infoLogger:  helpers.NewInfoLogger("users-info.log"),
		errorLogger: helpers.NewInfoLogger("users-error.log"),
	}
}

func (ctrl *PermissionController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewAuthenticationMiddleware(db)
	group := router.Group("/permissions", authMiddleware.CheckLoggedUser)
	group.Get("/", ctrl.index)
	group.Get("/create", ctrl.createForm)
}

func (ctrl *PermissionController) index(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	count, err := ctrl.service.CountPermissions(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	permissions, err := ctrl.service.GetAllPermissionsPaginated(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/permission/index", fiber.Map{
		"Title": "Permissions",
		"Permissions": permissions,
		"Count": count,
		"LoggedUser":  loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *PermissionController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("authentication/permission/create", fiber.Map{
		"Title":       "Create Permission",
		"LoggedUser":  loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}