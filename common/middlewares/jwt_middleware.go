package middlewares

import (
	"github.com/gofiber/fiber/v2"
	//"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
)

type JwtMiddleware struct {
	userService *services.UserService
	roleService *services.RoleService
	secretKey string
}

func NewJwtMiddleware(db *database.Database) *JwtMiddleware {
	return &JwtMiddleware{
		userService: services.NewUserService(db),
		roleService: services.NewRoleService(db),
		secretKey:   config.GetEnv("JWT_SECRET_KEY"),
	}
}

func (mid *JwtMiddleware) AllowRoles(roles ...string) fiber.Handler {
	return func (c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return unauthorizedResponse(c, "Unauthorized. Auth Header missing")
		}
		
		user, err := mid.userService.GetUserById(c.Context(), 1)
		if err != nil {
			return unauthorizedResponse(c, "Unauthorized. User not found")
		}
		
		hasRoles, err := mid.userService.UserHasRoles(c.Context(), user.UserId, roles...) 
		if err != nil {
			return unauthorizedResponse(c, "Unauthorized. "+err.Error())
		}
		if !hasRoles {
			return unauthorizedResponse(c, "Access denied. You do not have the required role(s) to access this resource.")
		}
		
		return c.Next()
	}
}

func (mid *JwtMiddleware) GenerateJwtToken() (string, error) {
	return "", nil
}

func (mid *JwtMiddleware) RefreshToken() error {
	return nil
}

func (mid *JwtMiddleware) VerifyToken() error {
	return nil
}

