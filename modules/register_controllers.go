package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication"
	"github.com/ortizdavid/golang-modular-software/modules/back_office"
	"github.com/ortizdavid/golang-modular-software/modules/company"
	"github.com/ortizdavid/golang-modular-software/modules/employees"
	"github.com/ortizdavid/golang-modular-software/modules/configurations"
	"github.com/ortizdavid/golang-modular-software/modules/references"
	"github.com/ortizdavid/golang-modular-software/modules/reports"
)

// RegisterControllers initializes and registers controllers (routes) from different modules
func RegisterControllers(router *fiber.App, db *database.Database) {
	
	// Register routes for configurations module
	configurations.RegisterModuleRoutes(router, db)

	// Register routes for reference module
	references.RegisterModuleRoutes(router, db)

	// Register routes for authentication module
	authentication.RegisterModuleRoutes(router, db)

	// Register routes for company module
	company.RegisterModuleRoutes(router, db)

	// Register routes for employees module
	employees.RegisterModuleRoutes(router, db)

	// Register routes for BackOffice
	back_office.RegisterModuleRoutes(router, db)

	// Register routes for BackOffice
	reports.RegisterModuleRoutes(router, db)

}