package main

import (
	"log"
	"github.com/ortizdavid/golang-modular-software/application"
)

func main() {
	// create and initialize the application
	app, err := application.NewApplication()
	if err != nil {
		log.Fatal(err.Error())
	}

	// start the application
	err = app.Start()
	if err != nil {
		log.Fatal(err.Error())
	}
}