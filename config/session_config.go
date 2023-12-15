package config

import (
	"os"
	"time"
	"github.com/gofiber/storage/postgres"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func GetSessionStore() *session.Store {
	storage := postgres.New(postgres.Config{
		ConnectionURI: ConnectionString(),
		Reset:         false,
		GCInterval:  time.Duration(SessionExpiration()) * time.Minute,
	})
	store := session.New(session.Config{
        Storage: storage, 
    })
	return store
}

func SessionExpiration() int {
	LoadDotEnv()
	return conversion.StringToInt(os.Getenv("APP_SESSION_EXPIRATION"))
}