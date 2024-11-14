package middlewares

import (
	"github.com/gofiber/fiber/v2"
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

func (mid *ApiKeyMiddleware) AllowRoles(roleCodes ...string) fiber.Handler {
	return func (c *fiber.Ctx) error {
		xUserId := c.Get("X-USER-ID")
		if xUserId == "" {
			return unauthorizedResponse(c, "Unauthorized. User ID missing")
		}
		xApiKey := c.Get("X-API-KEY")
		if xApiKey == "" {
			return unauthorizedResponse(c, "Unauthorized. API Key missing")
		}
		userApiKey, err := mid.userService.GetUserApiKey(c.Context(), xUserId)
		if err != nil {
			unauthorizedResponse(c, "Unauthorized. Error while getting User API Key")
		}	
		if !userApiKey.IsActive {
			return unauthorizedResponse(c, "Unauthorized. API Key is disabled")
		}
		if xApiKey != userApiKey.XApiKey {
			return unauthorizedResponse(c, "Unauthorized. Invalid API Key")
		}
		user, err := mid.userService.GetUserById(c.Context(), userApiKey.UserId)
		if err != nil {
			return unauthorizedResponse(c, "Unauthorized. User not found")
		}
		
		hasRoles, err := mid.userService.UserHasRoles(c.Context(), user.UserId, roleCodes...) 
		if err != nil {
			return unauthorizedResponse(c, "Unauthorized. "+err.Error())
		}
		if !hasRoles {
			return unauthorizedResponse(c, "Access denied. You do not have the required role(s) to access this resource.")
		}
		return c.Next()
	}
}

