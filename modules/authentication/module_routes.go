package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/api"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/controllers"
	"gorm.io/gorm"
)

func RegisterModuleRoutes(router *fiber.App, db *gorm.DB) {
	controllers.RegisterControllerRoutes(router, db)
	api.RegisterApiRoutes(router, db)
}