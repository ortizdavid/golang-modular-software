package main

import (
	"github.com/ortizdavid/golang-modular-software/application"
)

func main() {

	// create and initialize the application
	app := application.NewApplication()

	// start the application
	app.Start()
	app.Shutdown()
}