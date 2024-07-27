package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication"
	"github.com/ortizdavid/golang-modular-software/modules/company"
	"github.com/ortizdavid/golang-modular-software/modules/configurations"
)

// RegisterControllers initializes and registers controllers (routes) from different modules
func RegisterControllers(router *fiber.App, db *database.Database) {
	// Register routes for configurations module
	configurations.RegisterModuleRoutes(router, db)

	// Register routes for authentication module
	authentication.RegisterModuleRoutes(router, db)

	company.RegisterModuleRoutes(router, db)
}