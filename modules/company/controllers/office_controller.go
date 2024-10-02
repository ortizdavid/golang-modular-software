package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type OfficeController struct {
	service                 *services.OfficeService
	companyService          *services.CompanyService
	authService             *authentication.AuthService
	moduleFlagStatusService *configurations.ModuleFlagStatusService
	configService           *configurations.AppConfigurationService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
	shared.BaseController
}

func NewOfficeController(db *database.Database) *OfficeController {
	return &OfficeController{
		service:                 services.NewOfficeService(db),
		companyService:          services.NewCompanyService(db),
		authService:             authentication.NewAuthService(db),
		moduleFlagStatusService: configurations.NewModuleFlagStatusService(db),
		configService:           configurations.NewAppConfigurationService(db),
		infoLogger:              helpers.NewInfoLogger(infoLogFile),
		errorLogger:             helpers.NewErrorLogger(errorLogFile),
	}
}

func (ctrl *OfficeController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/company/offices")
	group.Get("", ctrl.index)
	group.Get("/:id/details", ctrl.details)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
}

func (ctrl *OfficeController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("company/office/index", fiber.Map{
		"Title":            "Offices",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"Pagination":       pagination,
		"CurrentPage":      pagination.MetaData.CurrentPage + 1,
		"TotalPages":       pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *OfficeController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	office, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("company/office/details", fiber.Map{
		"Title":            "Details",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"Office":           office,
	})
}

func (ctrl *OfficeController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	companies, err := ctrl.companyService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("company/office/create", fiber.Map{
		"Title":            "Create Office",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"Companies":        companies,
		"OfficeCode":       encryption.GenerateCode("OF-"),
	})
}

func (ctrl *OfficeController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateOfficeRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err := ctrl.service.Create(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Created office '"+request.OfficeName+"' successfully")
	return c.Redirect("/company/offices")
}

func (ctrl *OfficeController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	office, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("company/office/edit", fiber.Map{
		"Title":            "Edit Office",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"Office":           office,
	})
}

func (ctrl *OfficeController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	office, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.UpdateOfficeRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.Update(c.Context(), office.OfficeId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Updated Office '"+request.OfficeName+"' successfully")
	return c.Redirect("/company/offices/" + id + "/details")
}

func (ctrl *OfficeController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("company/office/search", fiber.Map{
		"Title":            "Search Offices",
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *OfficeController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchOfficeRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.Search(c.Context(), c, request, params)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("company/office/search-results", fiber.Map{
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
