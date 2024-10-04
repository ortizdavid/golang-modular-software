package api

import "github.com/gofiber/fiber/v2"

func (api *EmployeeApi) getDocuments(c *fiber.Ctx) error {
	id := c.Params("id")
	documents, err := api.documentService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(documents)
}

func (api *EmployeeApi) addDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	documents, err := api.documentService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(documents)
}

func (api *EmployeeApi) editDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	documents, err := api.documentService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(documents)
}
