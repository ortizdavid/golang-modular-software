package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/database"
	authEntities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	authServices "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/repositories"
)

type EmployeeAccountService struct {
	repository *repositories.EmployeeAccountRepository
	roleService *authServices.RoleService
}

func NewEmployeeAccountService(db *database.Database) *EmployeeAccountService {
	return &EmployeeAccountService{
		repository:  repositories.NewEmployeeAccountRepository(db),
		roleService: authServices.NewRoleService(db),
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

func (s *EmployeeAccountService) GetEmployeAllowedRoles(ctx context.Context) ([]authEntities.Role, error) {
	roles, err := s.roleService.GetAllEnaledRolesNotIn(ctx, s.GetAllowedRoles())
	if err != nil {
		return nil, apperrors.NewNotFoundError("roles not found")
	}
	return roles, nil
}

func (s *EmployeeAccountService) GetAllowedRoles() []string {
	return []string{
		authEntities.RoleEmployee.Code, 
		authEntities.RoleManager.Code, 
	}
}