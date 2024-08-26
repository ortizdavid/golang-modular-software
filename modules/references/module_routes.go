package references

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/api"
	"github.com/ortizdavid/golang-modular-software/modules/references/controllers"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	moduleName := "References"
	// Check Module flags
	moduleMiddleware := middlewares.NewModuleFlagMiddleware(db)
	router.Group("/references", moduleMiddleware.CheckModule(moduleName))
	router.Group("/api/references", moduleMiddleware.CheckModule(moduleName))
	// Register all routes
	controllers.RegisterControllerRoutes(router, db)
	api.RegisterApiRoutes(router, db)
}