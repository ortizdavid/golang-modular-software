package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

type ModuleRepository struct {
	db *database.Database
}

func NewModuleRepository(db *database.Database) *ModuleRepository {
	return &ModuleRepository{
		db: db,
	}
}

func (repo *ModuleRepository) Create(ctx context.Context, module entities.Module) error {
	result := repo.db.WithContext(ctx).Create(&module)
	return result.Error
}

func (repo *ModuleRepository) Update(ctx context.Context, module entities.Module) error {
	result := repo.db.WithContext(ctx).Save(&module)
	return result.Error
}

func (repo *ModuleRepository) Delete(ctx context.Context, module entities.Module) error {
	result := repo.db.WithContext(ctx).Delete(&module)
	return result.Error
}

func (repo *ModuleRepository) FindAll(ctx context.Context) ([]entities.Module, error) {
	var modules []entities.Module
	result := repo.db.WithContext(ctx).Find(&modules)
	return modules, result.Error
}

func (repo *ModuleRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.Module, error) {
	var modules []entities.Module
	result := repo.db.WithContext(ctx).Table("configurations.modules").Limit(limit).Offset(offset).Find(&modules)
	return modules, result.Error
}

func (repo *ModuleRepository) FindById(ctx context.Context, id int) (entities.Module, error) {
	var module entities.Module
	result := repo.db.WithContext(ctx).First(&module, id)
	return module, result.Error
}

func (repo *ModuleRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Module, error) {
	var module entities.Module
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&module)
	return module, result.Error
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

func (repo *ModuleRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("configurations.modules").Count(&count)
	return count, result.Error
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
