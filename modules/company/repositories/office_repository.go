package repositories

import "github.com/ortizdavid/golang-modular-software/database"

type OfficeRepository struct {
	db *database.Database
}

func NewOfficeRepository(db *database.Database) *OfficeRepository {
	return &OfficeRepository{
		db: db,
	}
}