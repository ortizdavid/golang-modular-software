package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

type EmployeeApi struct {
}

func NewEmployeeApi(db *database.Database) *EmployeeApi {
	return &EmployeeApi{}
}

func (ctrl *EmployeeApi) Routes(router *fiber.App) {

}