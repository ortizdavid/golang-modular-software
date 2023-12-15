package customers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/customers/controllers"
)


func RegisterRoutes(router *fiber.App) {
	controllers.CustomerController{}.Routes(router)
}