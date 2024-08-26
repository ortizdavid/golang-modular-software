package reports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/reports/api"
	"github.com/ortizdavid/golang-modular-software/modules/reports/controllers"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	moduleName := "Reports"
	// Check Module flags
	moduleMiddleware := middlewares.NewModuleFlagMiddleware(db)
	router.Group("/reports", moduleMiddleware.CheckModule(moduleName))
	router.Group("/api/reports", moduleMiddleware.CheckModule(moduleName))
	// Register all routes
	controllers.RegisterControllerRoutes(router, db)
	api.RegisterApiRoutes(router, db)
}