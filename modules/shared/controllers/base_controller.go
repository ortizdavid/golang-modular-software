package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
)

type BaseController struct {
}

// HandleErrorsWeb centralizes error handling for web-related operations
func (ctrl *BaseController) HandleErrorsWeb(c *fiber.Ctx, err error) error {
	errMap := fiber.Map{
		"Title":   "Error",
		"Message": err.Error(),
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
