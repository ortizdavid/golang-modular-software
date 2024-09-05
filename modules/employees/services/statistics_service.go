package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

type StatisticsService struct {
}

func NewStatisticsService(db *database.Database) *StatisticsService {
	return &StatisticsService{}
}

func (s *StatisticsService) GetStatistics(ctx context.Context) (entities.Statistics, error) {
	return entities.Statistics{
		Employees: 0,
		JobTitles: 0,
	}, nil
}