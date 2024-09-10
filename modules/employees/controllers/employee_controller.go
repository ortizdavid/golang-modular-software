package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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
	documentTypeService         *services.DocumentTypeService
	identTypeService		*references.IdentificationTypeService
	countryService			*references.CountryService
	maritalStatusService	*references.MaritalStatusService
	employmentStatusService	*references.EmploymentStatusService
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
		documentTypeService:     services.NewDocumentTypeService(db),
		identTypeService:        references.NewIdentificationTypeService(db),
		countryService:          references.NewCountryService(db),
		maritalStatusService:    references.NewMaritalStatusService(db),
		employmentStatusService: references.NewEmploymentStatusService(db),
		departmentService:       company.NewDepartmentService(db),
		authService:             authentication.NewAuthService(db),
		moduleFlagStatusService: configurations.NewModuleFlagStatusService(db),
		configService:           configurations.NewAppConfigurationService(db),
		infoLogger:              helpers.NewInfoLogger("employee-info.log"),
		errorLogger:             helpers.NewErrorLogger("employee-error.log"),
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
	group.Get("/:id/add-document", ctrl.addDocumentForm)
	group.Post("/:id/add-document", ctrl.addDocument)
	group.Get("/:empId/display-document/:docId", ctrl.addDocumentForm)
	group.Get("/:id/delete", ctrl.removeForm)
	group.Post("/:id/delete", ctrl.remove)
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
	return c.Render("employee/employee-info/details", fiber.Map{
		"Title":            "Details",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"Employee":      employee,
		"EmployeeDocuments": employeeDocuments,
		"CountDocuments": len(employeeDocuments),
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


func (ctrl *EmployeeController) displayDocument(c *fiber.Ctx) error {
	/**empId := c.Params("empId")
	docId := c.Params("docId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.RemoveEmployee(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}**
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' removed employment '%s'", loggedUser.UserName, employee.FirstName))*/
	return c.Redirect("/employees/employee-info")
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
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' removed employment '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info")
}


