package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"

)

type PermissionRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Permission]
}

func NewPermissionRepository(db *database.Database) *PermissionRepository {
	return &PermissionRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Permission](db),
	}
}

func (repo *PermissionRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.PermissionData, error) {
	var permissions []entities.PermissionData
	result := repo.db.WithContext(ctx).Table("authentication.view_permission_data").Limit(limit).Offset(offset).Find(&permissions)
	return permissions, result.Error
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

func (repo *PermissionRepository) ExistsByCode(ctx context.Context, code string) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.permissions").Where("code = ?", code).Count(&count)
	return count > 0, result.Error
}

func (repo *PermissionRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.PermissionData, error) {
	var permissions []entities.PermissionData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM authentication.view_permission_data WHERE permission_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&permissions)
	return permissions, result.Error
}

func (repo *PermissionRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM authentication.view_permission_data WHERE permission_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *PermissionRepository) FindUnassignedPermissionsByRole(ctx context.Context, roleId int) ([]entities.Permission, error) {
	var permissions []entities.Permission
	result := repo.db.WithContext(ctx).
		Table("authentication.permissions").
		Where("permission_id NOT IN(SELECT permission_id FROM authentication.permission_roles WHERE role_id = ?)", roleId).
		Find(&permissions)
	return permissions, result.Error
}
