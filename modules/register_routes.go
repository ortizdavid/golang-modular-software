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
	configEntities "github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	authEntities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
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
	module      configEntities.ModuleInfo
	middlewares []fiber.Handler
}

// createRouteGroups initializes and returns a slice of routeGroup
func createRouteGroups(db *database.Database) []routeGroup {
	// Initialize the middlewares
	flagMiddleware := middlewares.NewModuleFlagMiddleware(db)
	//authorization := middlewares.NewAuthorizationMiddleware(db)
	sessionMiddleware := middlewares.NewSessionAuthMiddleware(db)
	apiKeyMiddleware := middlewares.NewApiKeyMiddleware(db)
	//jwtMiddleware := middlewares.NewJwtMiddleware(db)

	return []routeGroup{
		{
			prefix: "/account",
			middlewares: []fiber.Handler{
				sessionMiddleware.CheckLoggedUser,

			},
		},
		{
			prefix: "/user-management",
			module: configEntities.ModuleAuthentication,
			middlewares: []fiber.Handler{
				sessionMiddleware.CheckLoggedUser,
				flagMiddleware.CheckModule(configEntities.ModuleAuthentication.Code),

			},
		},
		{
			prefix: "/api/user-management",
			module: configEntities.ModuleAuthentication,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(configEntities.ModuleAuthentication.Code),

			},
		},
		{
			prefix: "/configurations",
			module: configEntities.ModuleConfigurations,
			middlewares: []fiber.Handler{
				sessionMiddleware.CheckLoggedUser,
				flagMiddleware.CheckModule(configEntities.ModuleConfigurations.Code),
				
			},
		},
		{
			prefix: "/api/configurations",
			module: configEntities.ModuleConfigurations,
			middlewares: []fiber.Handler{
				apiKeyMiddleware.AllowRoles(authEntities.RoleSuperAdmin),
				flagMiddleware.CheckModule(configEntities.ModuleConfigurations.Code),

			},
		},
		{
			prefix: "/company",
			module: configEntities.ModuleCompany,
			middlewares: []fiber.Handler{
				sessionMiddleware.CheckLoggedUser,
				flagMiddleware.CheckModule(configEntities.ModuleCompany.Code),
				
			},
		},
		{
			prefix: "/api/company",
			module: configEntities.ModuleCompany,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(configEntities.ModuleCompany.Code),
				
			},
		},
		{
			prefix: "/employees",
			module: configEntities.ModuleEmployees,
			middlewares: []fiber.Handler{
				sessionMiddleware.CheckLoggedUser,
				flagMiddleware.CheckModule(configEntities.ModuleEmployees.Code),
				
			},
		},
		{
			prefix: "/api/employees",
			module: configEntities.ModuleEmployees,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(configEntities.ModuleEmployees.Code),
				
			},
		},
		{
			prefix: "/references",
			module: configEntities.ModuleReferences,
			middlewares: []fiber.Handler{
				sessionMiddleware.CheckLoggedUser,
				flagMiddleware.CheckModule(configEntities.ModuleReferences.Code),
				
			},
		},
		{
			prefix: "/api/references",
			module: configEntities.ModuleReferences,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(configEntities.ModuleReferences.Code),
				
			},
		},
		{
			prefix: "/reports",
			module: configEntities.ModuleReferences,
			middlewares: []fiber.Handler{
				sessionMiddleware.CheckLoggedUser,
				flagMiddleware.CheckModule(configEntities.ModuleReferences.Code),
				
			},
		},
		{
			prefix: "/api/reports",
			module: configEntities.ModuleReferences,
			middlewares: []fiber.Handler{
				flagMiddleware.CheckModule(configEntities.ModuleReferences.Code),
				
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
