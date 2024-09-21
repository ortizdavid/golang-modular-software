package company

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/api"
	"github.com/ortizdavid/golang-modular-software/modules/company/controllers"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	flagMiddleware := middlewares.NewCoreEntityFlagFlagMiddleware(db)
	// Core Entity Routes and Codes
	companyRoutes := []struct{
		Path string
		EntityCode string
	}{
		{"/company-info", entities.CoreEntityCompanyInfo.Code},
		{"/branches", entities.CoreEntityBranches.Code},
		{"/offices", entities.CoreEntityOffices.Code},
		{"/departments", entities.CoreEntityDepartments.Code},
		{"/rooms", entities.CoreEntityRooms.Code},
		{"/projects", entities.CoreEntityProjects.Code},
		{"/policies", entities.CoreEntityPolicies.Code},
	}
	// Register regular and API routes for each entity and apply Flags
	for _, route := range companyRoutes {
		router.Group("/company"+route.Path, flagMiddleware.CheckCoreEntityFlag(route.EntityCode))
		router.Group("/api/company"+route.Path, flagMiddleware.CheckCoreEntityFlag(route.EntityCode))
	}
	// Register additional controller and API routes
	controllers.RegisterRoutes(router, db)
	api.RegisterRoutes(router, db)
}