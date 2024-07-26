package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/authentication"
	"github.com/ortizdavid/golang-modular-software/modules/configurations"
	"github.com/ortizdavid/golang-modular-software/modules/company"
	"gorm.io/gorm"
)

// RegisterControllers initializes and registers controllers (routes) from different modules
func RegisterControllers(router *fiber.App, db *gorm.DB) {
	// Register routes for configurations module
	configurations.RegisterModuleRoutes(router, db)

	// Register routes for authentication module
	authentication.RegisterModuleRoutes(router, db)

	company.RegisterModuleRoutes(router, db)
}