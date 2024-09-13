package repositories

import (
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type EmailConfigurationRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.EmailConfiguration]
}

func NewEmailConfigurationRepository(db *database.Database) *EmailConfigurationRepository {
	return &EmailConfigurationRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.EmailConfiguration](db),
	}
}
