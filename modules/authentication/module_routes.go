package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/api"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/controllers"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	flagMiddleware := middlewares.NewCoreEntityFlagFlagMiddleware(db)
	// Core Entity Routes and Codes
	userRoutes := []struct{
		Path string
		EntityCode string
	}{
		{"/users", entities.CoreEntityUsers.Code},
		{"/active-users", entities.CoreEntityActiveUsers.Code},
		{"/inactive-users", entities.CoreEntityInactiveUsers.Code},
		{"/online-users", entities.CoreEntityOnlineUsers.Code},
		{"/offline-users", entities.CoreEntityOfflineUsers.Code},
		{"/roles", entities.CoreEntityRoles.Code},
		{"/permissions", entities.CoreEntityPermissions.Code},
		{"/login-activities", entities.CoreEntityLoginActivity.Code},
	}
	// Register regular and API routes for each entity and apply flag
	for _, route := range userRoutes {
		router.Group("/user-management"+route.Path, flagMiddleware.CheckCoreEntityFlag(route.EntityCode))
		router.Group("/api/user-management"+route.Path, flagMiddleware.CheckCoreEntityFlag(route.EntityCode))
	}
	// Register additional controller and API routes
	controllers.RegisterControllerRoutes(router, db)
	api.RegisterApiRoutes(router, db)
}
