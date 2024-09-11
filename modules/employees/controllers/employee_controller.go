package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	company "github.com/ortizdavid/golang-modular-software/modules/company/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/services"
	references "github.com/ortizdavid/golang-modular-software/modules/references/services"
)

type EmployeeController struct {
	service                 *services.EmployeeService
	jobTitleService			*services.JobTitleService
	documentService         *services.DocumentService
	phoneService    *services.EmployeePhoneService
	emailService    *services.EmployeeEmailService
	documentTypeService     *services.DocumentTypeService
	identTypeService		*references.IdentificationTypeService
	countryService			*references.CountryService
	maritalStatusService	*references.MaritalStatusService
	employmentStatusService	*references.EmploymentStatusService
	contactTypeService	    *references.ContactTypeService
	departmentService		*company.DepartmentService
	authService             *authentication.AuthService
	moduleFlagStatusService *configurations.ModuleFlagStatusService
	configService           *configurations.AppConfigurationService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
}

func NewEmployeeController(db *database.Database) *EmployeeController {
	return &EmployeeController{
		service:                 services.NewEmployeeService(db),
		jobTitleService:         services.NewJobTitleService(db),
		documentService:         services.NewDocumentService(db),
		phoneService:    services.NewEmployeePhoneService(db),
		emailService:    services.NewEmployeeEmailService(db),
		documentTypeService:     services.NewDocumentTypeService(db),
		identTypeService:        references.NewIdentificationTypeService(db),
		countryService:          references.NewCountryService(db),
		maritalStatusService:    references.NewMaritalStatusService(db),
		employmentStatusService: references.NewEmploymentStatusService(db),
		contactTypeService:      references.NewContactTypeService(db),
		departmentService:       company.NewDepartmentService(db),
		authService:             authentication.NewAuthService(db),
		moduleFlagStatusService: configurations.NewModuleFlagStatusService(db),
		configService:           configurations.NewAppConfigurationService(db),
		infoLogger:              helpers.NewInfoLogger("employees-info.log"),
		errorLogger:             helpers.NewErrorLogger("employees-error.log"),
	}
}

func (ctrl *EmployeeController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/employees/employee-info")
	group.Get("", ctrl.index)
	group.Get("/:id/details", ctrl.details)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
	group.Get("/:id/delete", ctrl.removeForm)
	group.Post("/:id/delete", ctrl.remove)
	group.Get("/:id/add-phone", ctrl.addPhoneForm)
	group.Post("/:id/add-phone", ctrl.addPhone)
	group.Get("/:id/add-email", ctrl.addEmailForm)
	group.Post("/:id/add-email", ctrl.addEmail)
	group.Get("/:id/add-document", ctrl.addDocumentForm)
	group.Post("/:id/add-document", ctrl.addDocument)
	group.Get("/:empId/edit-document/:docId", ctrl.editDocumentForm)
	group.Post("/:empId/edit-document/:docId", ctrl.editDocument)
	group.Get("/display-document/:id", ctrl.displayDocument)

}

