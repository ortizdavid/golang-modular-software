package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type DepartmentController struct {
	service        *services.DepartmentService
	companyService *services.CompanyService
	authService    *authentication.AuthService
	flagStatusService *configurations.ModuleFlagStatusService
	configService *configurations.AppConfigurationService
	infoLogger     *helpers.Logger
	errorLogger    *helpers.Logger
}

func NewDepartmentController(db *database.Database) *DepartmentController {
	return &DepartmentController{
		service:           services.NewDepartmentService(db),
		companyService:    services.NewCompanyService(db),
		authService:       authentication.NewAuthService(db),
		flagStatusService: configurations.NewModuleFlagStatusService(db),
		configService:     configurations.NewAppConfigurationService(db),
		infoLogger:        helpers.NewInfoLogger("company-info.log"),
		errorLogger:       helpers.NewErrorLogger("company-error.log"),
	}
}

func (ctrl *DepartmentController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/company/departments", authMiddleware.CheckLoggedUser)
	group.Get("", ctrl.index)
	group.Get("/:id/details", ctrl.details)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
}

func (ctrl *DepartmentController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllDepartmentsPaginated(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/department/index", fiber.Map{
		"Title":       "Departments",
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"LoggedUser":  loggedUser,
		"Pagination":  pagination,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages":  pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *DepartmentController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	department, err := ctrl.service.GetDepartmentByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/department/details", fiber.Map{
		"Title":      "Details",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"LoggedUser": loggedUser,
		"Department":     department,
	})
}

func (ctrl *DepartmentController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	companies, err := ctrl.companyService.GetAllCompanies(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/department/create", fiber.Map{
		"Title":      "Create Department",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"LoggedUser": loggedUser,
		"Companies":  companies,
		"DepartmentCode": encryption.GenerateCode("DPT-"),
	})
}

func (ctrl *DepartmentController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateDepartmentRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.CreateDepartment(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Created department '"+request.DepartmentName+"' successfully")
	return c.Redirect("/company/departments")
}

func (ctrl *DepartmentController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	companies, err := ctrl.companyService.GetAllCompanies(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	department, err := ctrl.service.GetDepartmentByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/department/edit", fiber.Map{
		"Title":      "Edit Department",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"LoggedUser": loggedUser,
		"Department":     department,
		"Companies":  companies,
	})
}

func (ctrl *DepartmentController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	department, err := ctrl.service.GetDepartmentByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.UpdateDepartmentRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.UpdateDepartment(c.Context(), department.DepartmentId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Updated Department '"+request.DepartmentName+"' successfully")
	return c.Redirect("/company/departments/" + id + "/details")
}

func (ctrl *DepartmentController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("company/department/search", fiber.Map{
		"Title":      "Search Departments",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
	})
}

func (ctrl *DepartmentController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchDepartmentRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.SearchDepartments(c.Context(), c, request, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("company/department/search-results", fiber.Map{
		"Title":       "Search Results",
		"LoggedUser":  loggedUser,
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"Pagination":  pagination,
		"Param":       request.SearchParam,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages":  pagination.MetaData.TotalPages + 1,
	})
}
