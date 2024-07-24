package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"gorm.io/gorm"
)

var configurationLogger = config.NewLogger("configurations-api.log")

func RegisterRoutes(router *fiber.App, db *gorm.DB) {
	NewBasicConfigurationApi(db).Routes(router)
	NewCompanyConfigurationApi(db).Routes(router)
	NewEmailConfigurationApi(db).Routes(router)
}