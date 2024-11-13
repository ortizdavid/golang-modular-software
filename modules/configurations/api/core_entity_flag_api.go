package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type CoreEntityFlagApi struct {
	service     *services.CoreEntityFlagService
	infoLogger  *helpers.Logger
	errorLogger *helpers.Logger
	shared.BaseController
}

func NewCoreEntityFlagApi(db *database.Database) *CoreEntityFlagApi {
	return &CoreEntityFlagApi{
		service:     services.NewCoreEntityFlagService(db),
		infoLogger:  helpers.NewInfoLogger(infoLogFile),
		errorLogger: helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *CoreEntityFlagApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/configurations/core-entity-flags")
	group.Get("", api.getAll)
	group.Get("/by-entity-id/:id", api.getByCoreEntityId)
	group.Get("/by-entity-code/:code", api.getByCoreEntityCode)
	group.Get("/by-module-id/:id", api.getAllByModuleId)
	group.Get("/by-module-code/:code", api.getAllByModuleCode)
}

func (api *CoreEntityFlagApi) getAll(c *fiber.Ctx) error {
	coreEntityFlags, err := api.service.GetAll(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(coreEntityFlags)
}

func (api *CoreEntityFlagApi) getByCoreEntityId(c *fiber.Ctx) error {
	id := conversion.StringToInt(c.Params("id"))
	coreEntityFlag, err := api.service.GetByCoreEntityId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(coreEntityFlag)
}

func (api *CoreEntityFlagApi) getByCoreEntityCode(c *fiber.Ctx) error {
	code := c.Params("code")
	coreEntityFlag, err := api.service.GetByCoreEntityCode(c.Context(), code)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(coreEntityFlag)
}

func (api *CoreEntityFlagApi) getAllByModuleId(c *fiber.Ctx) error {
	id := conversion.StringToInt(c.Params("id"))
	moduleFlags, err := api.service.GetAllByModuleId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(moduleFlags)
}

func (api *CoreEntityFlagApi) getAllByModuleCode(c *fiber.Ctx) error {
	code := c.Params("code")
	moduleFlags, err := api.service.GetAllByModuleCode(c.Context(), code)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(moduleFlags)
}
