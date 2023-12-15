package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/users/controllers"
)


func RegisterRoutes(router *fiber.App) {
	controllers.UserController{}.Routes(router)
}