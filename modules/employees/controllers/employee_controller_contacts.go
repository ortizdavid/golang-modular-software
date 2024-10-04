package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

func (ctrl *EmployeeController) addPhoneForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	contactTypes, err := ctrl.contactTypeService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	countries, err := ctrl.countryService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/add-phone", fiber.Map{
		"Title":            "Add Employee Phone",
		"Employee":         employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"ContactTypes":     contactTypes,
		"Countries": countries,
	})
}

func (ctrl *EmployeeController) addPhone(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.CreateEmployeePhoneRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.phoneService.Create(c.Context(), c, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' added phone for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/" + id + "/details")
}

func (ctrl *EmployeeController) editPhoneForm(c *fiber.Ctx) error {
	empId := c.Params("empId")
	contId := c.Params("contId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetByUniqueId(c.Context(), empId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employeePhone, err := ctrl.phoneService.GetByUniqueId(c.Context(), contId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	contactTypes, err := ctrl.contactTypeService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/edit-phone", fiber.Map{
		"Title":            "Edit Employee Phone",
		"Employee":         employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"ContactTypes":     contactTypes,
		"EmployeePhone":    employeePhone,
	})
}

func (ctrl *EmployeeController) editPhone(c *fiber.Ctx) error {
	empId := c.Params("empId")
	contId := c.Params("contId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetByUniqueId(c.Context(), empId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employeePhone, err := ctrl.phoneService.GetByUniqueId(c.Context(), contId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.UpdateEmployeePhoneRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.phoneService.Update(c.Context(), employeePhone.PhoneId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' edit phone for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/" + empId + "/details")
}

func (ctrl *EmployeeController) addEmailForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	contactTypes, err := ctrl.contactTypeService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/add-email", fiber.Map{
		"Title":            "Add Employee Email",
		"Employee":         employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"ContactTypes":     contactTypes,
	})
}

func (ctrl *EmployeeController) addEmail(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.CreateEmployeeEmailRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.emailService.Create(c.Context(), c, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' added email for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/" + id + "/details")
}

func (ctrl *EmployeeController) editEmailForm(c *fiber.Ctx) error {
	empId := c.Params("empId")
	contId := c.Params("contId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetByUniqueId(c.Context(), empId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employeeEmail, err := ctrl.emailService.GetByUniqueId(c.Context(), contId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	contactTypes, err := ctrl.contactTypeService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/edit-email", fiber.Map{
		"Title":            "Edit Employee Email",
		"Employee":         employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"ContactTypes":     contactTypes,
		"EmployeeEmail":    employeeEmail,
	})
}

func (ctrl *EmployeeController) editEmail(c *fiber.Ctx) error {
	empId := c.Params("empId")
	contId := c.Params("contId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetByUniqueId(c.Context(), empId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employeeEmail, err := ctrl.emailService.GetByUniqueId(c.Context(), contId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.UpdateEmployeeEmailRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.emailService.Update(c.Context(), employeeEmail.EmailId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' edit email for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/" + empId + "/details")
}
