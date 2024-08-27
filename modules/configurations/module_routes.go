package configurations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/api"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/controllers"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	controllers.RegisterRoutes(router, db)
	api.RegisterRoutes(router, db)
}