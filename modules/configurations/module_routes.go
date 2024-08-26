package configurations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/api"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/controllers"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	moduleName := "Configurations"
	// Check Module flags
	moduleMiddleware := middlewares.NewModuleFlagMiddleware(db)
	router.Group("/configurations", moduleMiddleware.CheckModule(moduleName))
	router.Group("/api/configurations", moduleMiddleware.CheckModule(moduleName))
	// Register all routes
	controllers.RegisterRoutes(router, db)
	api.RegisterRoutes(router, db)
}