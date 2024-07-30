package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)


func RegisterRoutes(router *fiber.App, db *database.Database) {
	NewBasicConfigurationController(db).Routes(router)
	NewCompanyConfigurationController(db).Routes(router)
	NewEmailConfigurationController(db).Routes(router)
}