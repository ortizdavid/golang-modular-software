package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/database"
)

func RegisterApiRoutes(router *fiber.App, db *database.Database) {
	NewCountryApi(db).Routes(router, db)
	NewCurrencyApi(db).Routes(router, db)
	NewIdentificationTypeApi(db).Routes(router, db)
	NewContactTypeApi(db).Routes(router, db)
	NewMaritalStatusApi(db).Routes(router, db)
}