func (ctrl *EmployeeController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllEmployeesPaginated(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("employee/employee-info/index", fiber.Map{
		"Title":            "Employees",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"Pagination":       pagination,
		"CurrentPage":      pagination.MetaData.CurrentPage + 1,
		"TotalPages":       pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *EmployeeController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	employeeDocuments, err := ctrl.documentService.GetAllEmployeeDocuments(c.Context(), employee.EmployeeId)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	employeePhones, err := ctrl.phoneService.GetAllEmployeePhones(c.Context(), employee.EmployeeId)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	employeeEmails, err := ctrl.emailService.GetAllEmployeeEmails(c.Context(), employee.EmployeeId)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("employee/employee-info/details", fiber.Map{
		"Title":            "Details",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"Employee":      employee,
		"EmployeeDocuments": employeeDocuments,
		"CountDocuments": len(employeeDocuments),
		"EmployeePhones": employeePhones,
		"CountPhones": len(employeePhones),
		"EmployeeEmails": employeeEmails,
		"CountEmails": len(employeeEmails),
	})
}

func (ctrl *EmployeeController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	
	jobTitles, err := ctrl.jobTitleService.GetAllJobTitles(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	identificationTypes, err := ctrl.identTypeService.GetAllIdentificationTypes(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	countries, err := ctrl.countryService.GetAllCountries(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	maritalStatuses, err := ctrl.maritalStatusService.GetAllStatuses(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	employmentStatuses, err := ctrl.employmentStatusService.GetAllStatuses(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	departments, err := ctrl.departmentService.GetAllDepartments(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("employee/employee-info/create", fiber.Map{
		"Title": "Create Employee",
		"AppConfig": ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser": loggedUser,
		"JobTitles": jobTitles,
		"IdentificationTypes": identificationTypes,
		"Countries": countries,
		"MaritalStatuses": maritalStatuses,
		"EmploymentStatuses": employmentStatuses,
		"Departments": departments,
	})
}

func (ctrl *EmployeeController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateEmployeeRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.CreateEmployee(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Created employment '"+request.FirstName+"' successfully")
	return c.Redirect("/employees/employee-info")
}

func (ctrl *EmployeeController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	
	jobTitles, err := ctrl.jobTitleService.GetAllJobTitles(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	identificationTypes, err := ctrl.identTypeService.GetAllIdentificationTypes(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	countries, err := ctrl.countryService.GetAllCountries(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	maritalStatuses, err := ctrl.maritalStatusService.GetAllStatuses(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	employmentStatuses, err := ctrl.employmentStatusService.GetAllStatuses(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	departments, err := ctrl.departmentService.GetAllDepartments(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("employee/employee-info/edit", fiber.Map{
		"Title":  "Edit Employee",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser": loggedUser,
		"Employee": employee,
		"JobTitles": jobTitles,
		"IdentificationTypes": identificationTypes,
		"Countries": countries,
		"MaritalStatuses": maritalStatuses,
		"EmploymentStatuses": employmentStatuses,
		"Departments": departments,
	})
}

func (ctrl *EmployeeController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.UpdateEmployeeRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.UpdateEmployee(c.Context(), employee.EmployeeId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Updated employment '"+request.FirstName+"' successfully")
	return c.Redirect("/employees/employee-info/" + id + "/details")
}

func (ctrl *EmployeeController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("employee/employee-info/search", fiber.Map{
		"Title":            "Search Employees",
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *EmployeeController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchEmployeeRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.SearchEmployees(c.Context(), c, request, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("employee/employee-info/search-results", fiber.Map{
		"Title":            "Search Results",
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"Pagination":       pagination,
		"Param":            request.SearchParam,
		"CurrentPage":      pagination.MetaData.CurrentPage + 1,
		"TotalPages":       pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *EmployeeController) removeForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("employee/employee-info/delete", fiber.Map{
		"Title":            "Remove Employee",
		"Employee":      employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *EmployeeController) remove(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.RemoveEmployee(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' removed employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info")
}


func (ctrl *EmployeeController) addPhoneForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	contactTypes, err := ctrl.contactTypeService.GetAllContactTypes(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("employee/employee-info/add-phone", fiber.Map{
		"Title":            "Add Employee Phone",
		"Employee":      employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"ContactTypes": contactTypes,
	})
}

func (ctrl *EmployeeController) addPhone(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.CreateEmployeePhoneRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.phoneService.CreateEmployeePhone(c.Context(), c, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' added phone for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/"+id+"/details")
}

func (ctrl *EmployeeController) addEmailForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	contactTypes, err := ctrl.contactTypeService.GetAllContactTypes(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("employee/employee-info/add-email", fiber.Map{
		"Title":            "Add Employee Email",
		"Employee":      employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"ContactTypes": contactTypes,
	})
}

func (ctrl *EmployeeController) addEmail(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.CreateEmployeeEmailRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.emailService.CreateEmployeeEmail(c.Context(), c, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' added email for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/"+id+"/details")
}

func (ctrl *EmployeeController) addDocumentForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	documentTypes, err := ctrl.documentTypeService.GetAllDocumentTypes(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("employee/employee-info/add-document", fiber.Map{
		"Title":            "Add Employee Document",
		"Employee":      employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"DocumentTypes": documentTypes,
	})
}

func (ctrl *EmployeeController) addDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.CreateDocumentRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.documentService.CreateDocument(c.Context(), c, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' added document for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/"+id+"/details")
}

func (ctrl *EmployeeController) editDocumentForm(c *fiber.Ctx) error {
	empId := c.Params("empId")
	docId := c.Params("docId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), empId)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	document, err := ctrl.documentService.GetDocumentByUniqueId(c.Context(), docId)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	documentTypes, err := ctrl.documentTypeService.GetAllDocumentTypes(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("employee/employee-info/edit-document", fiber.Map{
		"Title":            "Edit Document Info",
		"Employee":      employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"DocumentTypes": documentTypes,
		"Document": document,
	})
}

func (ctrl *EmployeeController) editDocument(c *fiber.Ctx) error {
	empId := c.Params("empId")
	docId := c.Params("docId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), empId)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	document, err := ctrl.documentService.GetDocumentByUniqueId(c.Context(), docId)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.UpdateDocumentRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.documentService.UpdateDocument(c.Context(), document.DocumentId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' edit document for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/"+empId+"/details")
}

func (ctrl *EmployeeController) displayDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	document, err := ctrl.documentService.GetDocumentByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	documentPath := config.UploadDocumentPath()+"/employees"
	
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' displayed document '%s'", loggedUser.UserName, document.DocumentName))
	return c.SendFile(documentPath+"/"+document.FileName)
}

