package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/repositories"
)

type EmployeeCompleteDataService struct {
	repository *repositories.EmployeeCompleteDataRepository
}

func NewEmployeeCompleteDataService(db *database.Database) *EmployeeCompleteDataService {
	return &EmployeeCompleteDataService{
		repository: repositories.NewEmployeeCompleteDataRepository(db),
	}
}

func (s *EmployeeCompleteDataService) GetByIdentificationNumber(ctx context.Context, identNumber string) (entities.EmployeeCompleteData, error) {
	employee, err := s.repository.GetByIdentificationNumber(ctx, identNumber)
	if err != nil {
		return entities.EmployeeCompleteData{}, apperrors.NotFoundError("employee not found")
	}
	if employee.EmployeeId == 0 {
		return entities.EmployeeCompleteData{}, apperrors.NotFoundError("employee not found")
	}
	return employee, nil
}

func (s *EmployeeCompleteDataService) GetByUniqueId(ctx context.Context, identNumber string) (entities.EmployeeCompleteData, error) {
	employee, err := s.repository.GetByUniqueId(ctx, identNumber)
	if err != nil {
		return entities.EmployeeCompleteData{}, apperrors.NotFoundError("employee not found")
	}
	if employee.EmployeeId == 0 {
		return entities.EmployeeCompleteData{}, apperrors.NotFoundError("employee not found")
	}
	return employee, nil
}
