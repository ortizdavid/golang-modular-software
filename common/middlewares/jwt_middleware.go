package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
)

type JwtMiddleware struct {
	userService *services.UserService
	roleService *services.RoleService
	jwtService  *services.JwtService
}

func NewJwtMiddleware(db *database.Database) *JwtMiddleware {
	jwtService := services.NewJwtService(config.JwtSecretKey())
	return &JwtMiddleware{
		userService: services.NewUserService(db),
		roleService: services.NewRoleService(db),
		jwtService:  jwtService,
	}
}

// AllowRoles checks if the user has the required roles to access a resource
func (mid *JwtMiddleware) AllowRoles(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Step 1: Extract JWT from Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return unauthorizedResponse(c, "Unauthorized. Authorization header missing")
		}
		
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return unauthorizedResponse(c, "Unauthorized. Invalid token format")
		}
		
		_, claims, err := mid.jwtService.ParseToken(tokenString)
		if err != nil {
			return unauthorizedResponse(c, "Unauthorized. Invalid or expired token")
		}
		
		userId, ok := claims["user_id"].(float64) // JWT library stores numbers as float64
		if !ok {
			return unauthorizedResponse(c, "Unauthorized. Invalid user ID in token")
		}
		
		user, err := mid.userService.GetUserById(c.Context(), int64(userId))
		if err != nil {
			return unauthorizedResponse(c, "Unauthorized. User not found")
		}
		
		hasRoles, err := mid.userService.UserHasRoles(c.Context(), user.UserId, roles...)
		if err != nil {
			return unauthorizedResponse(c, "Unauthorized. " + err.Error())
		}
		if !hasRoles {
			return unauthorizedResponse(c, "Access denied. You do not have the required role(s) to access this resource.")
		}
	
		return c.Next()
	}
}
