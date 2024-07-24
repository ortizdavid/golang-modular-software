package middlewares

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"gorm.io/gorm"
)

type AuthenticationMiddleware struct {
	service *authentication.AuthService
}

func NewAuthenticationMiddleware(db *gorm.DB) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		service: authentication.NewAuthService(db),
	}
}

func (mid *AuthenticationMiddleware) Handle(c *fiber.Ctx) error {
	requestedPath := c.Path()
	if requestedPath == "/" ||
		strings.HasPrefix(requestedPath, "/images") ||
		strings.HasPrefix(requestedPath, "/css") ||
		strings.HasPrefix(requestedPath, "/js") ||
		strings.HasPrefix(requestedPath, "/auth") {
		return c.Next()
	}
	if !mid.service.IsUserAuthenticated(c.Context(), c) {
		return c.Status(fiber.StatusUnauthorized).Render("errors/authentication", fiber.Map{
			"Title": "Authentication Error",
		})
	}
	return c.Next()
}
