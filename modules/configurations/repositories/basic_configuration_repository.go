package repositories

import (
	"context"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"gorm.io/gorm"
)

type BasicConfigurationRepository struct {
	db *gorm.DB
}

func NewBasicConfigurationRepository(db *gorm.DB) *BasicConfigurationRepository {
	return &BasicConfigurationRepository{
		db: db,
	}
}

func (repo *BasicConfigurationRepository) Update(ctx context.Context, configuration entities.BasicConfiguration) error {
	result := repo.db.WithContext(ctx).Save(&configuration)
	return result.Error
}

func (repo *BasicConfigurationRepository) FindById(ctx context.Context, id int) (entities.BasicConfiguration, error) {
	var conf entities.BasicConfiguration
	result := repo.db.WithContext(ctx).First(&conf, id)
	return conf, result.Error
}

func (repo *BasicConfigurationRepository) FindFirst(ctx context.Context) (entities.BasicConfiguration, error) {
	conf, err := repo.FindById(ctx, 1)
	return conf, err
}