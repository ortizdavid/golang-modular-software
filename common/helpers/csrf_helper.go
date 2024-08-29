package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/config"
)

func GenerateCsrfToken(c *fiber.Ctx) string {
	store := config.GetSessionStore()
	session, err := store.Get(c)
	if err != nil {
		return ""
	}
	csrfToken := encryption.GenerateRandomToken()
	session.Set("csrf_token", csrfToken)
	session.Save()
	return csrfToken
}

func ValidateCsrfToken(c *fiber.Ctx, requestToken any) bool {
	store := config.GetSessionStore()
	session, err := store.Get(c)
	if err != nil {
		return false
	}
	storedToken := session.Get("csrf_token")
	return storedToken == requestToken
}

func CheckCsrfToken(c *fiber.Ctx, token string) error {
	if !ValidateCsrfToken(c, token) {
		return c.Status(fiber.StatusForbidden).Render("_errors/error", fiber.Map{
			"Title":   "CSRF Error",
			"Message": "CSRF token validation failed ",
		})
	}
	return nil
}

