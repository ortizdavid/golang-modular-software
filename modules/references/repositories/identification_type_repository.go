package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type IdentificationTypeRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.IdentificationType]
}

func NewIdentificationTypeRepository(db *database.Database) *IdentificationTypeRepository {
	return &IdentificationTypeRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.IdentificationType](db),
	}
}

func (repo *IdentificationTypeRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.IdentificationType, error) {
	var identificationTypes []entities.IdentificationType
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.identification_types WHERE type_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&identificationTypes)
	return identificationTypes, result.Error
}

func (repo *IdentificationTypeRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.identification_types WHERE type_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *IdentificationTypeRepository) ExistsByName(ctx context.Context, identificationTypeName string) (bool, error) {
	var identificationType entities.IdentificationType
	result := repo.db.WithContext(ctx).Where("type_name=?", identificationTypeName).Find(&identificationType)
	return identificationType.TypeId != 0, result.Error
}
