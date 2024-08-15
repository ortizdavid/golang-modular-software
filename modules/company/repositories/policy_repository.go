package repositories

import "github.com/ortizdavid/golang-modular-software/database"

type PolicyRepository struct {
	db *database.Database
}

func NewPolicyRepository(db *database.Database) *PolicyRepository {
	return &PolicyRepository{
		db: db,
	}
}