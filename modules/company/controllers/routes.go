package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterControllerRoutes(router *fiber.App, db *gorm.DB) {
	NewCompanyController(db).Routes(router)
}