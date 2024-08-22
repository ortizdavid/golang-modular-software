package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type ContactTypeRepository struct {
	db *database.Database
}

func NewContactTypeRepository(db *database.Database) *ContactTypeRepository {
	return &ContactTypeRepository{
		db: db,
	}
}

func (repo *ContactTypeRepository) Create(ctx context.Context, contactType entities.ContactType) error {
	result := repo.db.WithContext(ctx).Create(&contactType)
	return result.Error
}

func (repo *ContactTypeRepository) Update(ctx context.Context, contactType entities.ContactType) error {
	result := repo.db.WithContext(ctx).Save(&contactType)
	return result.Error
}

func (repo *ContactTypeRepository) Delete(ctx context.Context, contactType entities.ContactType) error {
	result := repo.db.WithContext(ctx).Delete(&contactType)
	return result.Error
}

func (repo *ContactTypeRepository) FindAll(ctx context.Context) ([]entities.ContactType, error) {
	var contactTypes []entities.ContactType
	result := repo.db.WithContext(ctx).Find(&contactTypes)
	return contactTypes, result.Error
}

func (repo *ContactTypeRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.ContactType, error) {
	var contactTypes []entities.ContactType
	result := repo.db.WithContext(ctx).Table("reference.contact_types").Limit(limit).Offset(offset).Find(&contactTypes)
	return contactTypes, result.Error
}

func (repo *ContactTypeRepository) FindById(ctx context.Context, id int) (entities.ContactType, error) {
	var contactType entities.ContactType
	result := repo.db.WithContext(ctx).First(&contactType, id)
	return contactType, result.Error
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

func (repo *ContactTypeRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("reference.contact_types").Count(&count)
	return count, result.Error
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
