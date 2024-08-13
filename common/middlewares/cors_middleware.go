package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// CorsMiddleware struct for CORS handling
type CorsMiddleware struct {
	AllowedOrigins   string
	AllowedMethods   string
	AllowedHeaders   string
	ExposedHeaders   string
	AllowCredentials bool
	MaxAge           int
}

// NewCorsMiddleware creates a new instance of CorsMiddleware with default values
func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{
		AllowedOrigins:   "*",
		AllowedMethods:   "GET, POST, PUT, DELETE, OPTIONS",
		AllowedHeaders:   "Origin, Content-Type, Accept, Authorization",
		ExposedHeaders:   "",
		AllowCredentials: true,
		MaxAge:           0,
	}
}

// Handle processes the CORS middleware
func (cors *CorsMiddleware) Handle(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", cors.AllowedOrigins)
	c.Set("Access-Control-Allow-Methods", cors.AllowedMethods)
	c.Set("Access-Control-Allow-Headers", cors.AllowedHeaders)

	if cors.ExposedHeaders != "" {
		c.Set("Access-Control-Expose-Headers", cors.ExposedHeaders)
	}

	if cors.AllowCredentials {
		c.Set("Access-Control-Allow-Credentials", "true")
		// When AllowCredentials is true, Access-Control-Allow-Origin must not be "*"
		if cors.AllowedOrigins == "*" {
			c.Set("Access-Control-Allow-Origin", c.Get("Origin"))
		}
	}

	if cors.MaxAge > 0 {
		c.Set("Access-Control-Max-Age", strconv.Itoa(cors.MaxAge))
	}

	// Handle preflight requests
	if c.Method() == fiber.MethodOptions {
		return c.SendStatus(fiber.StatusNoContent)
	}

	return c.Next()
}
