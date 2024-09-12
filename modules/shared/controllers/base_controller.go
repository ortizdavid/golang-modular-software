package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

type BaseController interface {
	Routes(router *fiber.App, db *database.Database)
}