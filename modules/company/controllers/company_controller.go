package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

type CompanyController struct {
}

func NewCompanyController(db *database.Database) *CompanyController {
	return &CompanyController{}
}

func (ctrl *CompanyController) Routes(router *fiber.App) {

}