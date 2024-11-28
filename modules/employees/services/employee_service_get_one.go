package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

func (s *EmployeeService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeeData, error) {
	employee, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.EmployeeData{}, apperrors.NotFoundError("employee not found")
	}
	if employee.EmployeeId == 0 {
		return entities.EmployeeData{}, apperrors.NotFoundError("employee not found")
	}
	return employee, nil
}

func (s *EmployeeService) GetByIdentificationNumber(ctx context.Context, identNumber string) (entities.EmployeeData, error) {
	employee, err := s.repository.GetDataByIdentificationNumber(ctx, identNumber)
	if err != nil {
		return entities.EmployeeData{}, apperrors.NotFoundError("employee not found")
	}
	if employee.EmployeeId == 0 {
		return entities.EmployeeData{}, apperrors.NotFoundError("employee not found")
	}
	return employee, nil
}
