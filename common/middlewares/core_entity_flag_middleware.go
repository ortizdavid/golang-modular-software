package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type CoreEntityFlagMiddleware struct {
	flagService *services.CoreEntityFlagStatusService
}

func NewCoreEntityFlagFlagMiddleware(db *database.Database) *CoreEntityFlagMiddleware {
	return &CoreEntityFlagMiddleware{
		flagService: services.NewCoreEntityFlagStatusService(db),
	}
}

// CheckCoreEntityFlag dynamically checks if a core entity is enabled
func (mid *CoreEntityFlagMiddleware) CheckCoreEntityFlag(coreEntityCode string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		flagStatus, err := mid.flagService.GetAllEntityCoreFlags(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving module status")
		}
		if flagStatus[coreEntityCode]== "Disabled" {
			return c.Status(fiber.StatusForbidden).Render("_errors/error", fiber.Map{
				"Title":   "Core Entity Flag Error",
				"Message": "Core Entity '"+coreEntityCode+"' is Disabled",
			})
		}
		return c.Next()
	}
}
