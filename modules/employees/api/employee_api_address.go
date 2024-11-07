package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

func (api *EmployeeApi) getAddresses(c *fiber.Ctx) error {
	id := c.Params("id")
	addresses, err := api.addressService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(addresses)
}

func (api *EmployeeApi) addAddress(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.CreateAddressRequest
	request.EmployeeId = employee.EmployeeId
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.addressService.Create(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Employee '" + employee.IdentificationNumber + "' address added"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}

func (api *EmployeeApi) editAddress(c *fiber.Ctx) error {
	empId := c.Params("empId")
	addId := c.Params("addId")
	employee, err := api.service.GetByUniqueId(c.Context(), empId)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	address, err := api.addressService.GetByUniqueId(c.Context(), addId)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateAddressRequest
	request.EmployeeId = employee.EmployeeId
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.addressService.Update(c.Context(), address.AddressId, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Employee '" + employee.IdentificationNumber + "' address updated"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}
