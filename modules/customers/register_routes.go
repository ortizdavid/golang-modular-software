package customers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/customers/controllers"
	"gorm.io/gorm"
)


func RegisterRoutes(router *fiber.App, db *gorm.DB) {
	controllers.CustomerController{}.Routes(router, db)
}