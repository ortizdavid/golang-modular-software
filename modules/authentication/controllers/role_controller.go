package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type RoleController struct {
	service                 *services.RoleService
	authService             *services.AuthService
	permissionService       *services.PermissionService
	configService           *configurations.AppConfigurationService
	moduleFlagStatusService *configurations.ModuleFlagStatusService
	infoLogger              *helpers.Logger
	errorLogger             *helpers.Logger
	shared.BaseController
}

func NewRoleController(db *database.Database) *RoleController {
	return &RoleController{
		service:                 services.NewRoleService(db),
		authService:             services.NewAuthService(db),
		permissionService:       services.NewPermissionService(db),
		configService:           configurations.NewAppConfigurationService(db),
		moduleFlagStatusService: configurations.NewModuleFlagStatusService(db),
		infoLogger:              helpers.NewInfoLogger("users-info.log"),
		errorLogger:             helpers.NewInfoLogger("users-error.log"),
	}
}

func (ctrl *RoleController) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/user-management/roles")

	group.Get("/", ctrl.index)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/details", ctrl.details)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/:id/delete", ctrl.deleteForm)
	group.Post("/:id/delete", ctrl.delete)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
	group.Get("/:id/assign-permission", ctrl.assignPermissionForm)
	group.Post("/:id/assign-permission", ctrl.assignPermission)
	group.Get("/:roleId/remove-permission/:permissionRoleId", ctrl.removePermissionForm)
	group.Post("/:roleId/remove-permission/:permissionRoleId", ctrl.removePermission)

}
