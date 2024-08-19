package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type RootController struct {
	authService       *services.AuthService
	configService *configurations.AppConfigurationService
	statisticsService *services.StatisticsService
}

func NewRootController(db *database.Database) *RootController {
	return &RootController{
		authService:       services.NewAuthService(db),
		configService: configurations.NewAppConfigurationService(db),
		statisticsService: services.NewStatisticsService(db),
	}
}

func (ctrl *RootController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/users-management", authMiddleware.CheckLoggedUser)
	group.Get("/", ctrl.index)

}

func (ctrl *RootController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	statistics, _ := ctrl.statisticsService.GetStatistics(c.Context())
	return c.Render("authentication/_root/index", fiber.Map{
		"Title":      "Users Management",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"Statistics": statistics,
	})
}
