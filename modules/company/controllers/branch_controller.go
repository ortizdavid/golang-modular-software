package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	//"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
	configuration "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type BranchController struct {
	service *services.BranchService
	authService *authentication.AuthService
	appConfig *configuration.AppConfiguration
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewBranchController(db *database.Database) *BranchController {
	return &BranchController{
		service:     services.NewBranchService(db),
		authService: authentication.NewAuthService(db),
		appConfig:   configuration.LoadAppConfigurations(db),
		infoLogger:  helpers.NewInfoLogger("company-info.log"),
		errorLogger: helpers.NewErrorLogger("company-error.log"),
	}
}

func (ctrl *BranchController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/company/branches", authMiddleware.CheckLoggedUser)
	group.Get("", nil)
	group.Get("/:id/details", nil)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", nil)
	group.Get("/:id/edit", nil)
	group.Post("/:id/edit", nil)
	group.Get("/searcch", nil)
	group.Post("/search", nil)
}

func (ctrl *BranchController) createForm(c *fiber.Ctx) error  {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("company/branch/create", fiber.Map{
		"Title": "Create Branch",
		"AppConfig": ctrl.appConfig,
		"LoggedUser": loggedUser,
	})
}