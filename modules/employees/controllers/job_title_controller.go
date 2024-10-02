package controllers

import (
	"fmt"
    shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/services"
	
)

type JobTitleController struct {
	service                 *services.JobTitleService
	authService             *authentication.AuthService
	moduleFlagStatusService *configurations.ModuleFlagStatusService
	configService           *configurations.AppConfigurationService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
	shared.BaseController
}

func NewJobTitleController(db *database.Database) *JobTitleController {
	return &JobTitleController{
		service:                 services.NewJobTitleService(db),
		authService:             authentication.NewAuthService(db),
		moduleFlagStatusService: configurations.NewModuleFlagStatusService(db),
		configService:           configurations.NewAppConfigurationService(db),
		infoLogger:              helpers.NewInfoLogger(infoLogFile),
		errorLogger:             helpers.NewErrorLogger(errorLogFile),
	}
}

func (ctrl *JobTitleController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/employees/job-titles")
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
}

func (ctrl *JobTitleController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllPaginated(c.Context(), c, params)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/job-title/index", fiber.Map{
		"Title":            "Job Titles",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"Pagination":       pagination,
		"CurrentPage":      pagination.MetaData.CurrentPage + 1,
		"TotalPages":       pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *JobTitleController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	jobTitle, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/job-title/details", fiber.Map{
		"Title":            "Details",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"JobTitle":         jobTitle,
	})
}

func (ctrl *JobTitleController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("employee/job-title/create", fiber.Map{
		"Title":            "Create Job Title",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
	})
}

func (ctrl *JobTitleController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateJobTitleRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err := ctrl.service.Create(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Created job title '"+request.TitleName+"' successfully")
	return c.Redirect("/employees/job-titles")
}

func (ctrl *JobTitleController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	jobTitle, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employee/job-title/edit", fiber.Map{
		"Title":            "Edit Job Title",
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"LoggedUser":       loggedUser,
		"JobTitle":         jobTitle,
	})
}

func (ctrl *JobTitleController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	jobTitle, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.UpdateJobTitleRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.Update(c.Context(), jobTitle.JobTitleId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Updated job title '"+request.TitleName+"' successfully")
	return c.Redirect("/employees/job-titles/" + id + "/details")
}

func (ctrl *JobTitleController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("employee/job-title/search", fiber.Map{
		"Title":            "Search Job Titles",
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *JobTitleController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchJobTitleRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.Search(c.Context(), c, request, params)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("employee/job-title/search-results", fiber.Map{
		"Title":            "Search Results",
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"Pagination":       pagination,
		"Param":            request.SearchParam,
		"CurrentPage":      pagination.MetaData.CurrentPage + 1,
		"TotalPages":       pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *JobTitleController) removeForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	jobTitle, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("employees/job-title/delete", fiber.Map{
		"Title":            "Remove Job Title",
		"JobTitle":         jobTitle,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *JobTitleController) remove(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	jobTitle, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.service.Remove(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' removed job title '%s'", loggedUser.UserName, jobTitle.TitleName))
	return c.Redirect("/employees/job-titles")
}
