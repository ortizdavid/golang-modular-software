package company

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/api"
	"github.com/ortizdavid/golang-modular-software/modules/company/controllers"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	moduleName := "Company"
	// Check Module flags
	moduleMiddleware := middlewares.NewModuleFlagMiddleware(db)
	router.Group("/company", moduleMiddleware.CheckModule(moduleName))
	router.Group("/api/company", moduleMiddleware.CheckModule(moduleName))
	// Register all routes
	controllers.RegisterControllerRoutes(router, db)
	api.RegisterApiRoutes(router, db)
}