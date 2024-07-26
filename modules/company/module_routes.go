package company

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/company/api"
	"github.com/ortizdavid/golang-modular-software/modules/company/controllers"
	"gorm.io/gorm"
)

func RegisterModuleRoutes(router *fiber.App, db *gorm.DB) {
	controllers.RegisterControllerRoutes(router, db)
	api.RegisterApiRoutes(router, db)
}