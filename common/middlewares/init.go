package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func InitializeMiddlewares(app *fiber.App, db *database.Database) {
	app.Use(NewRequestLoggerMiddleware().Handle)
	app.Use(NewAuthenticationMiddleware(db).Handle)
	app.Use(NewCorsMiddleware().Handle)
}