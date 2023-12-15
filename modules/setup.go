package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/authentication"
	"github.com/ortizdavid/golang-modular-software/modules/configurations"
	"github.com/ortizdavid/golang-modular-software/modules/users"
)

func RegisterControllers(router *fiber.App) {
	configurations.RegisterRoutes(router)
	authentication.RegisterRoutes(router)
	users.RegisterRoutes(router)
}