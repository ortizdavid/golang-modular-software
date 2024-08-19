package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

type EmailConfigurationRepository struct {
	db *database.Database
}

func NewEmailConfigurationRepository(db *database.Database) *EmailConfigurationRepository {
	return &EmailConfigurationRepository{
		db: db,
	}
}

func (repo *EmailConfigurationRepository) Create(ctx context.Context, configuration entities.EmailConfiguration) error {
	result := repo.db.WithContext(ctx).Create(&configuration)
	return result.Error
}

func (repo *EmailConfigurationRepository) Update(ctx context.Context, configuration entities.EmailConfiguration) error {
	result := repo.db.WithContext(ctx).Save(&configuration)
	return result.Error
}

func (repo *EmailConfigurationRepository) FindById(ctx context.Context, id int) (entities.EmailConfiguration, error) {
	var conf entities.EmailConfiguration
	result := repo.db.WithContext(ctx).First(&conf, id)
	return conf, result.Error
}

func (repo *EmailConfigurationRepository) FindFirst(ctx context.Context) (entities.EmailConfiguration, error) {
	var conf entities.EmailConfiguration
	result := repo.db.WithContext(ctx).First(&conf)
	return conf, result.Error
}

func (repo *EmailConfigurationRepository) FindLast(ctx context.Context) (entities.EmailConfiguration, error) {
	var conf entities.EmailConfiguration
	result := repo.db.WithContext(ctx).Last(&conf)
	return conf, result.Error
}