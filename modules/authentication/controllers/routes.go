package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterControllerRoutes(router *fiber.App, db *database.Database) {
	NewUserController(db).Routes(router, db)
	NewRoleController(db).Routes(router)
	NewAuthController(db).Routes(router)
}