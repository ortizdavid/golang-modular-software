package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type ModuleFlagApi struct {
	service     *services.ModuleFlagService
	infoLogger  *helpers.Logger
	errorLogger *helpers.Logger
	shared.BaseController
}

func NewModuleFlagApi(db *database.Database) *ModuleFlagApi {
	return &ModuleFlagApi{
		service:     services.NewModuleFlagService(db),
		infoLogger:  helpers.NewInfoLogger(infoLogFile),
		errorLogger: helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *ModuleFlagApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/configurations/module-flags")
	group.Get("", api.getAll)
	group.Get("/by-module-id/:id", api.getByModuleId)
	group.Get("/by-module-code/:code", api.getByModuleCode)
}

func (api *ModuleFlagApi) getAll(c *fiber.Ctx) error {
	moduleFlags, err := api.service.GetAll(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(moduleFlags)
}

func (api *ModuleFlagApi) getByModuleId(c *fiber.Ctx) error {
	id := conversion.StringToInt(c.Params("id"))
	moduleFlag, err := api.service.GetByModuleId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(moduleFlag)
}

func (api *ModuleFlagApi) getByModuleCode(c *fiber.Ctx) error {
	code := c.Params("code")
	moduleFlag, err := api.service.GetByModuleCode(c.Context(), code)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(moduleFlag)
}
