package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type EmployeeApi struct {
	service					*services.EmployeeService
	completeDataService			*services.EmployeeCompleteDataService
	documentService         *services.DocumentService
	professionalInfoService *services.ProfessionalInfoService
	phoneService            *services.EmployeePhoneService
	emailService            *services.EmployeeEmailService
	addressService			*services.AddressService
	accountService          *services.EmployeeAccountService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
	shared.BaseController
}

func NewEmployeeApi(db *database.Database) *EmployeeApi {
	return &EmployeeApi{
		service:                 services.NewEmployeeService(db),
		completeDataService:         services.NewEmployeeCompleteDataService(db),
		documentService:         services.NewDocumentService(db),
		professionalInfoService: services.NewProfessionalInfoService(db),
		phoneService:            services.NewEmployeePhoneService(db),
		emailService:            services.NewEmployeeEmailService(db),
		addressService:          services.NewAddressService(db),
		accountService:          services.NewEmployeeAccountService(db),
		infoLogger:              helpers.NewInfoLogger(infoLogFile),
		errorLogger:             helpers.NewErrorLogger(errorLogFile),
		BaseController:          shared.BaseController{},
	}
}

func (api *EmployeeApi) Routes(router *fiber.App) {
	group := router.Group("/api/employees/employee-info")
	group.Get("", api.getAll)
	group.Get("/:id", api.getByUniqueId)
	group.Get("/by-identification/:identNumber", api.getByIdentificationNumber)
	group.Get("/:id/personal-info", api.getPersonalInfo)
	group.Get("/:id/professional-info", api.getProfessionalInfo)
	group.Get("/:id/addresses", api.getAddresses)
	group.Get("/:id/documents", api.getDocuments)
	group.Get("/:id/emails", api.getEmails)
	group.Get("/:id/phones", api.getPhones)
	group.Get("/:id/account", api.getAccount)
}

func (api *EmployeeApi) getAll(c *fiber.Ctx) error {
	params := helpers.GetPaginationParams(c)
	employees, err := api.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(employees)
}

func (api *EmployeeApi) getByUniqueId(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := api.completeDataService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(employee)
}

func (api *EmployeeApi) getByIdentificationNumber(c *fiber.Ctx) error {
	identNumber := c.Params("identNumber")
	employee, err := api.completeDataService.GetByIdentificationNumber(c.Context(), identNumber)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(employee)
}

func (api *EmployeeApi) getPersonalInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	personalInfo, err := api.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(personalInfo)
}

func (api *EmployeeApi) getProfessionalInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	professionalInfo, err := api.professionalInfoService.GetByEmployeeUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(professionalInfo)
}

func (api *EmployeeApi) getAddresses(c *fiber.Ctx) error {
	id := c.Params("id")
	addresses, err := api.completeDataService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(addresses)
}

func (api *EmployeeApi) getDocuments(c *fiber.Ctx) error {
	id := c.Params("id")
	documents, err := api.completeDataService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(documents)
}

func (api *EmployeeApi) getEmails(c *fiber.Ctx) error {
	id := c.Params("id")
	emails, err := api.completeDataService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(emails)
}

func (api *EmployeeApi) getPhones(c *fiber.Ctx) error {
	id := c.Params("id")
	phones, err := api.completeDataService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(phones)
}

func (api *EmployeeApi) getAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	account, err := api.completeDataService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return api.HandleErrorsApi(c, err)
	}
	return c.JSON(account)
}
