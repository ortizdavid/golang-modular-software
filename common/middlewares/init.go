package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func InitializeMiddlewares(app *fiber.App, db *gorm.DB) {
	app.Use(NewRequestLoggerMiddleware().Handle)
	app.Use(NewAuthenticationMiddleware(db).Handle)
}