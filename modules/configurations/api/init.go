package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterRoutes(router *fiber.App, db *database.Database) {
	NewBasicConfigurationApi(db).Routes(router)
	NewCompanyConfigurationApi(db).Routes(router)
	NewEmailConfigurationApi(db).Routes(router)
}