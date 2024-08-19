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

type OfficeController struct {
	service        *services.OfficeService
	companyService *services.CompanyService
	authService    *authentication.AuthService
	configService *configurations.AppConfigurationService
	infoLogger     *helpers.Logger
	errorLogger    *helpers.Logger
}

func NewOfficeController(db *database.Database) *OfficeController {
	return &OfficeController{
		service:        services.NewOfficeService(db),
		companyService: services.NewCompanyService(db),
		authService:    authentication.NewAuthService(db),
		configService: configurations.NewAppConfigurationService(db),
		infoLogger:     helpers.NewInfoLogger("company-info.log"),
		errorLogger:    helpers.NewErrorLogger("company-error.log"),
	}
}

func (ctrl *OfficeController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/company/offices", authMiddleware.CheckLoggedUser)
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
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllCompaniesPaginated(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/office/index", fiber.Map{
		"Title":       "Offices",
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":  loggedUser,
		"Pagination":  pagination,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages":  pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *OfficeController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	office, err := ctrl.service.GetOfficeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/office/details", fiber.Map{
		"Title":      "Details",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"Office":     office,
	})
}

func (ctrl *OfficeController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	companies, err := ctrl.companyService.GetAllCompanies(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/office/create", fiber.Map{
		"Title":      "Create Office",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"Companies":  companies,
		"OfficeCode": encryption.GenerateCode("OF-"),
	})
}

func (ctrl *OfficeController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateOfficeRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.CreateOffice(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Created office '"+request.OfficeName+"' successfully")
	return c.Redirect("/company/offices")
}

func (ctrl *OfficeController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	office, err := ctrl.service.GetOfficeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/office/edit", fiber.Map{
		"Title":      "Edit Office",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"Office":     office,
	})
}

func (ctrl *OfficeController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	office, err := ctrl.service.GetOfficeByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.UpdateOfficeRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.UpdateOffice(c.Context(), office.OfficeId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Updated Office '"+request.OfficeName+"' successfully")
	return c.Redirect("/company/offices/" + id + "/details")
}

func (ctrl *OfficeController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("company/office/search", fiber.Map{
		"Title":      "Search Offices",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
	})
}

func (ctrl *OfficeController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchOfficeRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.SearchOffices(c.Context(), c, request, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("company/office/search-results", fiber.Map{
		"Title":       "Search Results",
		"Pagination":  pagination,
		"Param":       request.SearchParam,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages":  pagination.MetaData.TotalPages + 1,
		"LoggedUser":  loggedUser,
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
	})
}
