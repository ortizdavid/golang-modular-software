package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type UserAssociationRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.UserAssociation]
}

func NewUserAssociationRepository(db *database.Database) *UserAssociationRepository {
	return &UserAssociationRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.UserAssociation](db),
	}
}

func (repo *UserAssociationRepository) Exists(ctx context.Context, entityId int64) (bool, error) {
	var association entities.UserAssociation
	result := repo.db.WithContext(ctx).Where("entity_id=?", entityId).Find(&association)
	return association.AssociationId !=0 , result.Error
}