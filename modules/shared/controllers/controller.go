package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

type Controller interface {
	Routes(router *fiber.App, db *database.Database)
	index(fiber.Ctx) error
	details(fiber.Ctx) error
	createForm(fiber.Ctx) error
	create(fiber.Ctx) error
	editForm(fiber.Ctx) error
	edit(fiber.Ctx) error
	searchForm(fiber.Ctx) error
	search(fiber.Ctx) error
}