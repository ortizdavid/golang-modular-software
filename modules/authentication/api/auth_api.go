package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
)

type AuthApi struct {
	service		*services.AuthService
	jwtService	*services.JwtService
	userService	*services.UserService
	infoLogger	*helpers.Logger
	errorLogger	*helpers.Logger
}

func NewAuthApi(db *database.Database) *AuthApi {
	return &AuthApi{
		service:     services.NewAuthService(db),
		jwtService:  services.NewJwtService(config.GetEnv("JWT_SECRET_KEY")),
		userService: services.NewUserService(db),
		infoLogger:  helpers.NewInfoLogger("auth-info.log"),
		errorLogger: helpers.NewInfoLogger("auth-error.log"),
	}
}

func (ctrl *AuthApi) Routes(router *fiber.App) {
	group := router.Group("/api/auth")
	group.Post("/login", ctrl.login)
	group.Post("/refresh-token", ctrl.refreshToken)
}

func (ctrl *AuthApi) login(c *fiber.Ctx) error {
	var request entities.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	token, err := ctrl.service.AuthenticateAPI(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, fmt.Sprintf("User '%s' failed to login", request.UserName))
		return helpers.HandleHttpErrorsApi(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' authenticated sucessful!", request.UserName))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"token":  token,
	})
}

func (ctrl *AuthApi) refreshToken(c *fiber.Ctx) error {
	// Extract refresh token from request (assuming it's in the Authorization header)
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return helpers.HandleHttpErrorsApi(c, apperrors.NewBadRequestError("Refresh token is required"))
	}
	// Validate and parse the refresh token
	parsedToken, claims, err := ctrl.jwtService.ParseToken(tokenString)
	if err != nil || !parsedToken.Valid {
		ctrl.errorLogger.Error(c, fmt.Sprintf("Invalid refresh token: %v", err))
		return helpers.HandleHttpErrorsApi(c, apperrors.NewUnauthorizedError("Invalid or expired refresh token"))
	}
	// Extract user ID from the token claims
	userId, ok := claims["user_id"].(int64)
	if !ok {
		ctrl.errorLogger.Error(c, "Failed to extract user ID from refresh token")
		return helpers.HandleHttpErrorsApi(c, apperrors.NewUnauthorizedError("Invalid refresh token"))
	}
	// Generate new access token
	newToken, err := ctrl.jwtService.GenerateJwtToken(userId)
	if err != nil {
		ctrl.errorLogger.Error(c, fmt.Sprintf("Error generating new token: %v", err))
		return helpers.HandleHttpErrorsApi(c, apperrors.NewInternalServerError("Error generating new token"))
	}
	// Log successful token refresh
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%d' refreshed token successfully", userId))
	// Return new token in response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"token":  newToken,
	})
}
