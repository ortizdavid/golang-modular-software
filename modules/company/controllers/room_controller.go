package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type RoomController struct {
	service        *services.RoomService
	companyService *services.CompanyService
	branchService *services.BranchService
	authService    *authentication.AuthService
	configService *configurations.AppConfigurationService
	infoLogger     *helpers.Logger
	errorLogger    *helpers.Logger
}

func NewRoomController(db *database.Database) *RoomController {
	return &RoomController{
		service:        services.NewRoomService(db),
		companyService: services.NewCompanyService(db),
		branchService:  services.NewBranchService(db),
		authService:    authentication.NewAuthService(db),
		configService:  configurations.NewAppConfigurationService(db),
		infoLogger:     helpers.NewInfoLogger("company-info.log"),
		errorLogger:    helpers.NewErrorLogger("company-error.log"),
	}
}

func (ctrl *RoomController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewSessionAuthMiddleware(db)
	group := router.Group("/company/rooms", authMiddleware.CheckLoggedUser)
	group.Get("", ctrl.index)
	group.Get("/:id/details", ctrl.details)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
}

func (ctrl *RoomController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllCompaniesPaginated(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/room/index", fiber.Map{
		"Title":       "Rooms",
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser":  loggedUser,
		"Pagination":  pagination,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages":  pagination.MetaData.TotalPages + 1,
	})
}

func (ctrl *RoomController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	room, err := ctrl.service.GetRoomByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/room/details", fiber.Map{
		"Title":      "Details",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"Room":     room,
	})
}

func (ctrl *RoomController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	companies, err := ctrl.companyService.GetAllCompanies(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	branches, err := ctrl.branchService.GetAllBranches(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/room/create", fiber.Map{
		"Title":      "Create Room",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"Companies":  companies,
		"Branches": branches,
	})
}

func (ctrl *RoomController) create(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	var request entities.CreateRoomRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.CreateRoom(c.Context(), request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Created room '"+request.RoomName+"' successfully")
	return c.Redirect("/company/rooms")
}

func (ctrl *RoomController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	room, err := ctrl.service.GetRoomByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	companies, err := ctrl.companyService.GetAllCompanies(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	branches, err := ctrl.branchService.GetAllBranches(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("company/room/edit", fiber.Map{
		"Title":      "Edit Room",
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
		"LoggedUser": loggedUser,
		"Room":     room,
		"Companies":  companies,
		"Branches": branches,
	})
}

func (ctrl *RoomController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	room, err := ctrl.service.GetRoomByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.UpdateRoomRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.UpdateRoom(c.Context(), room.RoomId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User "+loggedUser.UserName+" Updated Room '"+request.RoomName+"' successfully")
	return c.Redirect("/company/rooms/" + id + "/details")
}

func (ctrl *RoomController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("company/room/search", fiber.Map{
		"Title":      "Search Rooms",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.configService.LoadAppConfigurations(c.Context()),
	})
}

func (ctrl *RoomController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchRoomRequest{SearchParam: searcParam}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.SearchRooms(c.Context(), c, request, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
	return c.Render("company/room/search-results", fiber.Map{
		"Title":       "Search Results",
		"LoggedUser":  loggedUser,
		"AppConfig":   ctrl.configService.LoadAppConfigurations(c.Context()),
		"Pagination":  pagination,
		"Param":       request.SearchParam,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages":  pagination.MetaData.TotalPages + 1,
	})
}