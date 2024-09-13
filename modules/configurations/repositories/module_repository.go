package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type ModuleRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Module]
}

func NewModuleRepository(db *database.Database) *ModuleRepository {
	return &ModuleRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Module](db),
	}
}

func (repo *ModuleRepository) FindByCode(ctx context.Context, code string) (entities.Module, error) {
	var module entities.Module
	result := repo.db.WithContext(ctx).Where("code=?", code).First(&module)
	return module, result.Error
}

func (repo *ModuleRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.Module, error) {
	var module entities.Module
	result := repo.db.WithContext(ctx).Table("configurations.modules").Where("unique_id=?", uniqueId).First(&module)
	return module, result.Error
}

func (repo *ModuleRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.Module, error) {
	var modules []entities.Module
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM configurations.modules WHERE module_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&modules)
	return modules, result.Error
}

func (repo *ModuleRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM configurations.modules WHERE module_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *ModuleRepository) ExistsByName(ctx context.Context, moduleName string) (bool, error) {
	var module entities.Module
	result := repo.db.WithContext(ctx).Where("module_name=?", moduleName).Find(&module)
	return module.ModuleId != 0, result.Error
}

func (repo *ModuleRepository) ExistsByModule(ctx context.Context, moduleId int, moduleName string) (bool, error) {
	var module entities.Module
	result := repo.db.WithContext(ctx).Where("module_id=? AND module_name=?", moduleId, moduleName).Find(&module)
	return module.ModuleId != 0, result.Error
}
