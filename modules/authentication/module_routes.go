package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/api"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/controllers"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	moduleName := "Authentication"
	// Check Module flags
	moduleMiddleware := middlewares.NewModuleFlagMiddleware(db)
	router.Group("/users", moduleMiddleware.CheckModule(moduleName))
	router.Group("/roles", moduleMiddleware.CheckModule(moduleName))
	router.Group("/permissions", moduleMiddleware.CheckModule(moduleName))
	router.Group("/login-activities", moduleMiddleware.CheckModule(moduleName))
	router.Group("/api/users", moduleMiddleware.CheckModule(moduleName))
	router.Group("/api/roles", moduleMiddleware.CheckModule(moduleName))
	router.Group("/api/permissions", moduleMiddleware.CheckModule(moduleName))
	//-----------Register all module routes
	controllers.RegisterControllerRoutes(router, db)
	api.RegisterApiRoutes(router, db)
}