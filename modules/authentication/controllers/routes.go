package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterControllerRoutes(router *fiber.App, db *gorm.DB) {
	NewUserController(db).Routes(router)
	NewRoleController(db).Routes(router)
	NewAuthController(db).Routes(router)
}