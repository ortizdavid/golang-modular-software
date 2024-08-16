package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
	configuration "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type CompanyController struct {
	service *services.CompanyService
	authService *authentication.AuthService
	appConfig *configuration.AppConfiguration
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewCompanyController(db *database.Database) *CompanyController {
	return &CompanyController{
		service:     services.NewCompanyService(db),
		authService: authentication.NewAuthService(db),
		appConfig:   configuration.LoadAppConfigurations(db),
		infoLogger:  helpers.NewInfoLogger("company-info.log"),
		errorLogger: helpers.NewErrorLogger("company-error.log"),
	}
}

func (ctrl *CompanyController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/company/company-info", authMiddleware.CheckLoggedUser)
	group.Get("", ctrl.index)
	group.Get("/create", ctrl.createForm)
	group.Get("/:id/details", ctrl.details)
	group.Get("/:id/edit", ctrl.editForm)
}

func (ctrl *CompanyController) index(c *fiber.Ctx) error  {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("company/company-info/index", fiber.Map{
		"Title": "Company Info",
		"AppConfig": ctrl.appConfig,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *CompanyController) details(c *fiber.Ctx) error  {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("company/company-info/details", fiber.Map{
		"Title": "Company Details",
		"AppConfig": ctrl.appConfig,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *CompanyController) createForm(c *fiber.Ctx) error  {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("company/company-info/create", fiber.Map{
		"Title": "Create Company",
		"AppConfig": ctrl.appConfig,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *CompanyController) editForm(c *fiber.Ctx) error  {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("company/company-info/edit", fiber.Map{
		"Title": "Edit Company Info",
		"AppConfig": ctrl.appConfig,
		"LoggedUser": loggedUser,
	})
}