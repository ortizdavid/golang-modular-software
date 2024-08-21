package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type StatisticsService struct {
}

func NewStatisticsService(db *database.Database) *StatisticsService {
	return &StatisticsService{}
}

func (s *StatisticsService) GetStatistics(ctx context.Context)  (entities.Statistics, error) {
	return entities.Statistics{
		Countries:           0,
		Currencies:          0,
		IdentificationTypes: 0,
		ContactTypes:        0,
		MaritalStatuses:     0,
		TaskStatuses:        0,
		ApprovalStatuses:    0,
		DocumentStatuses:    0,
		WorkflowStatuses:    0,
	}, nil
}