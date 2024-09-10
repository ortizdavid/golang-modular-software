package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterControllerRoutes(router *fiber.App, db *database.Database) {
	NewRootController(db).Routes(router, db)
	NewEmployeeController(db).Routes(router, db)
	NewJobTitleController(db).Routes(router, db)
	NewDocumentTypeController(db).Routes(router, db)
}