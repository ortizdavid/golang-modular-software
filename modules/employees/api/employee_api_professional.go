package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

func (api *EmployeeApi) getProfessionalInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	professionalInfo, err := api.professionalInfoService.GetByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(professionalInfo)
}

func (api *EmployeeApi) addProfessionalInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.CreateProfessionalInfoRequest
	request.EmployeeId = employee.EmployeeId
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.professionalInfoService.CreateProfessionalInfo(c.Context(), c, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Professional info for employee '"+employee.IdentificationNumber+"' added"
	api.infoLogger.Info(c, msg)
	return c.JSON(msg)
}

func (api *EmployeeApi) editProfessionalInfo(c *fiber.Ctx) error {
	empId := c.Params("empId")
	proId := c.Params("proId")
	employee, err := api.service.GetByUniqueId(c.Context(), empId)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	professionalInfo, err := api.professionalInfoService.GetByUniqueId(c.Context(), proId)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateProfessionalInfoRequest
	request.EmployeeId = employee.EmployeeId
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.professionalInfoService.UpdateProfessionalInfo(c.Context(), professionalInfo.ProfessionalId, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Professional info for employee '"+employee.IdentificationNumber+"' updated"
	api.infoLogger.Info(c, msg)
	return c.JSON(msg)
}
