package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
)

type CompanyApi struct {
	service *services.CompanyService
}

func NewCompanyApi(db *database.Database) *CompanyApi {
	return &CompanyApi{
		service: services.NewCompanyService(db),
	}
}

func (ctrl *CompanyApi) Routes(router *fiber.App, db *database.Database) {

}