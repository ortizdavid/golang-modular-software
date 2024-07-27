package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type RoleController struct {
	service *services.RoleService
	authService *services.AuthService
	configService *configurations.BasicConfigurationService
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewRoleController(db *database.Database) *RoleController {
	return &RoleController{
		service:       services.NewRoleService(db),
		authService:   services.NewAuthService(db),
		configService: configurations.NewBasicConfigurationService(db),
		infoLogger:    helpers.NewInfoLogger("users-info.log"),
		errorLogger:   helpers.NewInfoLogger("users-error.log"),
	}
}

func (ctrl *RoleController) Routes(router *fiber.App) {
	group := router.Group("/roles")
	group.Get("/", ctrl.index)
	group.Get("/create", ctrl.createForm)
}

func (ctrl *RoleController) index(c *fiber.Ctx) error {
	var params helpers.PaginationParam
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
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
		"BasicConfig": basicConfig,
	})
}

func (ctrl *RoleController) createForm(c *fiber.Ctx) error {
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/roles/create", fiber.Map{
		"Title":       "Create Role",
		"LoggedUser":  loggedUser,
		"BasicConfig": basicConfig,
	})
}