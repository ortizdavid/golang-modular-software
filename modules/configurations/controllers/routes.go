package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterRoutes(router *fiber.App, db *database.Database) {
	NewRootController(db).Routes(router, db)
	NewBasicConfigurationController(db).Routes(router, db)
	NewCompanyConfigurationController(db).Routes(router, db)
	NewEmailConfigurationController(db).Routes(router, db)
	NewModuleFlagController(db).Routes(router, db)
	NewCoreEntityController(db).Routes(router, db)
	NewCoreEntityFlagController(db).Routes(router, db)
}
