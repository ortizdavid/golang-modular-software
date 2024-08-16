package services

import (
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
)

type PolicyService struct {
	repository *repositories.PolicyRepository
}

func NewPolicyService(db *database.Database) *PolicyService {
	return &PolicyService{
		repository: repositories.NewPolicyRepository(db),
	}
}