package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/controllers"
)

type UserApi struct {
	service       *services.UserService
	activityService *services.LoginActivityService
	roleService   *services.RoleService
	authService   *services.AuthService
	configService *configurations.BasicConfigurationService
	infoLogger    *helpers.Logger
	errorLogger   *helpers.Logger
	shared.BaseController
}

func NewUserApi(db *database.Database) *UserApi {
	return &UserApi{
		service:       services.NewUserService(db),
		activityService: services.NewLoginActivityService(db),
		roleService:   services.NewRoleService(db),
		authService:   services.NewAuthService(db),
		configService: configurations.NewBasicConfigurationService(db),
		infoLogger:    helpers.NewInfoLogger("users-info.log"),
		errorLogger:   helpers.NewInfoLogger("users-error.log"),
	}
}

func (api *UserApi) Routes(router *fiber.App, db *database.Database) {
	group := router.Group("/api/user-management/users")
	
	group.Get("", api.getAllUsers)
	group.Get("/active-users", api.getActiveUsers)
	group.Get("/inactive-users", api.getInactiveUsers)
	group.Get("/online-users", api.getOnlineUsers)
	group.Get("/offline-users", api.getOfflineUsers)
	group.Get("/search/:param", api.search)

	group.Post("", api.create)
	group.Put("/:id", api.edit)
	
	group.Get("/:id", api.getByUniqueId)
	group.Get("/by-name/:name", api.getByName)
	group.Get("/by-email/:email", api.getByEmail)
	group.Get("/by-token/:token", api.getByToken)

	group.Get("/:id/roles", api.getUserRoles)
	group.Get("/:name/login-activity", api.getLoginActivityByName)
	group.Get("/:name/api-info", api.getApiInfoByName)

	group.Post("/:id/activate", api.activate)
	group.Post("/:id/deactivate", api.deactivate)
}

