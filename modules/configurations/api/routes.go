package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterRoutes(router *fiber.App, db *database.Database) {
	NewBasicConfigurationApi(db).Routes(router, db)
	NewCompanyConfigurationApi(db).Routes(router, db)
	NewEmailConfigurationApi(db).Routes(router, db)
	NewModuleFlagApi(db).Routes(router, db)
	NewCoreEntityFlagApi(db).Routes(router, db)
}