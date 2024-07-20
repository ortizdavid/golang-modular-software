package middlewares

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/models"
)


func authenticationMiddleware(ctx *fiber.Ctx) error {
	requestedPath := ctx.Path()
	if requestedPath == "/" ||
		strings.HasPrefix(requestedPath, "/images") ||
		strings.HasPrefix(requestedPath, "/css") ||
		strings.HasPrefix(requestedPath, "/js") ||
		strings.HasPrefix(requestedPath, "/auth") {
		return ctx.Next()
	}
	if !authentication.IsUserAuthenticated(ctx) {
		return ctx.Status(fiber.StatusUnauthorized).Render("errors/authentication", fiber.Map{
			"Title": "Autentication Error",
		})
	}
	return ctx.Next()
}
