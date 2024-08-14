package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
)

type ApiKeyMiddleware struct {
	userService *services.UserService
	roleService *services.RoleService
}

func NewApiKeyMiddleware(db *database.Database) *ApiKeyMiddleware {
	return &ApiKeyMiddleware{
		userService: services.NewUserService(db),
		roleService: services.NewRoleService(db),
	}
}

func (mid *ApiKeyMiddleware) AllowRoles(roles ...string) fiber.Handler {
	return func (c *fiber.Ctx) error {
		xUserId := c.Get("X-User-ID")
		if xUserId == "" {
			return unauthorizedResponse(c, "Unauthorized. User ID missing")
		}
		xApiKey := c.Get("X-API-Key")
		if xApiKey == "" {
			return unauthorizedResponse(c, "Unauthorized. API Key missing")
		}	
		user, err := mid.userService.GetUserById(c.Context(), conversion.StringToInt64(xUserId))
		if err != nil {
			return unauthorizedResponse(c, "Unauthorized. User not found")
		}
		userApiKey, err := mid.userService.GetUserApiKey(c.Context(), user.UserId)
		if err != nil {
			unauthorizedResponse(c, "Unauthorized. Error while getting User API Key")
		}
		if !userApiKey.IsActive {
			return unauthorizedResponse(c, "Unauthorized. API Key is disabled")
		}
		if xApiKey != userApiKey.Key {
			return unauthorizedResponse(c, "Unauthorized. Invalid API Key")
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

