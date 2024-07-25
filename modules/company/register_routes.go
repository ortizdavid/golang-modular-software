package humanresources

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/modules/company/controllers"
	"gorm.io/gorm"
)


func RegisterRoutes(router *fiber.App, db *gorm.DB) {
	controllers.CompanyController{}.Routes(router)
}