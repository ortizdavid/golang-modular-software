package repositories

import "github.com/ortizdavid/golang-modular-software/database"

type BranchRepository struct {
	db *database.Database
}

func NewBranchRepository(db *database.Database) *BranchRepository {
	return &BranchRepository{
		db: db,
	}
}