package repositories

import "github.com/ortizdavid/golang-modular-software/database"

type ProjectRepository struct {
	db *database.Database
}

func NewProjectRepository(db *database.Database) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}