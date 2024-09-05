package employees

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/api"
	"github.com/ortizdavid/golang-modular-software/modules/employees/controllers"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	flagMiddleware := middlewares.NewCoreEntityFlagFlagMiddleware(db)
	// Core Entity Routes and Codes
	employeeRoutes := []struct{
		Path string
		EntityCode string
	}{
		{"/employee-info", entities.CoreEntityEmployees.Code},
		{"/job-titles", entities.CoreEntityJobTitles.Code},
	}
	// Register regular and API routes for each entity and apply Flags
	for _, route := range employeeRoutes {
		router.Group("/employees"+route.Path, flagMiddleware.CheckCoreEntityFlag(route.EntityCode))
		router.Group("/api/employees"+route.Path, flagMiddleware.CheckCoreEntityFlag(route.EntityCode))
	}
	// Register additional controller and API routes
	controllers.RegisterControllerRoutes(router, db)
	api.RegisterApiRoutes(router, db)
}