package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterControllerRoutes(router *fiber.App, db *database.Database) {
	NewAuthController(db).Routes(router)
	NewRootController(db).Routes(router, db)
	NewAccountController(db).Routes(router, db)
	NewRoleController(db).Routes(router, db)
	NewUserController(db).Routes(router, db)
	NewPermissionController(db).Routes(router, db)
	NewLoginActivityController(db).Routes(router, db)
}
