package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/controllers"
)


func RegisterRoutes(router *fiber.App) {
	controllers.AuthController{}.Routes(router)
}