package controllers

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type ApiRootController struct {
}

func (ctrl ApiRootController) Routes(router *fiber.App) {
	group := router.Group("/api")
	group.Get("", ctrl.index)
	group.Get("/download-collections", ctrl.downloadCollections)
}

func (ctrl ApiRootController) index(c *fiber.Ctx) error {
	return c.Render("_back_office/api-root", fiber.Map{
		"Title": "API Collection",
	})
}

func (ctrl *ApiRootController) downloadCollections(c *fiber.Ctx) error {
	path := "./docs/_api/"
	fileName := "Golang Modular Software.postman_collection.json"
	c.Set("Content-Disposition", "attachment; filename="+fileName)
	c.Set("Content-Type", "application/json")
	return c.SendFile(filepath.Join(path, fileName))
}
