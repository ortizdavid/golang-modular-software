package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type UserApi struct {
	service       *services.UserService
	roleService   *services.RoleService
	authService   *services.AuthService
	configService *configurations.BasicConfigurationService
	infoLogger    *helpers.Logger
	errorLogger   *helpers.Logger
}

func NewUserApi(db *database.Database) *UserApi {
	return &UserApi{
		service:       services.NewUserService(db),
		roleService:   services.NewRoleService(db),
		authService:   services.NewAuthService(db),
		configService: configurations.NewBasicConfigurationService(db),
		infoLogger:    helpers.NewInfoLogger("users-info.log"),
		errorLogger:   helpers.NewInfoLogger("users-error.log"),
	}
}

func (api *UserApi) Routes(router *fiber.App) {
	group := router.Group("/api/users")
	group.Get("", api.getAllUsers)
	router.Get("/active-users", api.getAllActiveUsers)
	router.Post("/inactive-users", api.getAllInactiveUsers)
}

func (ctrl *UserApi) getAllUsers(c *fiber.Ctx) error {
	_, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllUsers(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	return c.JSON(pagination)
}


func (ctrl *UserApi) getAllActiveUsers(c *fiber.Ctx) error {
	var params helpers.PaginationParam
	_, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	if err := c.QueryParser(&params); err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	pagination, err := ctrl.service.GetAllActiveUsers(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	return c.JSON(pagination)
}

func (ctrl *UserApi) getAllInactiveUsers(c *fiber.Ctx) error {
	var params helpers.PaginationParam
	_, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	if err := c.QueryParser(&params); err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	pagination, err := ctrl.service.GetAllInactiveUsers(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrorsApi(c, err)
	}
	return c.JSON(pagination)
}