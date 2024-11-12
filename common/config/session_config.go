package config

import (
	"time"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres"
	"github.com/ortizdavid/go-nopain/conversion"
)

func GetSessionStore() *session.Store {
	storage := postgres.New(postgres.Config{
		ConnectionURI: GetEnv("DATABASE_MAIN_URL"),
		Reset:         false,
		GCInterval:  time.Duration(SessionExpiration()) * time.Minute,
	})
	store := session.New(session.Config{
        Storage: storage, 
    })
	return store
}

func SessionExpiration() int {
	return conversion.StringToInt(GetEnv("APP_SESSION_EXPIRATION"))
}
