package application

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/config"
	"github.com/ortizdavid/golang-modular-software/middlewares"
	"github.com/ortizdavid/golang-modular-software/modules"
)

type Application struct {
	App *fiber.App
}

// initialize an application with static files, middlewares and controllers
func NewApplication() *Application {
	app := fiber.New(fiber.Config{
		Views: config.GetTemplateEngine(),		
	})
	// configure location of css, js, .jpg, .pdf and other files
	config.ConfigStaticFiles(app)
	// initialize all the middlewares needed
	middlewares.InitializeMiddlewares(app)
	// initialize all controllers containing routes of application
	modules.RegisterControllers(app)
	
	return &Application{
		App: app,
	}
}

// Start application in port defined at .ENV file
func (a *Application) Start()  {
	err := a.App.Listen(config.ListenAddr())
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}