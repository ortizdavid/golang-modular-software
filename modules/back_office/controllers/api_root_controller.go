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
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>API Collection</title>
		</head>
		<body>
			<h1>Golang Modular Software API!</h1>
			<p>To test the API in Postman, download the API collections <a target='_blank' href='/api/download-collections'>Here</a>
			</p>
		</body>
		</html>`
	c.Set("Content-Type", "text/html")
	return c.SendString(html)
}

func (ctrl *ApiRootController) downloadCollections(c * fiber.Ctx) error {
	path := "./_api_collections/"
	fileName := "Golang-Modular-Software.postman_collection.json"
	c.Set("Content-Disposition", "attachment; filename="+fileName)
	c.Set("Content-Type", "application/json")
	return c.SendFile(filepath.Join(path, fileName))
}