package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type DepartmentApi struct {
	service 	*services.DepartmentService
	infoLogger  *helpers.Logger
	errorLogger  *helpers.Logger
	shared.BaseController
}

func NewDepartmentApi(db *database.Database) *DepartmentApi {
	return &DepartmentApi{
		service: 		services.NewDepartmentService(db),
		infoLogger:   helpers.NewInfoLogger(infoLogFile),
		errorLogger:  helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *DepartmentApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/company/departments")
	group.Get("", api.getAll)
	group.Get("/not-paginated", api.getAllNotPaginated)
	group.Get("/:id", api.getByUniqueId)
	group.Post("", api.create)
	group.Put("/:id", api.edit)
}

func (api *DepartmentApi) getAll(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	departments, err := api.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(departments)
}

func (api *DepartmentApi) getAllNotPaginated(c *fiber.Ctx) error {
	departments, err := api.service.GetAll(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(departments)
}

func (api *DepartmentApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	department, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(department)
}

func (api *DepartmentApi) create(c *fiber.Ctx) error {
	var request entities.CreateDepartmentRequest
	if err := c.BodyParser(&request); err != nil {
		return  api.HandleErrorsApi(c, err)
	}
	err := api.service.Create(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return  api.HandleErrorsApi(c, err)
	}
	msg := "Created Department '"+request.DepartmentName+"' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"msg": msg})
}

func (api *DepartmentApi) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	department, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateDepartmentRequest
	if err := c.BodyParser(&request); err != nil {
		return  api.HandleErrorsApi(c, err)
	}
	err =  api.service.Update(c.Context(), id, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return  api.HandleErrorsApi(c, err)
	}
	msg := "Updated Department '"+department.DepartmentName+"' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"msg": msg})
}