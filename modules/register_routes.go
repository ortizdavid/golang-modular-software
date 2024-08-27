package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication"
	"github.com/ortizdavid/golang-modular-software/modules/back_office"
	"github.com/ortizdavid/golang-modular-software/modules/company"
	"github.com/ortizdavid/golang-modular-software/modules/configurations"
	"github.com/ortizdavid/golang-modular-software/modules/employees"
	"github.com/ortizdavid/golang-modular-software/modules/references"
	"github.com/ortizdavid/golang-modular-software/modules/reports"
	entities "github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

// RegisterRoutes initializes and registers controllers (routes) from different modules
func RegisterRoutes(router *fiber.App, db *database.Database) {
	
	// Create route groups
	routeGroups := createRouteGroups(db)

	// Register middleware for each route group
	registerRouteGroups(router, routeGroups)

	// Register module routes
	authentication.RegisterModuleRoutes(router, db)
	configurations.RegisterModuleRoutes(router, db)
	references.RegisterModuleRoutes(router, db)
	company.RegisterModuleRoutes(router, db)
	employees.RegisterModuleRoutes(router, db)
	back_office.RegisterModuleRoutes(router, db)
	reports.RegisterModuleRoutes(router, db)
}

// routeGroup represents a route group with its prefix, module, and middlewares
type routeGroup struct {
	prefix      string
	module      entities.ModuleInfo
	middlewares []fiber.Handler
}

// createRouteGroups initializes and returns a slice of routeGroup
func createRouteGroups(db *database.Database) []routeGroup {
	// Initialize the middlewares
	flagMiddleware := middlewares.NewModuleFlagMiddleware(db)
	//sessionMiddleware := middlewares.NewSessionAuthMiddleware(db)
	//apiKeyMiddleware := middlewares.NewApiKeyMiddleware(db)
	//jwtMiddleware := middlewares.NewJwtMiddleware(db)

	return []routeGroup{
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
				// Add other middlewares specific to configurations if needed
			},
		},
		{
			prefix: "/api/configurations",
			module: entities.ModuleConfigurations,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleConfigurations.Code),
				// Add other middlewares specific to configurations if needed
			},
		},
		{
			prefix: "/company",
			module: entities.ModuleCompany,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleCompany.Code),
				// Add other middlewares specific to company if needed
			},
		},
		{
			prefix: "/api/company",
			module: entities.ModuleCompany,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleCompany.Code),
				// Add other middlewares specific to company if needed
			},
		},
		{
			prefix: "/employees",
			module: entities.ModuleEmployees,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleEmployees.Code),
				// Add other middlewares specific to employees if needed
			},
		},
		{
			prefix: "/api/employees",
			module: entities.ModuleEmployees,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleEmployees.Code),
				// Add other middlewares specific to employees if needed
			},
		},
		{
			prefix: "/references",
			module: entities.ModuleReferences,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleReferences.Code),
				// Add other middlewares specific to references if needed
			},
		},
		{
			prefix: "/api/references",
			module: entities.ModuleReferences,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(entities.ModuleReferences.Code),
				// Add other middlewares specific to references if needed
			},
		},
	}
}

// registerRouteGroups registers routes for each routeGroup
func registerRouteGroups(router *fiber.App, routeGroups []routeGroup) {
	for _, group := range routeGroups {
		router.Group(group.prefix, group.middlewares...)
		// Optionally register API version routes
		router.Group("/api"+group.prefix, group.middlewares...)
	}
}
