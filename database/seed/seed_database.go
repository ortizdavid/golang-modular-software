package database

import (
	"github.com/ortizdavid/golang-modular-software/database"
)

type SeedingStatus struct {
    ID        uint `gorm:"primarykey"`
    Seeded    bool
}

func SeedDatabase(db *database.Database) error {
	if err := InitDatabaseScripts(db); err != nil {
		return err
	}
	if err := CreateAdminUsers(db); err != nil {
		return err
	}
	return nil
}