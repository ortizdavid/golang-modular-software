package services

import (
	"context"
	"fmt"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/repositories"
)

type ModuleFlagStatusService struct {
	repository *repositories.ModuleFlagRepository
}

func NewModuleFlagStatusService(db *database.Database) *ModuleFlagStatusService {
	return &ModuleFlagStatusService{
		repository: repositories.NewModuleFlagRepository(db),
	}
}

func (s *ModuleFlagStatusService) LoadModuleFlagStatus(ctx context.Context) (entities.ModuleFlagStatus, error) {
	// Fetch all module flags in a single query
	flagMap, err := s.GetAllModuleFlags(ctx)
	if err != nil {
		return entities.ModuleFlagStatus{}, err
	}
	// Construct and return the flag status
	flagStatus := entities.ModuleFlagStatus{
		Authentication: flagMap[entities.ModuleAuthentication.Code],
		Company:        flagMap[entities.ModuleCompany.Code],
		Employees:      flagMap[entities.ModuleEmployees.Code],
		References:     flagMap[entities.ModuleReferences.Code],
		Reports:        flagMap[entities.ModuleReports.Code],
		Configurations: flagMap[entities.ModuleConfigurations.Code],
	}
	return flagStatus, nil
}

func (s *ModuleFlagStatusService) GetAllModuleFlags(ctx context.Context) (map[string]string, error) {
	flagMap, err := s.repository.FindAllFlagsMap(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching all module flags: %w", err)
	}
	return flagMap, nil
}
