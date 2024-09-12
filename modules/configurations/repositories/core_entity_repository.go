package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

type CoreEntityRepository struct {
	db *database.Database
}

func NewCoreEntityRepository(db *database.Database) *CoreEntityRepository {
	return &CoreEntityRepository{
		db: db,
	}
}

func (repo *CoreEntityRepository) Create(ctx context.Context, coreEntity entities.CoreEntity) error {
	result := repo.db.WithContext(ctx).Create(&coreEntity)
	return result.Error
}

func (repo *CoreEntityRepository) Update(ctx context.Context, coreEntity entities.CoreEntity) error {
	result := repo.db.WithContext(ctx).Save(&coreEntity)
	return result.Error
}

func (repo *CoreEntityRepository) Delete(ctx context.Context, coreEntity entities.CoreEntity) error {
	result := repo.db.WithContext(ctx).Delete(&coreEntity)
	return result.Error
}

func (repo *CoreEntityRepository) FindAll(ctx context.Context) ([]entities.CoreEntity, error) {
	var coreEntities []entities.CoreEntity
	result := repo.db.WithContext(ctx).Find(&coreEntities)
	return coreEntities, result.Error
}

func (repo *CoreEntityRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.CoreEntityData, error) {
	var coreEntities []entities.CoreEntityData
	result := repo.db.WithContext(ctx).Table("configurations.view_core_entity_data").Limit(limit).Offset(offset).Find(&coreEntities)
	return coreEntities, result.Error
}

func (repo *CoreEntityRepository) FindById(ctx context.Context, id int) (entities.CoreEntity, error) {
	var coreEntity entities.CoreEntity
	result := repo.db.WithContext(ctx).First(&coreEntity, id)
	return coreEntity, result.Error
}

func (repo *CoreEntityRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.CoreEntity, error) {
	var coreEntity entities.CoreEntity
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&coreEntity)
	return coreEntity, result.Error
}

func (repo *CoreEntityRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.CoreEntityData, error) {
	var coreEntity entities.CoreEntityData
	result := repo.db.WithContext(ctx).Table("configurations.view_core_entity_data").Where("unique_id=?", uniqueId).First(&coreEntity)
	return coreEntity, result.Error
}

func (repo *CoreEntityRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("configurations.core_entities").Count(&count)
	return count, result.Error
}

func (repo *CoreEntityRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.CoreEntityData, error) {
	var coreEntities []entities.CoreEntityData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM configurations.view_core_entity_data WHERE entity_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&coreEntities)
	return coreEntities, result.Error
}

func (repo *CoreEntityRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM configurations.view_core_entity_data WHERE entity_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *CoreEntityRepository) ExistsByName(ctx context.Context, featureName string) (bool, error) {
	var feature entities.CoreEntity
	result := repo.db.WithContext(ctx).Where("entity_name=?", featureName).Find(&feature)
	return feature.EntityId != 0, result.Error
}

func (repo *CoreEntityRepository) ExistsByModule(ctx context.Context, moduleId int, featureName string) (bool, error) {
	var feature entities.CoreEntity
	result := repo.db.WithContext(ctx).Where("module_id=? AND entity_name=?", moduleId, featureName).Find(&feature)
	return feature.EntityId != 0, result.Error
}
