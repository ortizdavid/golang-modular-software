package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type UserController struct {
	service                 *services.UserService
	loginActivity           *services.LoginActivityService
	authService             *services.AuthService
	roleService             *services.RoleService
	configService           *configurations.AppConfigurationService
	moduleFlagStatusService *configurations.ModuleFlagStatusService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
	shared.BaseController
}

func NewUserController(db *database.Database) *UserController {
	return &UserController{
		service:                 services.NewUserService(db),
		loginActivity:           services.NewLoginActivityService(db),
		authService:             services.NewAuthService(db),
		roleService:             services.NewRoleService(db),
		configService:           configurations.NewAppConfigurationService(db),
		moduleFlagStatusService: configurations.NewModuleFlagStatusService(db),
		infoLogger:              helpers.NewInfoLogger(userInfoLogFile),
		errorLogger:             helpers.NewInfoLogger(userErrorLogFile),
	}
}

func (ctrl *UserController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/user-management/users")

	group.Get("/", ctrl.index)
	group.Get("/:id/details", ctrl.details)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)

	group.Get("/:id/assign-role", ctrl.assignRoleForm)
	group.Post("/:id/assign-role", ctrl.assignRole)
	group.Get("/:userId/remove-role/:userRoleId", ctrl.removeRoleForm)
	group.Post("/:userId/remove-role/:userRoleId", ctrl.removeRole)

	group.Get("/:id/deactivate", ctrl.deactivateForm)
	group.Post("/:id/deactivate", ctrl.deactivate)
	group.Get("/:id/activate", ctrl.activateForm)
	group.Post("/:id/activate", ctrl.activate)

	group.Get("/:id/reset-password", ctrl.resetPasswordForm)
	group.Post("/:id/reset-password", ctrl.resetPassword)

	group.Get("/active-users", ctrl.getActiveUsers)
	group.Get("/inactive-users", ctrl.getInactiveUsers)
	group.Get("/online-users", ctrl.getOnlineUsers)
	group.Get("/offline-users", ctrl.getOfflineUsers)
}
