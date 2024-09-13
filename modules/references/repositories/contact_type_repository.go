package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type ContactTypeRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.ContactType]
}

func NewContactTypeRepository(db *database.Database) *ContactTypeRepository {
	return &ContactTypeRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.ContactType](db),
	}
}

func (repo *ContactTypeRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.ContactType, error) {
	var contactType entities.ContactType
	result := repo.db.WithContext(ctx).Table("reference.contact_types").Where("unique_id=?", uniqueId).First(&contactType)
	return contactType, result.Error
}

func (repo *ContactTypeRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.ContactType, error) {
	var contactType entities.ContactType
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&contactType)
	return contactType, result.Error
}

func (repo *ContactTypeRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.ContactType, error) {
	var contactTypes []entities.ContactType
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.contact_types WHERE type_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&contactTypes)
	return contactTypes, result.Error
}

func (repo *ContactTypeRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.contact_types WHERE type_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *ContactTypeRepository) ExistsByName(ctx context.Context, contactTypeName string) (bool, error) {
	var contactType entities.ContactType
	result := repo.db.WithContext(ctx).Where("type_name=?", contactTypeName).Find(&contactType)
	return contactType.TypeId != 0, result.Error
}
