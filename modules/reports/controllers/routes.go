package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterRoutes(router *fiber.App, db *database.Database) {
	NewRootController(db).Routes(router, db)
	NewUserReportController(db).Routes(router, db)
	NewConfigurationReportController(db).Routes(router, db)
	NewCompanyReportController(db).Routes(router, db)
	NewEmployeeReportController(db).Routes(router, db)
}
