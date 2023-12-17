package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/configurations"
	"github.com/ortizdavid/golang-modular-software/modules/authentication"
	"github.com/ortizdavid/golang-modular-software/modules/humanresources"
	"github.com/ortizdavid/golang-modular-software/modules/customers"
)


// RegisterControllers initializes and registers controllers (routes) from different modules
func RegisterControllers(router *fiber.App) {
	// Register routes for configurations module
	configurations.RegisterRoutes(router)

	// Register routes for authentication module
	authentication.RegisterRoutes(router)

	// Register routes for human resources module
	humanresources.RegisterRoutes(router)

	// Register routes for customers module
	customers.RegisterRoutes(router)
}