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

type BranchController struct {
	service                 *services.BranchService
	companyService          *services.CompanyService
	moduleFlagStatusService *configurations.ModuleFlagStatusService
	authService             *authentication.AuthService
	configService           *configurations.AppConfigurationService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
	shared.BaseController
}

func NewBranchController(db *database.Database) *BranchController {
	return &BranchController{
		service:                 services.NewBranchService(db),
		companyService:          services.NewCompanyService(db),
		moduleFlagStatusService: configurations.NewModuleFlagStatusService(db),
		authService:             authentication.NewAuthService(db),
		configService:           configurations.NewAppConfigurationService(db),
		infoLogger:              helpers.NewInfoLogger(infoLogFile),
		errorLogger:             helpers.NewErrorLogger(errorLogFile),
	}
}

func (ctrl *BranchController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/company/branches")
	group.Get("", ctrl.index)
	group.Get("/:id/details", ctrl.details)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
}

func (ctrl *BranchController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("company/branch/index", fiber.Map{
		"Title":            "Branches",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"Pagination":       pagination,
		"CurrentPage":      pagination.MetaData.CurrentPage + 1,
		"TotalPages":       pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *BranchController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	branch, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("company/branch/details", fiber.Map{
		"Title":            "Details",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"Branch":           branch,
	})
}

func (ctrl *BranchController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	companies, err := ctrl.companyService.GetAll(c.Context())
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("company/branch/create", fiber.Map{
		"Title":            "Create Branch",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"Companies":        companies,
		"BranchCode":       encryption.GenerateCode("BRA-"),
	})
}

func (ctrl *BranchController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateBranchRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err := ctrl.service.Create(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Created branch '"+request.BranchName+"' successfully")
	return c.Redirect("/company/branches")
}

func (ctrl *BranchController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	branch, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("company/branch/edit", fiber.Map{
		"Title":            "Edit Branch",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"Branch":           branch,
	})
}

func (ctrl *BranchController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.UpdateBranchRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err := ctrl.service.Update(c.Context(), id, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Updated Branch '"+request.BranchName+"' successfully")
	return c.Redirect("/company/branches/" + id + "/details")
}

func (ctrl *BranchController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("company/branch/search", fiber.Map{
		"Title":      "Search Branches",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"FlagStatus": moduleFlagStatus,
	})
}

func (ctrl *BranchController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchBranchRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.Search(c.Context(), c, request, params)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("company/branch/search-results", fiber.Map{
		"Title":            "Search Results",
		"LoggedUser":       loggedUser,
		"ModuleFlagStatus": moduleFlagStatus,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"Pagination":       pagination,
		"Param":            request.SearchParam,
		"CurrentPage":      pagination.MetaData.CurrentPage + 1,
		"TotalPages":       pagination.MetaData.TotalPages + 1,
	})
}
