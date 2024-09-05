package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type ModuleFlagMiddleware struct {
	flagService *services.ModuleFlagStatusService
}

func NewModuleFlagMiddleware(db *database.Database) *ModuleFlagMiddleware {
	return &ModuleFlagMiddleware{
		flagService: services.NewModuleFlagStatusService(db),
	}
}

// CheckModule dynamically checks if a module is enabled
func (mid *ModuleFlagMiddleware) CheckModule(moduleCode string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		flagStatus, err := mid.flagService.GetAllModuleFlags(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving module status")
		}
		if flagStatus[moduleCode] == "Disabled" {
			return c.Status(fiber.StatusForbidden).Render("_errors/error", fiber.Map{
				"Title":   "Module Error",
				"Message": "Module '" + moduleCode + "' is Disabled",
			})
		}
		return c.Next()
	}
}
