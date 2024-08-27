package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/api"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/controllers"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	controllers.RegisterControllerRoutes(router, db)
	api.RegisterApiRoutes(router, db)
}
