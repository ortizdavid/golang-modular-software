package repositories

import (
	"context"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

type PermissionRepository struct {
	db *database.Database
}

func NewPermissionRepository(db *database.Database) *PermissionRepository {
	return &PermissionRepository {
		db: db,
	}
}

func (repo *PermissionRepository) Create(ctx context.Context, permission entities.Permission) error {
	result := repo.db.WithContext(ctx).Create(&permission)
	return result.Error
}

func (repo *PermissionRepository) Update(ctx context.Context, permission entities.Permission) error {
	result := repo.db.WithContext(ctx).Save(&permission)
	return result.Error
}

func (repo *PermissionRepository) Delete(ctx context.Context, permission entities.Permission) error {
	result := repo.db.WithContext(ctx).Table("authentication.permissions").Delete(&permission)
	return result.Error
}

func (repo *PermissionRepository) FindAll(ctx context.Context) ([]entities.Permission, error) {
	var permissions []entities.Permission
	result := repo.db.WithContext(ctx).Find(&permissions)
	return permissions, result.Error
}

func (repo *PermissionRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.PermissionData, error) {
	var permissions []entities.PermissionData
	result := repo.db.WithContext(ctx).Table("authentication.view_permission_data").Limit(limit).Offset(offset).Find(&permissions)
	return permissions, result.Error
}

func (repo *PermissionRepository) FindById(ctx context.Context, id int) (entities.Permission, error) {
	var permission entities.Permission
	result := repo.db.WithContext(ctx).First(&permission, id)
	return permission, result.Error
}

func (repo *PermissionRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Permission, error) {
	var permission entities.Permission
	result := repo.db.WithContext(ctx).Where("unqiue_id=?", uniqueId).First(&permission)
	return permission, result.Error
}

func (repo *PermissionRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.PermissionData, error) {
	var permission entities.PermissionData
	result := repo.db.WithContext(ctx).Table("authentication.view_permission_data").Where("unique_id=?", uniqueId).First(&permission)
	return permission, result.Error
}

func (repo *PermissionRepository) FindByName(ctx context.Context, name string) (entities.Permission, error) {
	var permission entities.Permission
	result := repo.db.WithContext(ctx).Where("permission_name=?", name).First(&permission)
	return permission, result.Error
}

func (repo *PermissionRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.permissions").Count(&count)
	return count, result.Error
}

func (repo *PermissionRepository) ExistsByCode(ctx context.Context, code string) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.permissions").Where("code = ?", code).Count(&count)
	return count > 0 , result.Error
}

func (repo *PermissionRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.PermissionData, error) {
	var permissions []entities.PermissionData
	likeParam := "%"+param+"%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM authentication.view_permission_data WHERE permission_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&permissions)
	return permissions, result.Error
}

func (repo *PermissionRepository) CountByParam(ctx context.Context, param string) (int64, error) {
    var count int64
	likeParam := "%"+param+"%"
    result := repo.db.WithContext(ctx).
        Raw("SELECT COUNT(*) FROM authentication.view_permission_data WHERE permission_name LIKE ? OR code LIKE ?", likeParam, likeParam).
        Scan(&count)
    return count, result.Error
}
