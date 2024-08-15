package repositories

import "github.com/ortizdavid/golang-modular-software/database"

type DepartmantRepository struct {
	db *database.Database
}

func NewDepartmantRepository(db *database.Database) *DepartmantRepository {
	return &DepartmantRepository{
		db: db,
	}
}