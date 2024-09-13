package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

func (ctrl *EmployeeController) addDocumentForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	documentTypes, err := ctrl.documentTypeService.GetAllDocumentTypes(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/add-document", fiber.Map{
		"Title":            "Add Employee Document",
		"Employee":         employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"DocumentTypes":    documentTypes,
	})
}

func (ctrl *EmployeeController) addDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.CreateDocumentRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.documentService.CreateDocument(c.Context(), c, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' added document for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/" + id + "/details")
}

func (ctrl *EmployeeController) editDocumentForm(c *fiber.Ctx) error {
	empId := c.Params("empId")
	docId := c.Params("docId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), empId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	document, err := ctrl.documentService.GetDocumentByUniqueId(c.Context(), docId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	documentTypes, err := ctrl.documentTypeService.GetAllDocumentTypes(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/edit-document", fiber.Map{
		"Title":            "Edit Document Info",
		"Employee":         employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"DocumentTypes":    documentTypes,
		"Document":         document,
	})
}

func (ctrl *EmployeeController) editDocument(c *fiber.Ctx) error {
	empId := c.Params("empId")
	docId := c.Params("docId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), empId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	document, err := ctrl.documentService.GetDocumentByUniqueId(c.Context(), docId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.UpdateDocumentRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.documentService.UpdateDocument(c.Context(), document.DocumentId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' edit document for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/" + empId + "/details")
}

func (ctrl *EmployeeController) displayDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	document, err := ctrl.documentService.GetDocumentByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	documentPath := config.UploadDocumentPath() + "/employees"

	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' displayed document '%s'", loggedUser.UserName, document.DocumentName))
	return c.SendFile(documentPath + "/" + document.FileName)
}
