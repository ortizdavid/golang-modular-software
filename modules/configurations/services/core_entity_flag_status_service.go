package services

import (
	"context"
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
	/*// Fetch the status of each module flag
	authenticationFlag, err := s.repository.FindById(ctx, entities.CoreEntityAuthentication.Id)
	if err != nil {
		return entities.CoreEntityFlagStatus{}, fmt.Errorf("error fetching authentication flag: %w", err)
	}
	companyFlag, err := s.repository.FindById(ctx, entities.CoreEntityCompany.Id)
	if err != nil {
		return entities.CoreEntityFlagStatus{}, fmt.Errorf("error fetching company flag: %w", err)
	}
	employeesFlag, err := s.repository.FindById(ctx, entities.CoreEntityEmployees.Id)
	if err != nil {
		return entities.CoreEntityFlagStatus{}, fmt.Errorf("error fetching employees flag: %w", err)
	}
	referencesFlag, err := s.repository.FindById(ctx, entities.CoreEntityReferences.Id)
	if err != nil {
		return entities.CoreEntityFlagStatus{}, fmt.Errorf("error fetching references flag: %w", err)
	}
	reportsFlag, err := s.repository.FindById(ctx, entities.CoreEntityReports.Id)
	if err != nil {
		return entities.CoreEntityFlagStatus{}, fmt.Errorf("error fetching reports flag: %w", err)
	}
	configurationsFlag, err := s.repository.FindById(ctx, entities.CoreEntityConfigurations.Id)
	if err != nil {
		return entities.CoreEntityFlagStatus{}, fmt.Errorf("error fetching configurations flag: %w", err)
	}*/
	// Construct and return the flag status
	flagStatus := entities.CoreEntityFlagStatus{
		AuthenticationFlags: entities.AuthenticationFlags{
			Users:         "",
			ActiveUsers:   "",
			InactiveUsers: "",
			OnlineUsers:   "",
			OfflineUsers:  "",
			Roles:         "",
			Permissions:   "",
			LoginActivity: "",
		},
		ConfigurationFlags:  entities.ConfigurationFlags{
			BasicConfigurations:   "",
			CompanyConfigurations: "",
			EmailConfigurations:   "",
			Modules:               "",
			CoreEntities:          "",
			ModuleFlags:           "",
			CoreEntityFlags:       "",
		},
		ReferenceFlags:      entities.ReferenceFlags{
			Countries:           "",
			Currencies:          "",
			IdentificationTypes: "",
			ContactTypes:        "",
			MaritalStatuses:     "",
			TaskStatuses:        "",
			ApprovalStatuses:    "",
			DocumentStatuses:    "",
			WorkflowStatuses:    "",
			EvaluationStatuses:  "",
			UserStatuses:        "",
			EmploymentStatuses:  "",
		},
		CompanyFlags:        entities.CompanyFlags{
			CompanyInfo: "",
			Branches:    "",
			Offices:     "",
			Departments: "",
			Rooms:       "",
			Projects:    "",
			Policies:    "",
		},
		EmployeeFlags:       entities.EmployeeFlags{
			Employees: "",
			JobTitles: "",
		},
		ReportFlags:         entities.ReportFlags{
			UserReports:          "",
			ConfigurationReports: "",
			CompanyReports:       "",
			EmployeeReports:      "",
			ReferenceReports:     "",
		},
	}
	return flagStatus, nil
}
