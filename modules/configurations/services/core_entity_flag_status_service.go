package services

import (
	"context"
	"fmt"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/repositories"
)

type CoreEntityFlagStatusService struct {
	repository *repositories.CoreEntityFlagRepository
}

func NewCoreEntityFlagStatusService(db *database.Database) *CoreEntityFlagStatusService {
	return &CoreEntityFlagStatusService{
		repository: repositories.NewCoreEntityFlagRepository(db),
	}
}

func (s *CoreEntityFlagStatusService) LoadCoreEntityFlagStatus(ctx context.Context) (entities.CoreEntityFlagStatus, error) {
	// Fetch the status of each core entity flag
	flagMap, err := s.GetAll(ctx)
	if err != nil {
		return entities.CoreEntityFlagStatus{}, err
	}
	// Initialize the flag status structure
	flagStatus := entities.CoreEntityFlagStatus{
		AuthenticationFlags: entities.AuthenticationFlags{
			Users:         flagMap[entities.CoreEntityUsers.Code],
			ActiveUsers:   flagMap[entities.CoreEntityActiveUsers.Code],
			InactiveUsers: flagMap[entities.CoreEntityInactiveUsers.Code],
			OnlineUsers:   flagMap[entities.CoreEntityOnlineUsers.Code],
			OfflineUsers:  flagMap[entities.CoreEntityOfflineUsers.Code],
			Roles:         flagMap[entities.CoreEntityRoles.Code],
			Permissions:   flagMap[entities.CoreEntityPermissions.Code],
			LoginActivity: flagMap[entities.CoreEntityLoginActivity.Code],
		},
		ConfigurationFlags: entities.ConfigurationFlags{
			BasicConfigurations:   flagMap[entities.CoreEntityBasicConfigurations.Code],
			CompanyConfigurations: flagMap[entities.CoreEntityCompanyConfigurations.Code],
			EmailConfigurations:   flagMap[entities.CoreEntityEmailConfigurations.Code],
			Modules:               flagMap[entities.CoreEntityModules.Code],
			CoreEntities:          flagMap[entities.CoreEntityCoreEntities.Code],
			ModuleFlags:           flagMap[entities.CoreEntityModuleFlags.Code],
			CoreEntityFlags:       flagMap[entities.CoreEntityCoreEntityFlags.Code],
		},
		ReferenceFlags: entities.ReferenceFlags{
			Countries:           flagMap[entities.CoreEntityCountries.Code],
			Currencies:          flagMap[entities.CoreEntityCurrencies.Code],
			IdentificationTypes: flagMap[entities.CoreEntityIdentificationTypes.Code],
			ContactTypes:        flagMap[entities.CoreEntityContactTypes.Code],
			MaritalStatuses:     flagMap[entities.CoreEntityMaritalStatuses.Code],
			TaskStatuses:        flagMap[entities.CoreEntityTaskStatuses.Code],
			ApprovalStatuses:    flagMap[entities.CoreEntityApprovalStatuses.Code],
			DocumentStatuses:    flagMap[entities.CoreEntityDocumentStatuses.Code],
			WorkflowStatuses:    flagMap[entities.CoreEntityWorkflowStatuses.Code],
			EvaluationStatuses:  flagMap[entities.CoreEntityEvaluationStatuses.Code],
			UserStatuses:        flagMap[entities.CoreEntityUserStatuses.Code],
			EmploymentStatuses:  flagMap[entities.CoreEntityEmploymentStatuses.Code],
		},
		CompanyFlags: entities.CompanyFlags{
			CompanyInfo: flagMap[entities.CoreEntityCompanyInfo.Code],
			Branches:    flagMap[entities.CoreEntityBranches.Code],
			Offices:     flagMap[entities.CoreEntityOffices.Code],
			Departments: flagMap[entities.CoreEntityDepartments.Code],
			Rooms:       flagMap[entities.CoreEntityRooms.Code],
			Projects:    flagMap[entities.CoreEntityProjects.Code],
			Policies:    flagMap[entities.CoreEntityPolicies.Code],
		},
		EmployeeFlags: entities.EmployeeFlags{
			Employees:     flagMap[entities.CoreEntityEmployees.Code],
			JobTitles:     flagMap[entities.CoreEntityJobTitles.Code],
			DocumentTypes: flagMap[entities.CoreEntityDocumentTypes.Code],
		},
		ReportFlags: entities.ReportFlags{
			UserReports:          flagMap[entities.CoreEntityUserReports.Code],
			ConfigurationReports: flagMap[entities.CoreEntityConfigurationReports.Code],
			CompanyReports:       flagMap[entities.CoreEntityCompanyReports.Code],
			EmployeeReports:      flagMap[entities.CoreEntityEmployeeReports.Code],
			ReferenceReports:     flagMap[entities.CoreEntityReferenceReports.Code],
		},
	}
	return flagStatus, nil
}


func (s *CoreEntityFlagStatusService) GetAll(ctx context.Context) (map[string]string, error) {
	flagMap, err := s.repository.FindAllFlagsMap(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching all core entity flags: %w", err)
	}
	return flagMap, nil
}
