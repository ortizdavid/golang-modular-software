package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type RootController struct {
	authService *authentication.AuthService
	configService *configurations.AppConfigurationService
	statisticsService *services.StatisticsService
}

func NewRootController(db *database.Database) *RootController {
	return &RootController{
		authService:       authentication.NewAuthService(db),
		configService:     configurations.NewAppConfigurationService(db),
		statisticsService: services.NewStatisticsService(db),
	}
}

func (ctrl *RootController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/company", authMiddleware.CheckLoggedUser)
	group.Get("/", ctrl.index)
}

func (ctrl *RootController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	statistics, _ := ctrl.statisticsService.GetStatistics(c.Context())
	return c.Render("company/_root/index", fiber.Map{
		"Title":      "Company Management",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"Statistics": statistics,
	})
}
