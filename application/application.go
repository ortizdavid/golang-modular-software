package application

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	seed "github.com/ortizdavid/golang-modular-software/database/seed"
	"github.com/ortizdavid/golang-modular-software/modules"
)

type Application struct {
	App *fiber.App
}

// initialize an application with static files, middlewares and controllers
func NewApplication() (*Application, error) {
	app := fiber.New(fiber.Config{
		Views: config.GetTemplateEngine(), // Template engines
	})

	// Configure location of css, js, .jpg, .pdf and other files
	config.ConfigStaticFiles(app)

	// Connect to Database
	dbConn, err := database.NewDBConnectionFromEnv("DATABASE_MAIN_URL") // Main database
	if err != nil {
		return nil, err
	}
	db := database.NewDatabase(dbConn.DB)

	// Seed Database
	if config.GetEnv("DATABASE_SEEDING_STATUS") == "false" {
		if err := seed.SeedDatabase(db); err != nil {
			return nil, err
		}
	}

	// Initialize all the middlewares needed
	middlewares.InitializeMiddlewares(app, db)
	
	// Initialize all controllers containing routes of application
	modules.RegisterRoutes(app, db)
	
	return &Application{
		App: app,
	}, nil
}

// Start application in port defined at .ENV file
func (a *Application) Start() error {
	err := a.App.Listen(config.ListenAddr())
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
		return err
	}
	return nil
}
