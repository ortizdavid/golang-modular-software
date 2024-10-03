package api

import "github.com/gofiber/fiber/v2"

func (api *EmployeeApi) getProfessionalInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	professionalInfo, err := api.professionalInfoService.GetByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(professionalInfo)
}

func (api *EmployeeApi) addProfessionalInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	professionalInfo, err := api.professionalInfoService.GetByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(professionalInfo)
}

func (api *EmployeeApi) updateProfessionalInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	professionalInfo, err := api.professionalInfoService.GetByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(professionalInfo)
}
