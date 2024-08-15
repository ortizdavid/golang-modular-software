package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterApiRoutes(router *fiber.App, db *database.Database) {
	NewAuthApi(db).Routes(router)
	NewUserApi(db).Routes(router, db)
	NewRoleApi(db).Routes(router, db)
}