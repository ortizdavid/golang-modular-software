package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/repositories"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

type StatisticsService struct {
	jobTitleRepository *repositories.JobTitleRepository
}

func NewStatisticsService(db *database.Database) *StatisticsService {
	return &StatisticsService{
		jobTitleRepository: repositories.NewJobTitleRepository(db),
	}
}

func (s *StatisticsService) GetStatistics(ctx context.Context) (entities.Statistics, error) {
	jobTitle, err := s.jobTitles(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	return entities.Statistics{
		Employees: 0,
		JobTitles: jobTitle,
	}, nil
}

func (s *StatisticsService) jobTitles(ctx context.Context) (int64, error) {
	return s.jobTitleRepository.Count(ctx)
}