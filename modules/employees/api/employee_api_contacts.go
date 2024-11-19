package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

func (api *EmployeeApi) getEmails(c *fiber.Ctx) error {
	id := c.Params("id")
	emails, err := api.emailService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(emails)
}

func (api *EmployeeApi) addEmail(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.CreateEmployeeEmailRequest
	request.EmployeeId = employee.EmployeeId
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.emailService.Create(c.Context(), c, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Employee '" + employee.IdentificationNumber + "' email added"
	api.infoLogger.Info(c, msg)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": msg})
}

func (api *EmployeeApi) editEmail(c *fiber.Ctx) error {
	empId := c.Params("empId")
	contId := c.Params("contId")
	employee, err := api.service.GetByUniqueId(c.Context(), empId)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	employeeEmail, err := api.emailService.GetByUniqueId(c.Context(), contId)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateEmployeeEmailRequest
	request.EmployeeId = employee.EmployeeId
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.emailService.Update(c.Context(), employeeEmail.EmailId, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Employee '" + employee.IdentificationNumber + "' email edited"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}

func (api *EmployeeApi) getPhones(c *fiber.Ctx) error {
	id := c.Params("id")
	phones, err := api.phoneService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(phones)
}

func (api *EmployeeApi) addPhone(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.CreateEmployeePhoneRequest
	request.EmployeeId = employee.EmployeeId
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.phoneService.Create(c.Context(), c, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Employee '" + employee.IdentificationNumber + "' phone added"
	api.infoLogger.Info(c, msg)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": msg})
}

func (api *EmployeeApi) editPhone(c *fiber.Ctx) error {
	empId := c.Params("empId")
	contId := c.Params("contId")
	employee, err := api.service.GetByUniqueId(c.Context(), empId)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	employeePhone, err := api.phoneService.GetByUniqueId(c.Context(), contId)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateEmployeePhoneRequest
	request.EmployeeId = employee.EmployeeId
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.phoneService.Update(c.Context(), employeePhone.PhoneId, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Employee '" + employee.IdentificationNumber + "' phone edited"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}
