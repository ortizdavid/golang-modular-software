package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication"
	"github.com/ortizdavid/golang-modular-software/modules/back_office"
	"github.com/ortizdavid/golang-modular-software/modules/company"
	"github.com/ortizdavid/golang-modular-software/modules/configurations"
	entities "github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees"
	"github.com/ortizdavid/golang-modular-software/modules/references"
	"github.com/ortizdavid/golang-modular-software/modules/reports"
)

// RegisterRoutes initializes and registers controllers (routes) from different modules
func RegisterRoutes(router *fiber.App, db *database.Database) {
	// initialize middleware
	flagMiddleware := middlewares.NewModuleFlagMiddleware(db)

	// Define route groups and middleware checks
	routeGroups := []struct {
		prefix     string
		module     entities.ModuleInfo
		middlewares []fiber.Handler
	}{
		{
			prefix: "/user-management",
			module: entities.ModuleAuthentication,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleAuthentication.Code),
			},
		},
		{
			prefix: "/api/user-management",
			module: entities.ModuleAuthentication,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleAuthentication.Code),
			},
		},
		{
			prefix: "/configurations",
			module: entities.ModuleConfigurations,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleConfigurations.Code),
			},
		},
		{
			prefix: "/api/configurations",
			module: entities.ModuleConfigurations,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleConfigurations.Code),
			},
		},
		{
			prefix: "/company",
			module: entities.ModuleCompany,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleCompany.Code),
			},
		},
		{
			prefix: "/api/company",
			module: entities.ModuleCompany,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleCompany.Code),
			},
		},
		{
			prefix: "/employees",
			module: entities.ModuleEmployees,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleEmployees.Code),
			},
		},
		{
			prefix: "/api/employees",
			module: entities.ModuleEmployees,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleEmployees.Code),
			},
		},
		{
			prefix: "/references",
			module: entities.ModuleReferences,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleReferences.Code),
			},
		},
		{
			prefix: "/api/references",
			module: entities.ModuleReferences,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleReferences.Code),
			},
		},
	}

	// Register middleware for each route group
	for _, group := range routeGroups {
		router.Group(group.prefix, group.middlewares...)
		router.Group("/api"+group.prefix, group.middlewares...)
	}

	// Register module routes
	authentication.RegisterModuleRoutes(router, db)
	configurations.RegisterModuleRoutes(router, db)
	references.RegisterModuleRoutes(router, db)
	company.RegisterModuleRoutes(router, db)
	employees.RegisterModuleRoutes(router, db)
	back_office.RegisterModuleRoutes(router, db)
	reports.RegisterModuleRoutes(router, db)

}