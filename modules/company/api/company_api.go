package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type CompanyApi struct {
	service *services.CompanyService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
	shared.BaseController
}

func NewCompanyApi(db *database.Database) *CompanyApi {
	return &CompanyApi{
		service: services.NewCompanyService(db),
		infoLogger:              helpers.NewInfoLogger(infoLogFile),
		errorLogger:             helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *CompanyApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/company/company-info")
	group.Get("", api.getInfo)
	group.Get("/:id", api.getByUniqueId)
	group.Put("/:id", api.edit)
}

func (api *CompanyApi) getInfo(c *fiber.Ctx) error {
	company, err := api.service.GetCurrent(c.Context())
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(company)
}

func (api *CompanyApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	company, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(company)
}

func (api *CompanyApi) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	company, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	var request entities.UpdateCompanyRequest
	if err := c.BodyParser(&request); err != nil {
		return  api.HandleErrorsApi(c, err)
	}
	err =  api.service.Update(c.Context(), id, request)
	if err != nil {
		api.errorLogger.Error(c, err.Error())
		return  api.HandleErrorsApi(c, err)
	}
	msg := "Updated Company '"+company.CompanyName+"' successfully"
	api.infoLogger.Info(c, msg)
	return c.JSON(fiber.Map{"msg": msg})
}