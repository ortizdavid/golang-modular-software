package middlewares

import "github.com/gofiber/fiber/v2"


func InitializeMiddlewares(app *fiber.App) {
	app.Use(requestLoggerMiddleware)
	app.Use(authenticationMiddleware)
}