package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterControllerRoutes(router *fiber.App, db *database.Database) {
	NewAuthController(db).Routes(router)
	NewRoleController(db).Routes(router, db)
	NewUserController(db).Routes(router, db)
}