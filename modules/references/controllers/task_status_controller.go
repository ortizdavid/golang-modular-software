package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	"github.com/ortizdavid/golang-modular-software/modules/references/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type TaskStatusController struct {
	service        *services.TaskStatusService
	authService    *authentication.AuthService
	configService *configurations.AppConfigurationService
	flagStatusService *configurations.ModuleFlagStatusService
	infoLogger     *helpers.Logger
	errorLogger    *helpers.Logger
}

func NewTaskStatusController(db *database.Database) *TaskStatusController {
	return &TaskStatusController{
		service:           services.NewTaskStatusService(db),
		authService:       authentication.NewAuthService(db),
		configService:     configurations.NewAppConfigurationService(db),
		flagStatusService: configurations.NewModuleFlagStatusService(db),
		infoLogger:        helpers.NewInfoLogger("references-info.log"),
		errorLogger:       helpers.NewErrorLogger("references-error.log"),
	}
}

func (ctrl *TaskStatusController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/references/task-statuses", authMiddleware.CheckLoggedUser)
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

func (ctrl *TaskStatusController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllStatusesPaginated(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("references/task-status/index", fiber.Map{
		"Title":       "Task Statuses",
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"LoggedUser":  loggedUser,
		"Pagination":  pagination,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages":  pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *TaskStatusController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	status, err := ctrl.service.GetTaskStatusByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("references/task-status/details", fiber.Map{
		"Title":      "Details",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"LoggedUser": loggedUser,
		"TaskStatus":     status,
	})
}

func (ctrl *TaskStatusController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("references/task-status/create", fiber.Map{
		"Title":      "Create Task Status",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"LoggedUser": loggedUser,
	})
}

func (ctrl *TaskStatusController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateStatusRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.CreateTaskStatus(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Created task status '"+request.StatusName+"' successfully")
	return c.Redirect("/references/task-statuses")
}

func (ctrl *TaskStatusController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	status, err := ctrl.service.GetTaskStatusByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("references/task-status/edit", fiber.Map{
		"Title":      "Edit Task Status",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"LoggedUser": loggedUser,
		"TaskStatus":     status,
	})
}

func (ctrl *TaskStatusController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	status, err := ctrl.service.GetTaskStatusByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.UpdateStatusRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.UpdateTaskStatus(c.Context(), status.StatusId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Updated task status '"+request.StatusName+"' successfully")
	return c.Redirect("/references/task-statuses/" + id + "/details")
}

func (ctrl *TaskStatusController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	return c.Render("references/task-status/search", fiber.Map{
		"Title":      "Search Statuses",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
	})
}

func (ctrl *TaskStatusController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchStatusRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.SearchStatuses(c.Context(), c, request, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("references/task-status/search-results", fiber.Map{
		"Title":       "Search Results",
		"LoggedUser":  loggedUser,
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
		"Pagination":  pagination,
		"Param":       request.SearchParam,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages":  pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *TaskStatusController) removeForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	flagStatus, _ := ctrl.flagStatusService.LoadModuleFlagStatus(c.Context())
	status, err := ctrl.service.GetTaskStatusByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("references/task-status/delete", fiber.Map{
		"Title":  "Remove Task Status",
		"TaskStatus": 	status,
		"LoggedUser":  loggedUser,
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": flagStatus,
	})
}

func (ctrl *TaskStatusController) remove(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	status, err := ctrl.service.GetTaskStatusByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.RemoveTaskStatus(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' removed task status '%s'", loggedUser.UserName, status.StatusName))
	return c.Redirect("/references/task-statuses")
}
