package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type ProjectApi struct {
	service 	*services.ProjectService
	infoLogger  *helpers.Logger
	errorLogger  *helpers.Logger
	shared.BaseController
}

func NewProjectApi(db *database.Database) *ProjectApi {
	return &ProjectApi{
		service: 		services.NewProjectService(db),
		infoLogger:   helpers.NewInfoLogger(infoLogFile),
		errorLogger:  helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *ProjectApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/company/projects")
	group.Get("", api.getAll)
	group.Get("/not-paginated", api.getAllNotPaginated)
	group.Get("/:id", api.getByUniqueId)
	group.Post("", api.create)
	group.Put("/:id", api.edit)
}

func (api *ProjectApi) getAll(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	projects, err := api.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(projects)
}

func (api *ProjectApi) getAllNotPaginated(c *fiber.Ctx) error {
	projects, err := api.service.GetAll(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(projects)
}

func (api *ProjectApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	company, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(company)
}

func (api *ProjectApi) create(c *fiber.Ctx) error {
	var request entities.CreateProjectRequest
	if err := c.BodyParser(&request); err != nil {
		return  api.HandleErrorsApi(c, err)
	}
	err := api.service.Create(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return  api.HandleErrorsApi(c, err)
	}
	msg := "Created Project '"+request.ProjectName+"' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"msg": msg})
}

func (api *ProjectApi) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	project, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateProjectRequest
	if err := c.BodyParser(&request); err != nil {
		return  api.HandleErrorsApi(c, err)
	}
	err =  api.service.Update(c.Context(), id, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return  api.HandleErrorsApi(c, err)
	}
	msg := "Updated Project '"+project.ProjectName+"' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"msg": msg})
}