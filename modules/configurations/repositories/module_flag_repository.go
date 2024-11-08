package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type ModuleFlagRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.ModuleFlag]
}

func NewModuleFlagRepository(db *database.Database) *ModuleFlagRepository {
	return &ModuleFlagRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.ModuleFlag](db),
	}
}

func (repo *ModuleFlagRepository) FindAllData(ctx context.Context) ([]entities.ModuleFlagData, error) {
	var moduleFlags []entities.ModuleFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_module_flag_data").Find(&moduleFlags)
	return moduleFlags, result.Error
}

func (repo *ModuleFlagRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.ModuleFlagData, error) {
	var moduleFlags []entities.ModuleFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_module_flag_data").Limit(limit).Offset(offset).Find(&moduleFlags)
	return moduleFlags, result.Error
}

func (repo *ModuleFlagRepository) FindAllFlagsMap(ctx context.Context) (map[string]string, error) {
	var moduleFlags []entities.ModuleFlagData
	err := repo.db.WithContext(ctx).Table("configurations.view_module_flag_data").Find(&moduleFlags).Error
	if err != nil {
		return nil, err
	}
	// Create a map to store the flags | Map the results by code
	flagMap := make(map[string]string)
	for _, flag := range moduleFlags {
		flagMap[flag.Code] = flag.Status
	}
	return flagMap, nil
}

func (repo *ModuleFlagRepository) FindByModuleCode(ctx context.Context, code string) (entities.ModuleFlagData, error) {
	var moduleFlag entities.ModuleFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_module_flag_data").Where("code=?", code).First(&moduleFlag)
	return moduleFlag, result.Error
}

func (repo *ModuleFlagRepository) FindByModuleId(ctx context.Context, id int) (entities.ModuleFlagData, error) {
	var moduleFlag entities.ModuleFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_module_flag_data").Where("module_id=?", id).First(&moduleFlag)
	return moduleFlag, result.Error
}

func (repo *ModuleFlagRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.ModuleFlagData, error) {
	var moduleFlag entities.ModuleFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_module_flag_data").Where("unique_id=?", uniqueId).First(&moduleFlag)
	return moduleFlag, result.Error
}

func (repo *ModuleFlagRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.ModuleFlagData, error) {
	var moduleFlags []entities.ModuleFlagData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM configurations.view_module_flag_data WHERE module_name LIKE ? OR email LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&moduleFlags)
	return moduleFlags, result.Error
}

func (repo *ModuleFlagRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM configurations.view_module_flag_data WHERE module_name LIKE ? OR email LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *ModuleFlagRepository) ExistsByModuleId(ctx context.Context, moduleId int) (bool, error) {
	var module entities.ModuleFlag
	result := repo.db.WithContext(ctx).Where("module_id=?", moduleId).Find(&module)
	return module.FlagId != 0, result.Error
}
