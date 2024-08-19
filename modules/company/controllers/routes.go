package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterControllerRoutes(router *fiber.App, db *database.Database) {
	NewRootController(db).Routes(router, db)
	NewCompanyController(db).Routes(router, db)
	NewBranchController(db).Routes(router, db)
	NewOfficeController(db).Routes(router, db)
}