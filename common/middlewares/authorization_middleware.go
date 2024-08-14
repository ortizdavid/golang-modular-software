package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
)

type AuthorizationMiddleware struct {
	authService *authentication.AuthService
	userService *authentication.UserService
}

func NewAuthorizationMiddleware(db *database.Database) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{
		authService: authentication.NewAuthService(db),
		userService: authentication.NewUserService(db),
	}
}

// AllowRoles creates a middleware handler for role-based access control with HTML responses.
func (mid *AuthorizationMiddleware) AllowRoles(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the logged-in user
		loggedUser, err := mid.authService.GetLoggedUser(c.Context(), c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).Render("_errors/authorization", fiber.Map{
				"Title":   "Authentication Error",
				"Message": "User not authenticated: " + err.Error(),
			})
		}
		// Check if the user has the required roles
		hasRoles, err := mid.userService.UserHasRoles(c.Context(), loggedUser.UserId, roles...)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).Render("_errors/authorization", fiber.Map{
				"Title":   "Internal Server Error",
				"Message": "Error checking user roles: " + err.Error(),
			})
		}
		if !hasRoles {
			return c.Status(fiber.StatusForbidden).Render("_errors/authorization", fiber.Map{
				"Title":   "Access Denied",
				"Message": "You do not have the necessary roles to access this resource.",
			})
		}
		
		return c.Next()
	}
}

