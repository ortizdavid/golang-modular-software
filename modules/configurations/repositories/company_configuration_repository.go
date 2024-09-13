package repositories

import (
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type CompanyConfigurationRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.CompanyConfiguration]
}

func NewCompanyConfigurationRepository(db *database.Database) *CompanyConfigurationRepository {
	return &CompanyConfigurationRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.CompanyConfiguration](db),
	}
}
