package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

func (ctrl *EmployeeController) addAddressForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	countries, err := ctrl.countryService.GetAllCountries(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/add-address", fiber.Map{
		"Title":            "Add Employee Address",
		"Employee":         employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"Countries": countries,
	})
}

func (ctrl *EmployeeController) addAddress(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.CreateAddressRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.addressService.CreateAddress(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' added address for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/" + id + "/details")
}

func (ctrl *EmployeeController) editAddressForm(c *fiber.Ctx) error {
	empId := c.Params("empId")
	addId := c.Params("addId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), empId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	countries, err := ctrl.countryService.GetAllCountries(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employeeAddress, err := ctrl.addressService.GetAddressByUniqueId(c.Context(), addId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/edit-address", fiber.Map{
		"Title":            "Edit Employee Address",
		"Employee":         employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"EmployeeAddress":	employeeAddress,
		"Countries": countries,
	})
}

func (ctrl *EmployeeController) editAddress(c *fiber.Ctx) error {
	empId := c.Params("empId")
	addId := c.Params("addId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetEmployeeByUniqueId(c.Context(), empId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	address, err := ctrl.addressService.GetAddressByUniqueId(c.Context(), addId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.UpdateAddressRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.addressService.UpdateAddress(c.Context(), address.AddressId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' updated address for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/" + empId + "/details")
}