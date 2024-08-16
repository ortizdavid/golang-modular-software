package repositories

import "github.com/ortizdavid/golang-modular-software/database"

type DepartmentRepository struct {
	db *database.Database
}

func NewDepartmentRepository(db *database.Database) *DepartmentRepository {
	return &DepartmentRepository{
		db: db,
	}
}