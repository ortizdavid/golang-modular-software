package references

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/references/api"
	"github.com/ortizdavid/golang-modular-software/modules/references/controllers"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	flagMiddleware := middlewares.NewCoreEntityFlagFlagMiddleware(db)
	// Core Entity Routes and Codes
	reportRoutes := []struct{
		Path string
		EntityCode string
	}{
		{"/users", entities.CoreEntityUserReports.Code},
		{"/configuration", entities.CoreEntityConfigurationReports.Code},
		{"/company", entities.CoreEntityCompanyReports.Code},
		{"/employees", entities.CoreEntityEmployeeReports.Code},
	}
	// Register regular and API routes for each entity and apply Flags
	for _, route := range reportRoutes {
		router.Group("/reports"+route.Path, flagMiddleware.CheckCoreEntityFlag(route.EntityCode))
		router.Group("/api/reports"+route.Path, flagMiddleware.CheckCoreEntityFlag(route.EntityCode))
	}
	// Register additional controller and API routes
	controllers.RegisterRoutes(router, db)
	api.RegisterApiRoutes(router, db)
}