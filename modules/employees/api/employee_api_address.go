package api

import "github.com/gofiber/fiber/v2"

func (api *EmployeeApi) getAddresses(c *fiber.Ctx) error {
	id := c.Params("id")
	addresses, err := api.addressService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(addresses)
}

func (api *EmployeeApi) addAddress(c *fiber.Ctx) error {
	id := c.Params("id")
	addresses, err := api.addressService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(addresses)
}

func (api *EmployeeApi) editAddress(c *fiber.Ctx) error {
	id := c.Params("id")
	addresses, err := api.addressService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(addresses)
}


