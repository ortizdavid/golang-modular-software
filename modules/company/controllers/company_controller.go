package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CompanyController struct {
}

func NewCompanyController(db *gorm.DB) *CompanyController {
	return &CompanyController{}
}

func (ctrl *CompanyController) Routes(router *fiber.App) {

}