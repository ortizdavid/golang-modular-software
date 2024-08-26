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
func (mid *ModuleFlagMiddleware) CheckModule(moduleName string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		flagStatus, err := mid.flagService.LoadModuleFlagStatus(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error loading module status")
		}
		var isEnabled bool

		// Check the status of the module based on its name
		switch moduleName {
		case "Authentication":
			isEnabled = flagStatus.Authentication == "Enabled"
		case "Configurations":
			isEnabled = flagStatus.Configurations == "Enabled"
		case "Company":
			isEnabled = flagStatus.Company == "Enabled"
		case "Employees":
			isEnabled = flagStatus.Employees == "Enabled"
		case "Reports":
			isEnabled = flagStatus.Reports == "Enabled"
		case "References":
			isEnabled = flagStatus.References == "Enabled"
		default:
			isEnabled = false
		}

		if !isEnabled {
			return c.Status(fiber.StatusForbidden).Render("_errors/error", fiber.Map{
				"Title":   "Module Error",
				"Message": "Module Disabled",
			})
		}
		return c.Next()
	}
}
