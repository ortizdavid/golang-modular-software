package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func InitializeMiddlewares(router *fiber.App, db *database.Database) {
	router.Use(NewRequestLoggerMiddleware().Handle)
	router.Use(NewCorsMiddleware().Handle)
}