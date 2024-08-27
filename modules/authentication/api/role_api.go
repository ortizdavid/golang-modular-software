package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
)

type RoleApi struct {
	service     *services.RoleService
	infoLogger  *helpers.Logger
	errorLogger *helpers.Logger
}

func NewRoleApi(db *database.Database) *RoleApi {
	return &RoleApi{
		service:     services.NewRoleService(db),
		infoLogger:  helpers.NewInfoLogger("users-info.log"),
		errorLogger: helpers.NewInfoLogger("users-error.log"),
	}
}

func (api *RoleApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/user-management/roles")
	group.Get("/", api.getAllRoles)
	group.Post("/", api.createRole)
}

func (api *RoleApi) getAllRoles(c *fiber.Ctx) error {
	var params helpers.PaginationParam
	//loggedUser := c.Locals("loggedUser").(entities.UserData)
	roles, err := api.service.GetAllRolesPaginated(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(roles)
}

func (api *RoleApi) createRole(c *fiber.Ctx) error {
	//loggedUser := c.Locals("loggedUser").(entities.UserData)
	return c.Status(fiber.StatusOK).JSON(nil)
}
