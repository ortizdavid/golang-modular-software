package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterRoutes(router *fiber.App, db *database.Database) {
	NewCompanyApi(db).Routes(router, db)
	NewBranchApi(db).Routes(router, db)
	NewOfficeApi(db).Routes(router, db)
	NewDepartmentApi(db).Routes(router, db)
}