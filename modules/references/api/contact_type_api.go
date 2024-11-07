package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	"github.com/ortizdavid/golang-modular-software/modules/references/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type ContactTypeApi struct {
	service     *services.ContactTypeService
	infoLogger  *helpers.Logger
	errorLogger *helpers.Logger
	shared.BaseController
}

func NewContactTypeApi(db *database.Database) *ContactTypeApi {
	return &ContactTypeApi{
		service:     services.NewContactTypeService(db),
		infoLogger:  helpers.NewInfoLogger(infoLogFile),
		errorLogger: helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *ContactTypeApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/references/contact-types")
	group.Get("", api.getAll)
	group.Post("", api.add)
	group.Put("/:id", api.edit)
	group.Delete("/:id", api.remove)
	group.Get("/by-uuid/:id", api.getByUniqueId)
	group.Get("/by-name/:name", api.getByName)
	group.Get("/by-code/:code", api.getByCode)
	group.Get("/not-paginated", api.getAllNotPaginated)
}

func (api *ContactTypeApi) getAll(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	types, err := api.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(types)
}

func (api *ContactTypeApi) add(c *fiber.Ctx) error {
	var request entities.CreateTypeRequest
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err := api.service.Create(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Created contact type '" + request.TypeName + "' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}

func (api *ContactTypeApi) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	var request entities.UpdateTypeRequest
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err := api.service.Update(c.Context(), id, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Updated contact type '" + request.TypeName + "' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}

func (ctrl *ContactTypeApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	iType, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(iType)
}

func (ctrl *ContactTypeApi) getByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	iType, err := ctrl.service.GetByCode(c.Context(), code)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(iType)
}

func (ctrl *ContactTypeApi) getByName(c *fiber.Ctx) error {
	name := c.Params("name")
	iType, err := ctrl.service.GetByName(c.Context(), name)
	if err != nil {
		return ctrl.HandleErrorsApi(c, err)
	}
	return c.JSON(iType)
}

func (api *ContactTypeApi) remove(c *fiber.Ctx) error {
	id := c.Params("id")
	identType, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.service.Remove(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	msg := "Removed contact type '" + identType.TypeName + "'"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}

func (api *ContactTypeApi) getAllNotPaginated(c *fiber.Ctx) error {
	types, err := api.service.GetAll(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(types)
}
