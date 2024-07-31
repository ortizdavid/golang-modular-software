package helpers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/conversion"
)

type PaginationParam struct {
	CurrentPage int `json:"current_page"`
	Limit       int `json:"limit"`
}

func (p *PaginationParam) Validate() error {
	if p.CurrentPage < 0 {
		return errors.New("current_page must be greater or equal than 0")
	}
	if p.Limit < 0 {
		return errors.New("limit must be greater than 0")
	}
	return nil
}

func GetPaginationParams(c *fiber.Ctx) PaginationParam {
	currentStr := c.Query("current_page")
	limitStr := c.Query("limit")
	if currentStr == "" {
		currentStr = "0"
	}
	if limitStr == "" {
		limitStr = "10"
	}
	currentPage := conversion.StringToInt(currentStr)
	limit := conversion.StringToInt(limitStr)
	return PaginationParam{
		CurrentPage: currentPage,
		Limit:       limit,
	}
}