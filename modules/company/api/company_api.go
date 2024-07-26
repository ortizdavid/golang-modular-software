package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CompanyApi struct {
}

func NewCompanyApi(db *gorm.DB) *CompanyApi {
	return &CompanyApi{}
}

func (ctrl *CompanyApi) Routes(router *fiber.App) {

}