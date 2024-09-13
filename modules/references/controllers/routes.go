package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterControllerRoutes(router *fiber.App, db *database.Database) {
	NewRootController(db).Routes(router, db)
	NewCountryController(db).Routes(router, db)
	NewCurrencyController(db).Routes(router, db)
	NewTaskStatusController(db).Routes(router, db)
	NewApprovalStatusController(db).Routes(router, db)
	NewWorkflowStatusController(db).Routes(router, db)
	NewEvaluationStatusController(db).Routes(router, db)
	NewUserStatusController(db).Routes(router, db)
	NewEmploymentStatusController(db).Routes(router, db)
	NewDocumentStatusController(db).Routes(router, db)
	NewMaritalStatusController(db).Routes(router, db)
	NewContactTypeController(db).Routes(router, db)
	NewIdentificationTypeController(db).Routes(router, db)
}
