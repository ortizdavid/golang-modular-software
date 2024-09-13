package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type CoreEntityFlagRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.CoreEntityFlag]
}

func NewCoreEntityFlagRepository(db *database.Database) *CoreEntityFlagRepository {
	return &CoreEntityFlagRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.CoreEntityFlag](db),
	}
}

func (repo *CoreEntityFlagRepository) FindAllData(ctx context.Context) ([]entities.CoreEntityFlagData, error) {
	var coreEntityFlags []entities.CoreEntityFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_core_entity_flag_data").Find(&coreEntityFlags)
	return coreEntityFlags, result.Error
}

func (repo *CoreEntityFlagRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.CoreEntityFlagData, error) {
	var coreEntityFlags []entities.CoreEntityFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_core_entity_flag_data").Limit(limit).Offset(offset).Find(&coreEntityFlags)
	return coreEntityFlags, result.Error
}

func (repo *CoreEntityFlagRepository) FindAllFlagsMap(ctx context.Context) (map[string]string, error) {
	var coreEntityFlags []entities.CoreEntityFlagData
	err := repo.db.WithContext(ctx).Table("configurations.view_core_entity_flag_data").Find(&coreEntityFlags).Error
	if err != nil {
		return nil, err
	}
	// Create a map to store the flags | Map the results by code
	flagMap := make(map[string]string)
	for _, flag := range coreEntityFlags {
		flagMap[flag.Code] = flag.Status
	}
	return flagMap, nil
}

func (repo *CoreEntityFlagRepository) FindByModuleCode(ctx context.Context, module string) (entities.CoreEntityFlagData, error) {
	var coreEntityFlag entities.CoreEntityFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_core_entity_flag_data").Where("module_code=?", module)
	return coreEntityFlag, result.Error
}

func (repo *CoreEntityFlagRepository) FindByEntityCode(ctx context.Context, module string) (entities.CoreEntityFlagData, error) {
	var coreEntityFlag entities.CoreEntityFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_core_entity_flag_data").Where("code=?", module)
	return coreEntityFlag, result.Error
}

func (repo *CoreEntityFlagRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.CoreEntityFlagData, error) {
	var coreEntityFlag entities.CoreEntityFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_core_entity_flag_data").Where("unique_id=?", uniqueId).First(&coreEntityFlag)
	return coreEntityFlag, result.Error
}

func (repo *CoreEntityFlagRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.CoreEntityFlagData, error) {
	var coreEntityFlags []entities.CoreEntityFlagData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM configurations.view_core_entity_flag_data WHERE module_name LIKE ? OR email LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&coreEntityFlags)
	return coreEntityFlags, result.Error
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
