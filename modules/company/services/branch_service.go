package services

import (
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
)

type BranchService struct {
	repository *repositories.BranchRepository
}

func NewBranchService(db *database.Database) *BranchService {
	return &BranchService{
		repository: repositories.NewBranchRepository(db),
	}
}