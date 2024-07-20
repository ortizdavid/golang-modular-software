package config

import (
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
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", 
		GetEnv("DB_USER"), 
		GetEnv("DB_PASSWORD"), 
		GetEnv("DB_HOST"),
		GetEnv("DB_PORT"), 
		GetEnv("DB_NAME"),
	)
}

