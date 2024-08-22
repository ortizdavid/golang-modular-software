package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type IdentificationTypeRepository struct {
	db *database.Database
}

func NewIdentificationTypeRepository(db *database.Database) *IdentificationTypeRepository {
	return &IdentificationTypeRepository{
		db: db,
	}
}

func (repo *IdentificationTypeRepository) Create(ctx context.Context, identificationType entities.IdentificationType) error {
	result := repo.db.WithContext(ctx).Create(&identificationType)
	return result.Error
}

func (repo *IdentificationTypeRepository) Update(ctx context.Context, identificationType entities.IdentificationType) error {
	result := repo.db.WithContext(ctx).Save(&identificationType)
	return result.Error
}

func (repo *IdentificationTypeRepository) Delete(ctx context.Context, identificationType entities.IdentificationType) error {
	result := repo.db.WithContext(ctx).Delete(&identificationType)
	return result.Error
}

func (repo *IdentificationTypeRepository) FindAll(ctx context.Context) ([]entities.IdentificationType, error) {
	var identificationTypes []entities.IdentificationType
	result := repo.db.WithContext(ctx).Find(&identificationTypes)
	return identificationTypes, result.Error
}

func (repo *IdentificationTypeRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.IdentificationType, error) {
	var identificationTypes []entities.IdentificationType
	result := repo.db.WithContext(ctx).Table("reference.identification_types").Limit(limit).Offset(offset).Find(&identificationTypes)
	return identificationTypes, result.Error
}

func (repo *IdentificationTypeRepository) FindById(ctx context.Context, id int) (entities.IdentificationType, error) {
	var identificationType entities.IdentificationType
	result := repo.db.WithContext(ctx).First(&identificationType, id)
	return identificationType, result.Error
}

func (repo *IdentificationTypeRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.IdentificationType, error) {
	var identificationType entities.IdentificationType
	result := repo.db.WithContext(ctx).Table("reference.identification_types").Where("unique_id=?", uniqueId).First(&identificationType)
	return identificationType, result.Error
}

func (repo *IdentificationTypeRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.IdentificationType, error) {
	var identificationType entities.IdentificationType
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&identificationType)
	return identificationType, result.Error
}

func (repo *IdentificationTypeRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("reference.identification_types").Count(&count)
	return count, result.Error
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
