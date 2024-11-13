package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type OfficeApi struct {
	service 	*services.OfficeService
	infoLogger  *helpers.Logger
	errorLogger  *helpers.Logger
	shared.BaseController
}

func NewOfficeApi(db *database.Database) *OfficeApi {
	return &OfficeApi{
		service: 		services.NewOfficeService(db),
		infoLogger:   helpers.NewInfoLogger(infoLogFile),
		errorLogger:  helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *OfficeApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/company/offices")
	group.Get("", api.getAll)
	group.Get("/not-paginated", api.getAllNotPaginated)
	group.Get("/:id", api.getByUniqueId)
	group.Post("", api.create)
	group.Put("/:id", api.edit)
}

func (api *OfficeApi) getAll(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	offices, err := api.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(offices)
}

func (api *OfficeApi) getAllNotPaginated(c *fiber.Ctx) error {
	offices, err := api.service.GetAll(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(offices)
}

func (api *OfficeApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	office, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(office)
}

func (api *OfficeApi) create(c *fiber.Ctx) error {
	var request entities.CreateOfficeRequest
	if err := c.BodyParser(&request); err != nil {
		return  api.HandleErrorsApi(c, err)
	}
	err := api.service.Create(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return  api.HandleErrorsApi(c, err)
	}
	msg := "Created Office '"+request.OfficeName+"' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"msg": msg})
}

func (api *OfficeApi) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	office, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateOfficeRequest
	if err := c.BodyParser(&request); err != nil {
		return  api.HandleErrorsApi(c, err)
	}
	err =  api.service.Update(c.Context(), id, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return  api.HandleErrorsApi(c, err)
	}
	msg := "Updated Office '"+office.OfficeName+"' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"msg": msg})
}