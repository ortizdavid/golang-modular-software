package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
}

func GetEnv(key string) string {
	LoadDotEnv()
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not found", key)
	}
	return value
}