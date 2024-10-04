package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/employees/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type EmployeeApi struct {
	service					*services.EmployeeService
	completeDataService		*services.EmployeeCompleteDataService
	documentService         *services.DocumentService
	professionalInfoService *services.ProfessionalInfoService
	phoneService            *services.EmployeePhoneService
	emailService            *services.EmployeeEmailService
	addressService			*services.AddressService
	accountService          *services.EmployeeAccountService
	userService 			*authentication.UserService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
	shared.BaseController
}

func NewEmployeeApi(db *database.Database) *EmployeeApi {
	return &EmployeeApi{
		service:                 services.NewEmployeeService(db),
		completeDataService:     services.NewEmployeeCompleteDataService(db),
		documentService:         services.NewDocumentService(db),
		professionalInfoService: services.NewProfessionalInfoService(db),
		phoneService:            services.NewEmployeePhoneService(db),
		emailService:            services.NewEmployeeEmailService(db),
		addressService:          services.NewAddressService(db),
		accountService:          services.NewEmployeeAccountService(db),
		userService:             authentication.NewUserService(db),
		infoLogger:              helpers.NewInfoLogger(infoLogFile),
		errorLogger:             helpers.NewErrorLogger(errorLogFile),
	}
}

func (api *EmployeeApi) Routes(router *fiber.App) {
	group := router.Group("/api/employees/employee-info")
	
	group.Get("", api.getAll)
	group.Post("", api.create)
	group.Put("/:id", api.edit)
	group.Get("/:id", api.getByUniqueId)
	group.Get("/by-identification/:identNumber", api.getByIdentificationNumber)
	group.Get("/search-results", api.search)
	group.Get("/:id/personal-info", api.getPersonalInfo)

	group.Get("/:id/professional-info", api.getProfessionalInfo)
	group.Post("/:id/professional-info", api.addProfessionalInfo)
	group.Put("/:id/professional-info", api.editProfessionalInfo)

	group.Get("/:id/addresses", api.getAddresses)
	group.Post("/:id/addresses", api.addAddress)
	group.Put("/:id/addresses", api.editAddress)

	group.Get("/:id/documents", api.getDocuments)
	group.Post("/:id/documents", api.addDocument)
	group.Put("/:id/documents", api.editDocument)

	group.Get("/:id/emails", api.getEmails)
	group.Post("/:id/emails", api.addEmail)
	group.Put("/:id/emails", api.editEmail)

	group.Get("/:id/phones", api.getPhones)
	group.Post("/:id/phones", api.addPhone)
	group.Put("/:id/phones", api.editPhone)
	
	group.Get("/:id/account", api.getUserAccount)
	group.Post("/:id/account", api.addUserAccount)
	group.Post("/:id/account-association", api.associateUserAccount)
}

