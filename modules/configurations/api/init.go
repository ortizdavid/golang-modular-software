package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router *fiber.App, db *gorm.DB) {
	NewBasicConfigurationApi(db).Routes(router)
	NewCompanyConfigurationApi(db).Routes(router)
	NewEmailConfigurationApi(db).Routes(router)
}