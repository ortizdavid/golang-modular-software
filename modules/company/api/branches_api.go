package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type BranchApi struct {
	service 	*services.BranchService
	infoLogger  *helpers.Logger
	errorLogger  *helpers.Logger
	shared.BaseController
}

func NewBranchApi(db *database.Database) *BranchApi {
	return &BranchApi{
		service: 		services.NewBranchService(db),
		infoLogger:   helpers.NewInfoLogger(infoLogFile),
		errorLogger:  helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *BranchApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/company/branches")
	group.Get("", api.getAll)
	group.Get("/not-paginated", api.getAllNotPaginated)
	group.Get("/:id", api.getByUniqueId)
	group.Post("", api.create)
	group.Put("/:id", api.edit)
}

func (api *BranchApi) getAll(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	branches, err := api.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(branches)
}

func (api *BranchApi) getAllNotPaginated(c *fiber.Ctx) error {
	branches, err := api.service.GetAll(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(branches)
}

func (api *BranchApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	company, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(company)
}

func (api *BranchApi) create(c *fiber.Ctx) error {
	var request entities.CreateBranchRequest
	if err := c.BodyParser(&request); err != nil {
		return  api.HandleErrorsApi(c, err)
	}
	err := api.service.Create(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return  api.HandleErrorsApi(c, err)
	}
	msg := "Created Branch '"+request.BranchName+"' successfully"
	api.infoLogger.Info(c, msg)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"msg": msg})
}

func (api *BranchApi) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	branch, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateBranchRequest
	if err := c.BodyParser(&request); err != nil {
		return  api.HandleErrorsApi(c, err)
	}
	err =  api.service.Update(c.Context(), id, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return  api.HandleErrorsApi(c, err)
	}
	msg := "Updated Branch '"+branch.BranchName+"' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"msg": msg})
}