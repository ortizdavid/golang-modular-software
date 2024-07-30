package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

type CompanyConfigurationRepository struct {
	db *database.Database
}

func NewCompanyConfigurationRepository(db *database.Database) *CompanyConfigurationRepository {
	return &CompanyConfigurationRepository{
		db: db,
	}
}

func (repo *CompanyConfigurationRepository) Create(ctx context.Context, configuration entities.CompanyConfiguration) error {
	result := repo.db.WithContext(ctx).Create(&configuration)
	return result.Error
}

func (repo *CompanyConfigurationRepository) Update(ctx context.Context, configuration entities.CompanyConfiguration) error {
	result := repo.db.WithContext(ctx).Save(&configuration)
	return result.Error
}

func (repo *CompanyConfigurationRepository) FindById(ctx context.Context, id int) (entities.CompanyConfiguration, error) {
	var configuration entities.CompanyConfiguration
	result := repo.db.WithContext(ctx).First(&configuration, id)
	return configuration, result.Error
}

func (repo *CompanyConfigurationRepository) FindFirst(ctx context.Context) (entities.CompanyConfiguration, error) {
	configuration, err := repo.FindById(ctx, 1)
	return configuration, err
}
