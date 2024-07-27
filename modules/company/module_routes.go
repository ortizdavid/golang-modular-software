package company

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/api"
	"github.com/ortizdavid/golang-modular-software/modules/company/controllers"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	controllers.RegisterControllerRoutes(router, db)
	api.RegisterApiRoutes(router, db)
}