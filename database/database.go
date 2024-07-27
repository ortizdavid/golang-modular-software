package database

import "gorm.io/gorm"

type Database struct {
	*gorm.DB
}

func NewDatabase(db *gorm.DB) *Database {
	return &Database{
		DB: db,
	}
}