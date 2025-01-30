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
	skipPaths := []string{"/api", "/images", "/js", "/lib", "/auth/login"}
	
	if requestedPath == "/" {
		return c.Next()
	}
	for _, path := range skipPaths  {
		if strings.HasPrefix(requestedPath, path) {
			return c.Next()
		}
	} 
	if !mid.service.IsUserAuthenticated(c.Context(), c) {
		return c.Status(fiber.StatusUnauthorized).Render("_errors/authentication", fiber.Map{
			"Title": "Authentication Error",
		})
	}
	return c.Next()
}

func (mid *SessionAuthMiddleware)  CheckLoggedUser(c *fiber.Ctx) error {
    if strings.HasPrefix(c.Path(), "/api") {
        return c.Next() 
    }
	loggedUser, err := mid.service.GetLoggedUser(c.Context(), c)
	if err != nil || loggedUser.UserId == 0 || loggedUser.IsActive == "No" {
		return c.Redirect("/auth/login")
	}
	return c.Next()
}