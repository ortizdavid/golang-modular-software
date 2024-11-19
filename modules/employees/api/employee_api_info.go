package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

func (api *EmployeeApi) create(c *fiber.Ctx) error {
	var request entities.CreateEmployeeRequest
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err := api.service.Create(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := fmt.Sprintf("Employee '%s %s' created", request.FirstName, request.LastName)
	api.infoLogger.Info(c, msg)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": msg})
}

func (api *EmployeeApi) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	var request entities.UpdateEmployeeRequest
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err := api.service.Update(c.Context(), id, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := fmt.Sprintf("Employee '%s %s' editd", request.FirstName, request.LastName)
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}

func (api *EmployeeApi) getAll(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	employees, err := api.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(employees)
}

func (api *EmployeeApi) search(c *fiber.Ctx) error {
	searchParam := c.Query("param")
	params := helpers.GetPaginationParams(c)
	request := entities.SearchEmployeeRequest{
		SearchParam: searchParam,
	}
	results, err := api.service.Search(c.Context(), c, request, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(results)
}

func (api *EmployeeApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := api.completeDataService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(employee)
}

func (api *EmployeeApi) getByIdentificationNumber(c *fiber.Ctx) error {
	identNumber := c.Params("identNumber")
	employee, err := api.completeDataService.GetByIdentificationNumber(c.Context(), identNumber)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(employee)
}

func (api *EmployeeApi) getPersonalInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	personalInfo, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(personalInfo)
}
