package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	"github.com/ortizdavid/golang-modular-software/modules/references/repositories"
)

type StatisticsService struct {
	countryRepository            *repositories.CountryRepository
	currencyRepository           *repositories.CurrencyRepository
	identificationTypeRepository *repositories.IdentificationTypeRepository
	contactTypeRepository        *repositories.ContactTypeRepository
	maritalStatusRepository      *repositories.MaritalStatusRepository
	taskStatusRepository         *repositories.TaskStatusRepository
	approvalStatusRepository     *repositories.ApprovalStatusRepository
	documentStatusRepository     *repositories.DocumentStatusRepository
	workflowStatusRepository     *repositories.WorkflowStatusRepository
	evaluationStatusRepository     *repositories.EvaluationStatusRepository
	userStatusRepository         *repositories.UserStatusRepository
	employmentStatusRepository         *repositories.EmploymentStatusRepository
}

func NewStatisticsService(db *database.Database) *StatisticsService {
	return &StatisticsService{
		countryRepository:            repositories.NewCountryRepository(db),
		currencyRepository:           repositories.NewCurrencyRepository(db),
		identificationTypeRepository: repositories.NewIdentificationTypeRepository(db),
		contactTypeRepository:        repositories.NewContactTypeRepository(db),
		maritalStatusRepository:      repositories.NewMaritalStatusRepository(db),
		taskStatusRepository:         repositories.NewTaskStatusRepository(db),
		approvalStatusRepository:     repositories.NewApprovalStatusRepository(db),
		documentStatusRepository:     repositories.NewDocumentStatusRepository(db),
		workflowStatusRepository:     repositories.NewWorkflowStatusRepository(db),
		evaluationStatusRepository:   repositories.NewEvaluationStatusRepository(db),
		userStatusRepository:         repositories.NewUserStatusRepository(db),
		employmentStatusRepository:   repositories.NewEmploymentStatusRepository(db),
	}
}

func (s *StatisticsService) GetStatistics(ctx context.Context) (entities.Statistics, error) {
	countries, err := s.countries(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	currencies, err := s.currencies(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	identificationTypes, err := s.identificationTypes(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	contactTypes, err := s.contactTypes(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	maritalStatuses, err := s.maritalStatuses(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	taskStatuses, err := s.taskStatuses(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	approvalStatuses, err := s.approvalStatuses(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	documentStatuses, err := s.documentStatuses(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	workflowStatuses, err := s.workflowStatuses(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	evaluationStatuses, err := s.evaluationStatuses(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	userStatuses, err := s.userStatuses(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	employmentStatuses, err := s.employmentStatuses(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}

	return entities.Statistics{
		Countries:           countries,
		Currencies:          currencies,
		IdentificationTypes: identificationTypes,
		ContactTypes:        contactTypes,
		MaritalStatuses:     maritalStatuses,
		TaskStatuses:        taskStatuses,
		ApprovalStatuses:    approvalStatuses,
		DocumentStatuses:    documentStatuses,
		WorkflowStatuses:    workflowStatuses,
		EvaluationStatuses:  evaluationStatuses,
		UserStatuses:        userStatuses,
		EmploymentStatuses:  employmentStatuses,
	}, nil
}

func (s *StatisticsService) countries(ctx context.Context) (int64, error) {
	return s.countryRepository.Count(ctx)
}

func (s *StatisticsService) currencies(ctx context.Context) (int64, error) {
	return s.currencyRepository.Count(ctx)
}

func (s *StatisticsService) identificationTypes(ctx context.Context) (int64, error) {
	return s.identificationTypeRepository.Count(ctx)
}

func (s *StatisticsService) contactTypes(ctx context.Context) (int64, error) {
	return s.contactTypeRepository.Count(ctx)
}

func (s *StatisticsService) maritalStatuses(ctx context.Context) (int64, error) {
	return s.maritalStatusRepository.Count(ctx)
}

func (s *StatisticsService) taskStatuses(ctx context.Context) (int64, error) {
	return s.taskStatusRepository.Count(ctx)
}

func (s *StatisticsService) approvalStatuses(ctx context.Context) (int64, error) {
	return s.approvalStatusRepository.Count(ctx)
}

func (s *StatisticsService) documentStatuses(ctx context.Context) (int64, error) {
	return s.documentStatusRepository.Count(ctx)
}

func (s *StatisticsService) workflowStatuses(ctx context.Context) (int64, error) {
	return s.workflowStatusRepository.Count(ctx)
}

func (s *StatisticsService) evaluationStatuses(ctx context.Context) (int64, error) {
	return s.evaluationStatusRepository.Count(ctx)
}

func (s *StatisticsService) userStatuses(ctx context.Context) (int64, error) {
	return s.userStatusRepository.Count(ctx)
}

func (s *StatisticsService) employmentStatuses(ctx context.Context) (int64, error) {
	return s.employmentStatusRepository.Count(ctx)
}

