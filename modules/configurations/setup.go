package configurations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/api"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/controllers"
	"gorm.io/gorm"
)

func RegisterRoutes(router *fiber.App, db *gorm.DB) {
	controllers.RegisterRoutes(router, db)
	api.RegisterRoutes(router, db)
}