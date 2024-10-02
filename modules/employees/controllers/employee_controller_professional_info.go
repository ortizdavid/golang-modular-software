package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

func (ctrl *EmployeeController) addProfessionalInfoForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	jobTitles, err := ctrl.jobTitleService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employmentStatuses, err := ctrl.employmentStatusService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	departments, err := ctrl.departmentService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/add-professional-info", fiber.Map{
		"Title":            "Add Professional Info",
		"Employee":         employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"EmploymentStatuses":  employmentStatuses,
		"Departments":         departments,
		"JobTitles": jobTitles,
	})
}

func (ctrl *EmployeeController) addProfessionalInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.CreateProfessionalInfoRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.professionalInfoService.CreateProfessionalInfo(c.Context(), c, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' added professional info for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/" + id + "/details")
}

func (ctrl *EmployeeController) editProfessionalInfoForm(c *fiber.Ctx) error {
	empId := c.Params("empId")
	proId := c.Params("proId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	employee, err := ctrl.service.GetByUniqueId(c.Context(), empId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	professionalInfo, err := ctrl.professionalInfoService.GetByUniqueId(c.Context(), proId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	jobTitles, err := ctrl.jobTitleService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	employmentStatuses, err := ctrl.employmentStatusService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	departments, err := ctrl.departmentService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/employee-info/edit-professional-info", fiber.Map{
		"Title":            "Edit Document Info",
		"Employee":         employee,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"ProfessionalInfo": professionalInfo,
		"EmploymentStatuses":  employmentStatuses,
		"Departments":         departments,
		"JobTitles": jobTitles,
	})
}

func (ctrl *EmployeeController) editProfessionalInfo(c *fiber.Ctx) error {
	empId := c.Params("empId")
	proId := c.Params("proId")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	employee, err := ctrl.service.GetByUniqueId(c.Context(), empId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	professionalInfo, err := ctrl.professionalInfoService.GetByUniqueId(c.Context(), proId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.UpdateProfessionalInfoRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.professionalInfoService.UpdateProfessionalInfo(c.Context(), professionalInfo.ProfessionalId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' edit professionalInfo for employee '%s'", loggedUser.UserName, employee.FirstName))
	return c.Redirect("/employees/employee-info/" + empId + "/details")
}
