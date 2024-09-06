package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type StatisticsRepository struct {
	db *database.Database
}

func NewStatisticsRepository(db *database.Database) *StatisticsRepository {
	return &StatisticsRepository{
		db: db,
	}
}

func (repo *StatisticsRepository) GetStatistics(ctx context.Context) (entities.Statistics, error) {
	var statistics entities.Statistics
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM reference.view_statistics_data;").Scan(&statistics)
	return statistics, result.Error
}