package services

import (
	"context"
	"fmt"
	//"fmt"

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
	// Fetch the status of each module flag

	// List of all IDs needed
	ids := []int{
		entities.CoreEntityUsers.Id,
		entities.CoreEntityActiveUsers.Id,
		entities.CoreEntityInactiveUsers.Id,
		entities.CoreEntityOnlineUsers.Id,
		entities.CoreEntityOfflineUsers.Id,
		entities.CoreEntityRoles.Id,
		entities.CoreEntityPermissions.Id,
		entities.CoreEntityLoginActivity.Id,

		entities.CoreEntityBasicConfigurations.Id,
		entities.CoreEntityCompanyConfigurations.Id,
		entities.CoreEntityEmailConfigurations.Id,
		entities.CoreEntityModules.Id,
		entities.CoreEntityCoreEntities.Id,
		entities.CoreEntityModuleFlags.Id,
		entities.CoreEntityCoreEntityFlags.Id,

		entities.CoreEntityCountries.Id,
		entities.CoreEntityCurrencies.Id,
		entities.CoreEntityIdentificationTypes.Id,
		entities.CoreEntityContactTypes.Id,
		entities.CoreEntityMaritalStatuses.Id,
		entities.CoreEntityTaskStatuses.Id,
		entities.CoreEntityApprovalStatuses.Id,
		entities.CoreEntityDocumentStatuses.Id,
		entities.CoreEntityWorkflowStatuses.Id,
		entities.CoreEntityEvaluationStatuses.Id,
		entities.CoreEntityUserStatuses.Id,
		entities.CoreEntityEmploymentStatuses.Id,

		entities.CoreEntityCompanyInfo.Id,
		entities.CoreEntityBranches.Id,
		entities.CoreEntityOffices.Id,
		entities.CoreEntityDepartments.Id,
		entities.CoreEntityRooms.Id,
		entities.CoreEntityProjects.Id,
		entities.CoreEntityPolicies.Id,

		entities.CoreEntityEmployees.Id,
		entities.CoreEntityJobTitles.Id,

		entities.CoreEntityUserReports.Id,
		entities.CoreEntityConfigurationReports.Id,
		entities.CoreEntityCompanyReports.Id,
		entities.CoreEntityEmployeeReports.Id,
		entities.CoreEntityReferenceReports.Id,
	}

	// Fetch all entity flags in a single query
	flagMap, err := s.repository.FindByIdBatch(ctx, ids)
	if err != nil {
		return entities.CoreEntityFlagStatus{}, fmt.Errorf("error fetching core entity flags: %w", err)
	}

	// Construct and return the flag status
	// Initialize the flag status structure
	flagStatus := entities.CoreEntityFlagStatus{
		AuthenticationFlags: entities.AuthenticationFlags{
			Users:         flagMap[entities.CoreEntityUsers.Id],
			ActiveUsers:   flagMap[entities.CoreEntityActiveUsers.Id],
			InactiveUsers: flagMap[entities.CoreEntityInactiveUsers.Id],
			OnlineUsers:   flagMap[entities.CoreEntityOnlineUsers.Id],
			OfflineUsers:  flagMap[entities.CoreEntityOfflineUsers.Id],
			Roles:         flagMap[entities.CoreEntityRoles.Id],
			Permissions:   flagMap[entities.CoreEntityPermissions.Id],
			LoginActivity: flagMap[entities.CoreEntityLoginActivity.Id],
		},
		ConfigurationFlags: entities.ConfigurationFlags{
			BasicConfigurations:   flagMap[entities.CoreEntityBasicConfigurations.Id],
			CompanyConfigurations: flagMap[entities.CoreEntityCompanyConfigurations.Id],
			EmailConfigurations:   flagMap[entities.CoreEntityEmailConfigurations.Id],
			Modules:               flagMap[entities.CoreEntityModules.Id],
			CoreEntities:          flagMap[entities.CoreEntityCoreEntities.Id],
			ModuleFlags:           flagMap[entities.CoreEntityModuleFlags.Id],
			CoreEntityFlags:       flagMap[entities.CoreEntityCoreEntityFlags.Id],
		},
		ReferenceFlags: entities.ReferenceFlags{
			Countries:           flagMap[entities.CoreEntityCountries.Id],
			Currencies:          flagMap[entities.CoreEntityCurrencies.Id],
			IdentificationTypes: flagMap[entities.CoreEntityIdentificationTypes.Id],
			ContactTypes:        flagMap[entities.CoreEntityContactTypes.Id],
			MaritalStatuses:     flagMap[entities.CoreEntityMaritalStatuses.Id],
			TaskStatuses:        flagMap[entities.CoreEntityTaskStatuses.Id],
			ApprovalStatuses:    flagMap[entities.CoreEntityApprovalStatuses.Id],
			DocumentStatuses:    flagMap[entities.CoreEntityDocumentStatuses.Id],
			WorkflowStatuses:    flagMap[entities.CoreEntityWorkflowStatuses.Id],
			EvaluationStatuses:  flagMap[entities.CoreEntityEvaluationStatuses.Id],
			UserStatuses:        flagMap[entities.CoreEntityUserStatuses.Id],
			EmploymentStatuses:  flagMap[entities.CoreEntityEmploymentStatuses.Id],
		},
		CompanyFlags: entities.CompanyFlags{
			CompanyInfo: flagMap[entities.CoreEntityCompanyInfo.Id],
			Branches:    flagMap[entities.CoreEntityBranches.Id],
			Offices:     flagMap[entities.CoreEntityOffices.Id],
			Departments: flagMap[entities.CoreEntityDepartments.Id],
			Rooms:       flagMap[entities.CoreEntityRooms.Id],
			Projects:    flagMap[entities.CoreEntityProjects.Id],
			Policies:    flagMap[entities.CoreEntityPolicies.Id],
		},
		EmployeeFlags: entities.EmployeeFlags{
			Employees: flagMap[entities.CoreEntityEmployees.Id],
			JobTitles: flagMap[entities.CoreEntityJobTitles.Id],
		},
		ReportFlags: entities.ReportFlags{
			UserReports:          flagMap[entities.CoreEntityUserReports.Id],
			ConfigurationReports: flagMap[entities.CoreEntityConfigurationReports.Id],
			CompanyReports:       flagMap[entities.CoreEntityCompanyReports.Id],
			EmployeeReports:      flagMap[entities.CoreEntityEmployeeReports.Id],
			ReferenceReports:     flagMap[entities.CoreEntityReferenceReports.Id],
		},
	}

	return flagStatus, nil
}
