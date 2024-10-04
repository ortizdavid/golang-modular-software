package api

import "github.com/gofiber/fiber/v2"

func (api *EmployeeApi) getEmails(c *fiber.Ctx) error {
	id := c.Params("id")
	emails, err := api.emailService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(emails)
}

func (api *EmployeeApi) addEmail(c *fiber.Ctx) error {
	id := c.Params("id")
	emails, err := api.emailService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(emails)
}

func (api *EmployeeApi) editEmail(c *fiber.Ctx) error {
	id := c.Params("id")
	emails, err := api.emailService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(emails)
}

func (api *EmployeeApi) getPhones(c *fiber.Ctx) error {
	id := c.Params("id")
	phones, err := api.phoneService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(phones)
}

func (api *EmployeeApi) addPhone(c *fiber.Ctx) error {
	id := c.Params("id")
	phones, err := api.phoneService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(phones)
}

func (api *EmployeeApi) editPhone(c *fiber.Ctx) error {
	id := c.Params("id")
	phones, err := api.phoneService.GetAllByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(phones)
}

