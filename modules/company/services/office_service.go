package services

import (
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
)

type OfficeService struct {
	repository *repositories.OfficeRepository
}

func NewOfficeService(db *database.Database) *OfficeService {
	return &OfficeService{
		repository: repositories.NewOfficeRepository(db),
	}
}