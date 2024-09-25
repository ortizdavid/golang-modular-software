package controllers

import (
	"html/template"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
)

type BaseController struct {
}

// HandleErrorsWeb centralizes error handling for web-related operations
func (ctrl *BaseController) HandleErrorsWeb(c *fiber.Ctx, err error) error {
	errMap := fiber.Map{
		"Title":   "Error",
		"Message": template.HTML(formatHtmlErrors(err.Error())),
	}
	if e, ok := err.(*apperrors.HttpError); ok {
		return c.Status(e.StatusCode).Render("_errors/error", errMap)
	}
	return c.Status(fiber.StatusInternalServerError).Render("_errors/error", errMap)
}

// HandleErrorsApi centralizes error handling for API-related operations
func (ctrl *BaseController) HandleErrorsApi(c *fiber.Ctx, err error) error {
	if e, ok := err.(*apperrors.HttpError); ok {
		return c.Status(e.StatusCode).JSON(fiber.Map{
			"error": e.Error(),
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func (ctrl *BaseController) HandleLoginError(c *fiber.Ctx, err error) error {
	errMap := fiber.Map{
		"Title":   "Login",
		"ErrorMessage": err.Error(),
	}
	return c.Status(fiber.StatusUnauthorized).Render("authentication/auth/login", errMap)
}

func formatHtmlErrors(errorMessage string) string {
	str := strings.ReplaceAll(errorMessage, "\n", "<br/>")
	return strings.ReplaceAll(str, "\t", "&nbsp;")
}