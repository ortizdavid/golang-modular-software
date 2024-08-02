package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type RoleController struct {
	service *services.RoleService
	appConfig *configurations.AppConfiguration
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewRoleController(db *database.Database) *RoleController {
	return &RoleController{
		service:       services.NewRoleService(db),
		appConfig: 	configurations.LoadAppConfigurations(db),
		infoLogger:    helpers.NewInfoLogger("users-info.log"),
		errorLogger:   helpers.NewInfoLogger("users-error.log"),
	}
}

func (ctrl *RoleController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewAuthenticationMiddleware(db)
	group := router.Group("/roles", authMiddleware.CheckLoggedUser)
	group.Get("/", ctrl.index)
	group.Get("/create", ctrl.createForm)
}

func (ctrl *RoleController) index(c *fiber.Ctx) error {
	var params helpers.PaginationParam
	loggedUser := c.Locals("loggedUser").(entities.UserData)
	count, err := ctrl.service.CountRoles(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	roles, err := ctrl.service.GetAllRolesPaginated(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/roles/index", fiber.Map{
		"Title": "User Details",
		"Roles": roles,
		"Count": count,
		"LoggedUser":  loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *RoleController) createForm(c *fiber.Ctx) error {
	loggedUser := c.Locals("loggedUser").(entities.UserData)
	return c.Render("authentication/roles/create", fiber.Map{
		"Title":       "Create Role",
		"LoggedUser":  loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}