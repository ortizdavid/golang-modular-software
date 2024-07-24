package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"gorm.io/gorm"
)

var configurationLogger = config.NewLogger("configurations.log")

func RegisterRoutes(router *fiber.App, db *gorm.DB) {
	NewBasicConfigurationController(db).Routes(router)
	NewCompanyConfigurationController(db).Routes(router)
	NewEmailConfigurationController(db).Routes(router)
}