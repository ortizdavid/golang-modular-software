package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func RegisterRoutes(router *fiber.App, db *gorm.DB) {
	NewBasicConfigurationController(db).Routes(router)
	NewCompanyConfigurationController(db).Routes(router)
	NewEmailConfigurationController(db).Routes(router)
}