package humanresources

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/humanresources/controllers"
)


func RegisterRoutes(router *fiber.App) {
	controllers.EmployeeController{}.Routes(router)
}