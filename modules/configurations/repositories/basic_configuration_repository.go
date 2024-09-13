package repositories

import (
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type BasicConfigurationRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.BasicConfiguration]
}

func NewBasicConfigurationRepository(db *database.Database) *BasicConfigurationRepository {
	return &BasicConfigurationRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.BasicConfiguration](db),
	}
}
