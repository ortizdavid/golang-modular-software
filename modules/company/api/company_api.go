package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

type CompanyApi struct {
}

func NewCompanyApi(db *database.Database) *CompanyApi {
	return &CompanyApi{}
}

func (ctrl *CompanyApi) Routes(router *fiber.App) {

}