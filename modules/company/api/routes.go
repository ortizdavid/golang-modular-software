package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterApiRoutes(router *fiber.App, db *gorm.DB) {
	NewCompanyApi(db).Routes(router)
}