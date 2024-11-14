package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type RoomApi struct {
	service 	*services.RoomService
	infoLogger  *helpers.Logger
	errorLogger  *helpers.Logger
	shared.BaseController
}

func NewRoomApi(db *database.Database) *RoomApi {
	return &RoomApi{
		service: 		services.NewRoomService(db),
		infoLogger:   helpers.NewInfoLogger(infoLogFile),
		errorLogger:  helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *RoomApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/company/rooms")
	group.Get("", api.getAll)
	group.Get("/not-paginated", api.getAllNotPaginated)
	group.Get("/:id", api.getByUniqueId)
	group.Post("", api.create)
	group.Put("/:id", api.edit)
}

func (api *RoomApi) getAll(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	rooms, err := api.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(rooms)
}

func (api *RoomApi) getAllNotPaginated(c *fiber.Ctx) error {
	rooms, err := api.service.GetAll(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(rooms)
}

func (api *RoomApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	company, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(company)
}

func (api *RoomApi) create(c *fiber.Ctx) error {
	var request entities.CreateRoomRequest
	if err := c.BodyParser(&request); err != nil {
		return  api.HandleErrorsApi(c, err)
	}
	err := api.service.Create(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return  api.HandleErrorsApi(c, err)
	}
	msg := "Created Room '"+request.RoomName+"' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"msg": msg})
}

func (api *RoomApi) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	room, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateRoomRequest
	if err := c.BodyParser(&request); err != nil {
		return  api.HandleErrorsApi(c, err)
	}
	err =  api.service.Update(c.Context(), id, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return  api.HandleErrorsApi(c, err)
	}
	msg := "Updated Room '"+room.RoomName+"' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"msg": msg})
}