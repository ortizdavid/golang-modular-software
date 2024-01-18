package main

import (
	"fmt"

	"github.com/ortizdavid/golang-modular-software/config"
)

func main() {
	fmt.Println(config.GetEnv("ENV"))
	fmt.Println(config.GetEnv("APP_HOST"))
	fmt.Println(config.GetEnv("APP_PORT"))
	fmt.Println(config.GetEnv("REQUESTS_PER_SECONDS"))
	fmt.Println(config.GetEnv("DB_HOST"))
	fmt.Println(config.GetEnv("DB_PASSWORD"))
}
