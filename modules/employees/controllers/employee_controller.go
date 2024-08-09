package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

type EmployeeController struct {
}

func NewEmployeeController(db *database.Database) *EmployeeController {
	return &EmployeeController{}
}

func (ctrl *EmployeeController) Routes(router *fiber.App, db *database.Database) {
}
