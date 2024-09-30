package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

func (ctrl *EmployeeController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllEmployeesPaginated(c.Context(), c, params)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
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
		return ctrl.HandleErrorsWeb(c, err)
	}
	employeeDocuments, err := ctrl.documentService.GetAllEmployeeDocuments(c.Context(), employee.EmployeeId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employeeProfessionalInfo, _ := ctrl.professionalInfoService.GetProfessionalInfoByEmployeeId(c.Context(), employee.EmployeeId)

	employeePhones, err := ctrl.phoneService.GetAllEmployeePhones(c.Context(), employee.EmployeeId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employeeEmails, err := ctrl.emailService.GetAllEmployeeEmails(c.Context(), employee.EmployeeId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employeeAddresses, err := ctrl.addressService.GetAllEmployeeAddresses(c.Context(), employee.EmployeeId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employeeAccount, err := ctrl.accountService.GetEmployeeAccountById(c.Context(), employee.EmployeeId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/details", fiber.Map{
		"Title":             "Details",
		"AppConfig":         ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus":  moduleFlagStatus,
		"LoggedUser":        loggedUser,
		"Employee":          employee,
		"EmployeeDocuments": employeeDocuments,
		"EmployeeAddresses": employeeAddresses,
		"CountAddresses": len(employeeAddresses),
		"EmployeeProfessionalInfo": employeeProfessionalInfo,
		"CountDocuments":    len(employeeDocuments),
		"EmployeePhones":    employeePhones,
		"CountPhones":       len(employeePhones),
		"EmployeeEmails":    employeeEmails,
		"CountEmails":       len(employeeEmails),
		"EmployeeAccount": employeeAccount,
	})
}

func (ctrl *EmployeeController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())

	jobTitles, err := ctrl.jobTitleService.GetAllJobTitles(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	identificationTypes, err := ctrl.identTypeService.GetAllIdentificationTypes(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	countries, err := ctrl.countryService.GetAllCountries(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	maritalStatuses, err := ctrl.maritalStatusService.GetAllStatuses(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employmentStatuses, err := ctrl.employmentStatusService.GetAllStatuses(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	departments, err := ctrl.departmentService.GetAllDepartments(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/create", fiber.Map{
		"Title":               "Create Employee",
		"AppConfig":           ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus":    moduleFlagStatus,
		"LoggedUser":          loggedUser,
		"JobTitles":           jobTitles,
		"IdentificationTypes": identificationTypes,
		"Countries":           countries,
		"MaritalStatuses":     maritalStatuses,
		"EmploymentStatuses":  employmentStatuses,
		"Departments":         departments,
		"DateInfo": helpers.GetDateInfo(),
	})
}

func (ctrl *EmployeeController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateEmployeeRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err := ctrl.service.CreateEmployee(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Created employment '"+request.IdentificationNumber+"' successfully")
	return c.Redirect("/employees/employee-info")
}

func (ctrl *EmployeeController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}

	jobTitles, err := ctrl.jobTitleService.GetAllJobTitles(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	identificationTypes, err := ctrl.identTypeService.GetAllIdentificationTypes(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	countries, err := ctrl.countryService.GetAllCountries(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	maritalStatuses, err := ctrl.maritalStatusService.GetAllStatuses(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employmentStatuses, err := ctrl.employmentStatusService.GetAllStatuses(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	departments, err := ctrl.departmentService.GetAllDepartments(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/edit", fiber.Map{
		"Title":               "Edit Employee",
		"AppConfig":           ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus":    moduleFlagStatus,
		"LoggedUser":          loggedUser,
		"Employee":            employee,
		"JobTitles":           jobTitles,
		"IdentificationTypes": identificationTypes,
		"Countries":           countries,
		"MaritalStatuses":     maritalStatuses,
		"EmploymentStatuses":  employmentStatuses,
		"Departments":         departments,
	})
}

func (ctrl *EmployeeController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.UpdateEmployeeRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.UpdateEmployee(c.Context(), employee.EmployeeId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Updated employee '"+request.IdentificationNumber+"' successfully")
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
		return ctrl.HandleErrorsWeb(c, err)
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
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/delete", fiber.Map{
		"Title":            "Remove Employee",
		"Employee":         employee,
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
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.RemoveEmployee(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' removed employee '%s'", loggedUser.UserName, employee.IdentificationNumber))
	return c.Redirect("/employees/employee-info")
}
