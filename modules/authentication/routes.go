package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/controllers"
	"gorm.io/gorm"
)


func RegisterRoutes(router *fiber.App, db *gorm.DB) {
	controllers.UserController{}.Routes(router)
	controllers.AuthController{}.Routes(router)
}