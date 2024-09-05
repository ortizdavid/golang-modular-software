package configurations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/api"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/controllers"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

func RegisterModuleRoutes(router *fiber.App, db *database.Database) {
	flagMiddleware := middlewares.NewCoreEntityFlagFlagMiddleware(db)
	// Core Entity Routes and Codes
	configurationRoutes := []struct{
		Path string
		EntityCode string
	}{
		{"/basic-configurations", entities.CoreEntityBasicConfigurations.Code},
		{"/company-configurations", entities.CoreEntityCompanyConfigurations.Code},
		{"/email-configurations", entities.CoreEntityEmailConfigurations.Code},
		{"/modules", entities.CoreEntityModules.Code},
		{"/core-entities", entities.CoreEntityCoreEntities.Code},
		{"/module-flags", entities.CoreEntityModuleFlags.Code},
		{"/core-entity-flags", entities.CoreEntityCoreEntityFlags.Code},
	}
	// Register regular and API routes for each entity and apply Flags
	for _, route := range configurationRoutes {
		router.Group("/configurations"+route.Path, flagMiddleware.CheckCoreEntityFlag(route.EntityCode))
		router.Group("/api/configurations"+route.Path, flagMiddleware.CheckCoreEntityFlag(route.EntityCode))
	}
	// Register additional controller and API routes
	controllers.RegisterRoutes(router, db)
	api.RegisterRoutes(router, db)
}