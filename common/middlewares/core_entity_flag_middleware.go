package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type CoreEntityFlagFlagMiddleware struct {
	flagService *services.CoreEntityFlagStatusService
}

func NewCoreEntityFlagFlagMiddleware(db *database.Database) *CoreEntityFlagFlagMiddleware {
	return &CoreEntityFlagFlagMiddleware{
		flagService: services.NewCoreEntityFlagStatusService(db),
	}
}

// CheckCoreEntityFlag dynamically checks if a module is enabled
func (mid *CoreEntityFlagFlagMiddleware) CheckCoreEntityFlag(moduleCode string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		/*flagStatus, err := mid.flagService.LoadCoreEntityFlagStatus(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error loading module status")
		}*/
		var isEnabled bool

		// Check the status of the module based on its name
		switch moduleCode {
		/*case "authentication":
			isEnabled = flagStatus.Authentication == "Enabled"
		case "configurations":
			isEnabled = flagStatus.Configurations == "Enabled"
		case "company":
			isEnabled = flagStatus.Company == "Enabled"
		case "employees":
			isEnabled = flagStatus.Employees == "Enabled"
		case "reports":
			isEnabled = flagStatus.Reports == "Enabled"
		case "references":
			isEnabled = flagStatus.References == "Enabled"*/
		default:
			isEnabled = false
		}

		if !isEnabled {
			return c.Status(fiber.StatusForbidden).Render("_errors/error", fiber.Map{
				"Title":   "CoreEntityFlag Error",
				"Message": "CoreEntityFlag '"+moduleCode+"' is Disabled",
			})
		}
		return c.Next()
	}
}