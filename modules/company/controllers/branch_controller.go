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

type BranchController struct {
	service        *services.BranchService
	companyService *services.CompanyService
	authService    *authentication.AuthService
	configService *configurations.AppConfigurationService
	infoLogger     *helpers.Logger
	errorLogger    *helpers.Logger
}

func NewBranchController(db *database.Database) *BranchController {
	return &BranchController{
		service:        services.NewBranchService(db),
		companyService: services.NewCompanyService(db),
		authService:    authentication.NewAuthService(db),
		configService: configurations.NewAppConfigurationService(db),
		infoLogger:     helpers.NewInfoLogger("company-info.log"),
		errorLogger:    helpers.NewErrorLogger("company-error.log"),
	}
}

func (ctrl *BranchController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/company/branches", authMiddleware.CheckLoggedUser)
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
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllCompaniesPaginated(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/branch/index", fiber.Map{
		"Title":       "Branches",
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":  loggedUser,
		"Pagination":  pagination,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages":  pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *BranchController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	branch, err := ctrl.service.GetBranchByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/branch/details", fiber.Map{
		"Title":      "Details",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"Branch":     branch,
	})
}

func (ctrl *BranchController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	companies, err := ctrl.companyService.GetAllCompanies(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/branch/create", fiber.Map{
		"Title":      "Create Branch",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"Companies":  companies,
		"BranchCode": encryption.GenerateCode("BRA-"),
	})
}

func (ctrl *BranchController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateBranchRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.CreateBranch(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Created branch '"+request.BranchName+"' successfully")
	return c.Redirect("/company/branches")
}

func (ctrl *BranchController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	branch, err := ctrl.service.GetBranchByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/branch/edit", fiber.Map{
		"Title":      "Edit Branch",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"Branch":     branch,
	})
}

func (ctrl *BranchController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	branch, err := ctrl.service.GetBranchByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.UpdateBranchRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.UpdateBranch(c.Context(), branch.BranchId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Updated Branch '"+request.BranchName+"' successfully")
	return c.Redirect("/company/branches/" + id + "/details")
}

func (ctrl *BranchController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("company/branch/search", fiber.Map{
		"Title":      "Search Branches",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
	})
}

func (ctrl *BranchController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchBranchRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.SearchBranches(c.Context(), c, request, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("company/branch/search-results", fiber.Map{
		"Title":       "Search Results",
		"Pagination":  pagination,
		"Param":       request.SearchParam,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages":  pagination.MetaData.TotalPages + 1,
		"LoggedUser":  loggedUser,
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
	})
}
