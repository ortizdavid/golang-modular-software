package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	"github.com/ortizdavid/golang-modular-software/modules/references/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type CurrencyApi struct {
	service			*services.CurrencyService
	infoLogger		*helpers.Logger
	errorLogger		*helpers.Logger
	shared.BaseController
}

func NewCurrencyApi(db *database.Database) *CurrencyApi {
	return &CurrencyApi{
		service:        services.NewCurrencyService(db),
		infoLogger:     helpers.NewInfoLogger(infoLogFile),
		errorLogger:    helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *CurrencyApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/references/currencies")
	group.Get("", api.getAll)
	group.Post("", api.add)
	group.Put("/:id", api.edit)
	group.Get("/not-paginated/", api.getAllNotPaginated)
	group.Get("/by-uuid/:id", api.getByUniqueId)
	group.Get("/by-name/:name", api.getByCurrencyName)
	group.Get("/by-code/:code", api.getByCode)
	
}

func (api *CurrencyApi) getAll(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	countries, err := api.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(countries)
}

func (api *CurrencyApi) add(c *fiber.Ctx) error {
	var request entities.CreateCurrencyRequest
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err := api.service.Create(c.Context(), request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Currency '"+request.CurrencyName+"' added"
	api.infoLogger.Info(c, msg)
	return c.JSON(msg)
}

func (api *CurrencyApi) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	country, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateCurrencyRequest
	if err := c.BodyParser(&request); err != nil {
		return api.HandleErrorsApi(c, err)
	}
	err = api.service.Update(c.Context(), id, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return api.HandleErrorsApi(c, err)
	}
	msg := "Currency '"+country.CurrencyName+"' edited"
	api.infoLogger.Info(c, msg)
	return c.JSON(msg)
}


func (api *CurrencyApi) getAllNotPaginated(c *fiber.Ctx) error {
	countries, err := api.service.GetAll(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(countries)
}

func (api *CurrencyApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	country, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(country)
}

func (api *CurrencyApi) getByCode(c *fiber.Ctx) error {
	isoCode := c.Params("code")
	country, err := api.service.GetByCode(c.Context(), isoCode)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(country)
}

func (api *CurrencyApi) getByCurrencyName(c *fiber.Ctx) error {
	name := c.Params("name")
	country, err := api.service.GetByName(c.Context(), name)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(country)
}