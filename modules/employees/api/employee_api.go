package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type EmployeeApi struct {
	service					*services.EmployeeService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
	shared.BaseController
}

func NewEmployeeApi(db *database.Database) *EmployeeApi {
	return &EmployeeApi{
		service:        services.NewEmployeeService(db),
		infoLogger:     helpers.NewInfoLogger(infoLogFile),
		errorLogger:    helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *EmployeeApi) Routes(router *fiber.App) {
	group := router.Group("/api/employees/employee-info")
	group.Get("", api.getAll)
	group.Get("/:id", api.getByUniqueId)
	group.Get("/by-identification/:identNumber", api.getByIdentificationNumber)
}

func (api *EmployeeApi) getAll(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	employees, err := api.service.GetAllEmployeesPaginated(c.Context(), c, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(employees)
}

func (api *EmployeeApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := api.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(employee)
}

func (api *EmployeeApi) getByIdentificationNumber(c *fiber.Ctx) error {
	identNumber := c.Params("identNumber")
	employee, err := api.service.GetEmployeeByIdentificationNumber(c.Context(), identNumber)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(employee)
}