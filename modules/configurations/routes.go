package configurations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/controllers"
)


func RegisterRoutes(router *fiber.App) {
	controllers.BasicConfigController{}.Routes(router)
	controllers.EmailConfigController{}.Routes(router)
}