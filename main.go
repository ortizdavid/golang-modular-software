package main

import (
	"log"

	"github.com/ortizdavid/golang-modular-software/application"
)

func main() {
	// create and initialize the application
	app, err := application.NewApplication()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// start the application
	app.Start()
}