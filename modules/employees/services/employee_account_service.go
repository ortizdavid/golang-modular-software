package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/repositories"
)

type EmployeeAccountService struct {
	repository *repositories.EmployeeAccountRepository
}

func NewEmployeeAccountService(db *database.Database) *EmployeeAccountService {
	return &EmployeeAccountService{
		repository: repositories.NewEmployeeAccountRepository(db),
	}
}

func (s *EmployeeAccountService) GetEmployeeAccountById(ctx context.Context, employeeId int64) (entities.EmployeeAccountData, error) {
	employee, err := s.repository.GetDataByEmployeeId(ctx, employeeId)
	if err != nil {
		return entities.EmployeeAccountData{}, apperrors.NewNotFoundError("employee account not found")
	}
	return employee, nil
}

func (s *EmployeeAccountService) GetEmployeeAccountByIdentificationNumber(ctx context.Context, identNumber string) (entities.EmployeeAccountData, error) {
	employee, err := s.repository.GetDataByIdentificationNumber(ctx, identNumber)
	if err != nil {
		return entities.EmployeeAccountData{}, apperrors.NewNotFoundError("employee account not found")
	}
	return employee, nil
}

