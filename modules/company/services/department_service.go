package services

import (
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
)

type DepartmentService struct {
	repository *repositories.DepartmentRepository
}

func NewDepartmentService(db *database.Database) *DepartmentService {
	return &DepartmentService{
		repository: repositories.NewDepartmentRepository(db),
	}
}