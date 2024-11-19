package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type RoleApi struct {
	service     *services.RoleService
	infoLogger  *helpers.Logger
	errorLogger *helpers.Logger
	shared.BaseController
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
	group.Get("/", api.getAll)
	group.Get("/not-paginated", api.getAllNotPaginated)
	group.Post("/", api.create)
	group.Put("/:id", api.edit)
	group.Get("/:id", api.getByUniqueId)
	group.Get("/by-code/:code", api.getByCode)
}

func (api *RoleApi) getAll(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	roles, err := api.service.GetAllRolesPaginated(c.Context(), c, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(roles)
}

func (api *RoleApi) getAllNotPaginated(c *fiber.Ctx) error {
	roles, err := api.service.GetAllRoles(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(roles)
}

func (api *RoleApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	role, err := api.service.GetRoleByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(role)
}

func (api *RoleApi) getByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	role, err := api.service.GetRoleByCode(c.Context(), code)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(role)
}

func (api *RoleApi) create(c *fiber.Ctx) error {
	var request entities.CreateRoleRequest
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err := api.service.CreateRole(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Role '"+request.RoleName+"' created successfully"
	api.infoLogger.Info(c, msg)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"msg": msg})
}

func (api *RoleApi) edit(c *fiber.Ctx) error {
	id := c.Params(("id"))
	var request entities.UpdateRoleRequest
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err := api.service.UpdateRole(c.Context(), id, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Role '"+request.RoleName+"' updated successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"msg": msg})
}
