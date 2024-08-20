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
	NewRoomController(db).Routes(router, db)
	NewDepartmentController(db).Routes(router, db)
	NewPolicyController(db).Routes(router, db)
	NewProjectController(db).Routes(router, db)
}