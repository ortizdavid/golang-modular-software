package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
)
type StatisticsService struct {
	repository *repositories.StatisticsRepository
}

func NewStatisticsService(db *database.Database) *StatisticsService {
	return &StatisticsService{
		repository: repositories.NewStatisticsRepository(db),
	}
}

func (s *StatisticsService) GetStatistics(ctx context.Context)  (entities.Statistics, error) {
	statistics, err := s.repository.GetStatistics(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	return statistics, nil
}
