package services

import (
	"context"
	"fmt"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/repositories"
)

type ModuleFlagStatusService struct {
	repository        *repositories.ModuleFlagRepository
}

func NewModuleFlagStatusService(db *database.Database) *ModuleFlagStatusService {
	return &ModuleFlagStatusService{
		repository:       repositories.NewModuleFlagRepository(db),
	}
}

func (s *ModuleFlagStatusService) LoadModuleFlagStatus(ctx context.Context) (entities.ModuleFlagStatus, error) {
    // Fetch the status of each module flag
    authenticationFlag, err := s.repository.FindById(ctx, entities.ModuleAuthentication)
    if err != nil {
        return entities.ModuleFlagStatus{}, fmt.Errorf("error fetching authentication flag: %w", err)
    }
    companyFlag, err := s.repository.FindById(ctx, entities.ModuleCompany)
    if err != nil {
        return entities.ModuleFlagStatus{}, fmt.Errorf("error fetching company flag: %w", err)
    }
    employeesFlag, err := s.repository.FindById(ctx, entities.ModuleEmployees)
    if err != nil {
        return entities.ModuleFlagStatus{}, fmt.Errorf("error fetching employees flag: %w", err)
    }
    referencesFlag, err := s.repository.FindById(ctx, entities.ModuleReferences)
    if err != nil {
        return entities.ModuleFlagStatus{}, fmt.Errorf("error fetching references flag: %w", err)
    }
    reportsFlag, err := s.repository.FindById(ctx, entities.ModuleReports)
    if err != nil {
        return entities.ModuleFlagStatus{}, fmt.Errorf("error fetching reports flag: %w", err)
    }
    configurationsFlag, err := s.repository.FindById(ctx, entities.ModuleConfigurations)
    if err != nil {
        return entities.ModuleFlagStatus{}, fmt.Errorf("error fetching configurations flag: %w", err)
    }
    // Construct and return the flag status
	flagStatus := entities.ModuleFlagStatus{
        Authentication: authenticationFlag.Status,
        Company:        companyFlag.Status,
        Employees:      employeesFlag.Status,
        References:     referencesFlag.Status,
        Reports:        reportsFlag.Status,
        Configurations: configurationsFlag.Status,
    }
    return flagStatus, nil
}
