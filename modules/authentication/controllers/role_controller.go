package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type RoleController struct {
	service *services.RoleService
	authService *services.AuthService
	appConfig *configurations.AppConfiguration
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewRoleController(db *database.Database) *RoleController {
	return &RoleController{
		service:     services.NewRoleService(db),
		authService: services.NewAuthService(db),
		appConfig:   configurations.LoadAppConfigurations(db),
		infoLogger:  helpers.NewInfoLogger("users-info.log"),
		errorLogger: helpers.NewInfoLogger("users-error.log"),
	}
}

func (ctrl *RoleController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewAuthenticationMiddleware(db)
	group := router.Group("/roles", authMiddleware.CheckLoggedUser)
	group.Get("/", ctrl.index)
	group.Get("/create", ctrl.createForm)
}

func (ctrl *RoleController) index(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	pagination, err := ctrl.service.GetAllRolesPaginated(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/role/index", fiber.Map{
		"Title": "Roles",
		"Pagination": pagination,
		"LoggedUser":  loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *RoleController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("authentication/roles/create", fiber.Map{
		"Title":       "Create Role",
		"LoggedUser":  loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}