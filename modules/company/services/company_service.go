package services

import (
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
)

type CompanyService struct {
	repository *repositories.CompanyRepository
}

func NewCompanyService(db *database.Database) *CompanyService {
	return &CompanyService{
		repository: repositories.NewCompanyRepository(db),
	}
}