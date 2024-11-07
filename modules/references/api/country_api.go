package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	"github.com/ortizdavid/golang-modular-software/modules/references/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type CountryApi struct {
	service     *services.CountryService
	infoLogger  *helpers.Logger
	errorLogger *helpers.Logger
	shared.BaseController
}

func NewCountryApi(db *database.Database) *CountryApi {
	return &CountryApi{
		service:     services.NewCountryService(db),
		infoLogger:  helpers.NewInfoLogger(infoLogFile),
		errorLogger: helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *CountryApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/references/countries")
	group.Get("", api.getAll)
	group.Post("", api.add)
	group.Put("/:id", api.edit)
	group.Get("/not-paginated", api.getAllNotPaginated)
	group.Get("/by-uuid/:id", api.getByUniqueId)
	group.Get("/by-name/:name", api.getByCountryName)
	group.Get("/by-iso-code/:code", api.getByIsoCode)
}

func (api *CountryApi) getAll(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	countries, err := api.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(countries)
}

func (api *CountryApi) add(c *fiber.Ctx) error {
	var request entities.CreateCountryRequest
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err := api.service.Create(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Country '" + request.CountryName + "' added"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}

func (api *CountryApi) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	country, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateCountryRequest
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.service.Update(c.Context(), id, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Country '" + country.CountryName + "' edited"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"message": msg})
}

func (api *CountryApi) getAllNotPaginated(c *fiber.Ctx) error {
	countries, err := api.service.GetAll(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(countries)
}

func (api *CountryApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	country, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(country)
}

func (api *CountryApi) getByIsoCode(c *fiber.Ctx) error {
	isoCode := c.Params("code")
	country, err := api.service.GetByIsoCode(c.Context(), isoCode)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(country)
}

func (api *CountryApi) getByCountryName(c *fiber.Ctx) error {
	name := c.Params("name")
	country, err := api.service.GetByName(c.Context(), name)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(country)
}
