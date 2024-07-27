package application

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules"
)

type Application struct {
	App *fiber.App
}

// initialize an application with static files, middlewares and controllers
func NewApplication() (*Application, error) {
	app := fiber.New(fiber.Config{
		Views: config.GetTemplateEngine(),		
	})
	// Main database
	dbConn, err := database.NewDBConnectionFromEnv("DATABASE_MAIN_URL")
	if err != nil {
		return nil, err
	}
	database := database.NewDatabase(dbConn.DB)
	// configure location of css, js, .jpg, .pdf and other files
	config.ConfigStaticFiles(app)
	// initialize all the middlewares needed
	middlewares.InitializeMiddlewares(app, database)
	// initialize all controllers containing routes of application
	modules.RegisterControllers(app, database)
	
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
