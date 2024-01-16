package config

import (
	"os"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := ConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func DisconnectDB(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic("Failed to disconnect DB")
	}
	conn.Close()
}

func ConnectionString() string {
	LoadDotEnv()
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), 
		os.Getenv("DB_NAME"),
	)
}