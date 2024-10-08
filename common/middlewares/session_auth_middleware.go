package middlewares

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
)

type SessionAuthMiddleware struct {
	service *authentication.AuthService
}

func NewSessionAuthMiddleware(db *database.Database) *SessionAuthMiddleware {
	return &SessionAuthMiddleware{
		service: authentication.NewAuthService(db),
	}
}

func (mid *SessionAuthMiddleware) AuthenticateRequests(c *fiber.Ctx) error {
	requestedPath := c.Path()
	if requestedPath == "/" ||
		strings.HasPrefix(requestedPath, "/api") ||
		strings.HasPrefix(requestedPath, "/images") ||
		strings.HasPrefix(requestedPath, "/css") ||
		strings.HasPrefix(requestedPath, "/js") ||
		strings.HasPrefix(requestedPath, "/lib") ||
		strings.HasPrefix(requestedPath, "/auth/login") {
		return c.Next()
	}
	if !mid.service.IsUserAuthenticated(c.Context(), c) {
		return c.Status(fiber.StatusUnauthorized).Render("_errors/authentication", fiber.Map{
			"Title": "Authentication Error",
		})
	}
	return c.Next()
}

func (mid *SessionAuthMiddleware)  CheckLoggedUser(c *fiber.Ctx) error {
	// Skip session check for API routes
    if strings.HasPrefix(c.Path(), "/api") {
        return c.Next() 
    }
	loggedUser, err := mid.service.GetLoggedUser(c.Context(), c)
	if err != nil || loggedUser.UserId == 0 || loggedUser.IsActive == "No" {
		return c.Redirect("/auth/login")
	}
	return c.Next()
}