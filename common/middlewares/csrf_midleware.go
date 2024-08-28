package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/ortizdavid/golang-modular-software/common/config"
)

type CsrfMiddleware struct {
	cfg csrf.Config
}

func NewCsrfMiddleware() *CsrfMiddleware { 
	return &CsrfMiddleware{
		cfg: csrf.Config{
			CookieSecure:  config.CsrfCookieSecure(),
			CookieHTTPOnly: true,
			Expiration:    time.Duration(config.CsrfExpiration()) * time.Hour,
			ErrorHandler: csrfErrorHandler,
		},
	}
}

func (mid * CsrfMiddleware) Handle() fiber.Handler {
	return csrf.New(mid.cfg)
}

func csrfErrorHandler(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusForbidden).Render("_errors/authorization", fiber.Map{
			"Title":   "CSRF Error",
			"Message": "CSRF token validation failed: " + err.Error(),
		})
	}
	return nil
}