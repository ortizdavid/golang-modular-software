package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

func (api *EmployeeApi) getDocuments(c *fiber.Ctx) error {
	id := c.Params("id")
	documents, err := api.documentService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(documents)
}

func (api *EmployeeApi) addDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.CreateDocumentRequest
	request.EmployeeId = employee.EmployeeId
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.documentService.Create(c.Context(), c, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Document '" + request.DocumentName + "' for employee '" + employee.FirstName + "' added"
	api.infoLogger.Info(c, msg)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": msg})
}

func (api *EmployeeApi) editDocument(c *fiber.Ctx) error {
	empId := c.Params("empId")
	docId := c.Params("docId")
	employee, err := api.service.GetByUniqueId(c.Context(), empId)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	document, err := api.documentService.GetByUniqueId(c.Context(), docId)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateDocumentRequest
	request.EmployeeId = employee.EmployeeId
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.documentService.Update(c.Context(), document.DocumentId, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Document '" + request.DocumentName + "' info for employee '" + employee.FirstName + " edited "
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}
