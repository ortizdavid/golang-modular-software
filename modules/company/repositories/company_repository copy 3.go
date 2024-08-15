package repositories

import "github.com/ortizdavid/golang-modular-software/database"

type CompanyRepository struct {
	db *database.Database
}

func NewCompanyRepository(db *database.Database) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}