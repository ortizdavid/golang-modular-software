package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	company "github.com/ortizdavid/golang-modular-software/modules/company/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	"github.com/ortizdavid/golang-modular-software/modules/employees/services"
	references "github.com/ortizdavid/golang-modular-software/modules/references/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type EmployeeController struct {
	service                 *services.EmployeeService
	jobTitleService         *services.JobTitleService
	documentService         *services.DocumentService
	professionalInfoService *services.ProfessionalInfoService
	phoneService            *services.EmployeePhoneService
	emailService            *services.EmployeeEmailService
	addressService			*services.AddressService
	accountService          *services.EmployeeAccountService
	documentTypeService     *services.DocumentTypeService
	identTypeService        *references.IdentificationTypeService
	countryService          *references.CountryService
	maritalStatusService    *references.MaritalStatusService
	employmentStatusService *references.EmploymentStatusService
	contactTypeService      *references.ContactTypeService
	departmentService       *company.DepartmentService
	authService             *authentication.AuthService
	userService				*authentication.UserService
	moduleFlagStatusService *configurations.ModuleFlagStatusService
	configService           *configurations.AppConfigurationService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
	shared.BaseController
}

func NewEmployeeController(db *database.Database) *EmployeeController {
	return &EmployeeController{
		service:                 services.NewEmployeeService(db),
		jobTitleService:         services.NewJobTitleService(db),
		documentService:         services.NewDocumentService(db),
		professionalInfoService: services.NewProfessionalInfoService(db),
		phoneService:            services.NewEmployeePhoneService(db),
		emailService:            services.NewEmployeeEmailService(db),
		addressService:          services.NewAddressService(db),
		accountService:          services.NewEmployeeAccountService(db),
		documentTypeService:     services.NewDocumentTypeService(db),
		identTypeService:        references.NewIdentificationTypeService(db),
		countryService:          references.NewCountryService(db),
		maritalStatusService:    references.NewMaritalStatusService(db),
		employmentStatusService: references.NewEmploymentStatusService(db),
		contactTypeService:      references.NewContactTypeService(db),
		departmentService:       company.NewDepartmentService(db),
		authService:             authentication.NewAuthService(db),
		userService:             authentication.NewUserService(db),
		moduleFlagStatusService: configurations.NewModuleFlagStatusService(db),
		configService:           configurations.NewAppConfigurationService(db),
		infoLogger:              helpers.NewInfoLogger(infoLogFile),
		errorLogger:             helpers.NewErrorLogger(errorLogFile),
	}
}

func (ctrl *EmployeeController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/employees/employee-info")

	group.Get("", ctrl.index)
	group.Get("/:id/details", ctrl.details)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
	group.Get("/:id/delete", ctrl.removeForm)
	group.Post("/:id/delete", ctrl.remove)

	group.Get("/:id/add-phone", ctrl.addPhoneForm)
	group.Post("/:id/add-phone", ctrl.addPhone)
	group.Get("/:empId/edit-phone/:contId", ctrl.editPhoneForm)
	group.Post("/:empId/edit-phone/:contId", ctrl.editPhone)

	group.Get("/:id/add-email", ctrl.addEmailForm)
	group.Post("/:id/add-email", ctrl.addEmail)
	group.Get("/:empId/edit-email/:contId", ctrl.editEmailForm)
	group.Post("/:empId/edit-email/:contId", ctrl.editEmail)

	group.Get("/:id/add-address", ctrl.addAddressForm)
	group.Post("/:id/add-address", ctrl.addAddress)
	group.Get("/:empId/edit-address/:addId", ctrl.editAddressForm)
	group.Post("/:empId/edit-address/:addId", ctrl.editAddress)

	group.Get("/:id/add-document", ctrl.addDocumentForm)
	group.Post("/:id/add-document", ctrl.addDocument)
	group.Get("/:empId/edit-document/:docId", ctrl.editDocumentForm)
	group.Post("/:empId/edit-document/:docId", ctrl.editDocument)
	group.Get("/display-document/:id", ctrl.displayDocument)

	group.Get("/:id/add-professional-info", ctrl.addProfessionalInfoForm)
	group.Post("/:id/add-professional-info", ctrl.addProfessionalInfo)
	group.Get("/:empId/edit-professional-info/:proId", ctrl.editProfessionalInfoForm)
	group.Post("/:empId/edit-professional-info/:proId", ctrl.editProfessionalInfo)

	group.Get("/:id/add-user-account", ctrl.addUserAccountForm)
	group.Post("/:id/add-user-account", ctrl.addUserAccount)
	group.Get("/:id/associate-user-account", ctrl.associateUserAccountForm)
	group.Post("/:id/associate-user-account", ctrl.associateUserAccount)
}
