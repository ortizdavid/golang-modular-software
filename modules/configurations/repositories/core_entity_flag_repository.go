package repositories

import (
	"context"
	
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

type CoreEntityFlagRepository struct {
	db *database.Database
}

func NewCoreEntityFlagRepository(db *database.Database) *CoreEntityFlagRepository {
	return &CoreEntityFlagRepository{
		db: db,
	}
}

func (repo *CoreEntityFlagRepository) Create(ctx context.Context, moduleFlag entities.CoreEntityFlag) error {
	result := repo.db.WithContext(ctx).Create(&moduleFlag)
	return result.Error
}

func (repo *CoreEntityFlagRepository) Update(ctx context.Context, moduleFlag entities.CoreEntityFlag) error {
	result := repo.db.WithContext(ctx).Save(&moduleFlag)
	return result.Error
}

func (repo *CoreEntityFlagRepository) UpdateBatch(ctx context.Context, moduleFlags []entities.CoreEntityFlag) error {
    result := repo.db.WithContext(ctx).Save(&moduleFlags)
    return result.Error
}

func (repo *CoreEntityFlagRepository) Delete(ctx context.Context, moduleFlag entities.CoreEntityFlag) error {
	result := repo.db.WithContext(ctx).Delete(&moduleFlag)
	return result.Error
}

func (repo *CoreEntityFlagRepository) FindAll(ctx context.Context) ([]entities.CoreEntityFlagData, error) {
	var moduleFlags []entities.CoreEntityFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_core_entity_flag_data").Find(&moduleFlags)
	return moduleFlags, result.Error
}

func (repo *CoreEntityFlagRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.CoreEntityFlagData, error) {
	var moduleFlags []entities.CoreEntityFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_core_entity_flag_data").Limit(limit).Offset(offset).Find(&moduleFlags)
	return moduleFlags, result.Error
}

func (repo *CoreEntityFlagRepository) FindById(ctx context.Context, id int) (entities.CoreEntityFlag, error) {
	var moduleFlag entities.CoreEntityFlag
	result := repo.db.WithContext(ctx).First(&moduleFlag, id)
	return moduleFlag, result.Error
}

func (repo *CoreEntityFlagRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.CoreEntityFlag, error) {
	var moduleFlag entities.CoreEntityFlag
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&moduleFlag)
	return moduleFlag, result.Error
}

func (repo *CoreEntityFlagRepository) FindByModule(ctx context.Context, module string) (entities.CoreEntityFlagData, error) {
	var moduleFlag entities.CoreEntityFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_core_entity_flag_data").Where("module_code=?", module)
	return moduleFlag, result.Error
}

func (repo *CoreEntityFlagRepository) FindByCoreEntity(ctx context.Context, module string) (entities.CoreEntityFlagData, error) {
	var moduleFlag entities.CoreEntityFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_core_entity_flag_data").Where("code=?", module)
	return moduleFlag, result.Error
}

func (repo *CoreEntityFlagRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.CoreEntityFlagData, error) {
	var moduleFlag entities.CoreEntityFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_core_entity_flag_data").Where("unique_id=?", uniqueId).First(&moduleFlag)
	return moduleFlag, result.Error
}

func (repo *CoreEntityFlagRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("configurations.core_entity_flag").Count(&count)
	return count, result.Error
}

func (repo *CoreEntityFlagRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.CoreEntityFlagData, error) {
	var moduleFlags []entities.CoreEntityFlagData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM configurations.view_core_entity_flag_data WHERE module_name LIKE ? OR email LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&moduleFlags)
	return moduleFlags, result.Error
}

func (repo *CoreEntityFlagRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM configurations.view_core_entity_flag_data WHERE module_name LIKE ? OR email LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *CoreEntityFlagRepository) ExistsByModuleId(ctx context.Context, moduleId int) (bool, error) {
	var module entities.CoreEntityFlag
	result := repo.db.WithContext(ctx).Where("module_id=?", moduleId).Find(&module)
	return module.FlagId != 0, result.Error
}
