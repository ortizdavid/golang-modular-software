package repositories

import (
	"context"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

type PermissionRoleRepository struct {
	db *database.Database
}

func NewPermissionRoleRepository(db *database.Database) *PermissionRoleRepository {
	return &PermissionRoleRepository {
		db: db,
	}
}

func (repo *PermissionRoleRepository) Create(ctx context.Context, userRole entities.PermissionRole) error {
	result := repo.db.WithContext(ctx).Create(&userRole)
	return result.Error
}

func (repo *PermissionRoleRepository) Update(ctx context.Context, userRole entities.PermissionRole) error {
	result := repo.db.WithContext(ctx).Save(&userRole)
	return result.Error
}

func (repo *PermissionRoleRepository) Delete(ctx context.Context, permissionRole entities.PermissionRole) error {
	result := repo.db.WithContext(ctx).Delete(&permissionRole)
	return result.Error
}

func (repo *PermissionRoleRepository) FindAll(ctx context.Context) ([]entities.PermissionRole, error) {
	var roles []entities.PermissionRole
	result := repo.db.WithContext(ctx).Find(&roles)
	return roles, result.Error
}

func (repo *PermissionRoleRepository) FindAllByRoleId(ctx context.Context, roleId int) ([]entities.PermissionRoleData, error) {
	var roles []entities.PermissionRoleData
	result := repo.db.WithContext(ctx).Table("authentication.view_permission_role_data").Where("role_id = ?", roleId).Find(&roles)
	return roles, result.Error
}

func (repo *PermissionRoleRepository) FindAllDataByUserId(ctx context.Context, userId int64) ([]entities.PermissionRoleData, error) {
	var roles []entities.PermissionRoleData
	result := repo.db.WithContext(ctx).Table("view_user_role_data").Where("user_id = ?", userId).Find(&roles)
	return roles, result.Error
}

func (repo *PermissionRoleRepository) FindById(ctx context.Context, id int) (entities.PermissionRole, error) {
	var permissionRole entities.PermissionRole
	result := repo.db.WithContext(ctx).First(&permissionRole, id)
	return permissionRole, result.Error
}

func (repo *PermissionRoleRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.PermissionRole, error) {
	var permissionRole entities.PermissionRole
	result := repo.db.WithContext(ctx).Where("unique_id = ?", uniqueId).First(&permissionRole)
	return permissionRole, result.Error
}

func (repo *PermissionRoleRepository) FindByRoleId(ctx context.Context, roleId int) (entities.PermissionRole, error) {
	var permissionRole entities.PermissionRole
	result := repo.db.WithContext(ctx).Where("role_id=?", roleId).First(&permissionRole)
	return permissionRole, result.Error
}

func (repo *PermissionRoleRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.permission_roles").Count(&count)
	return count, result.Error
}

func (repo *PermissionRoleRepository) ExistsByRoleCode(ctx context.Context, code string) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("view_permission_role_data").Where("role_code = ?", code).Count(&count)
	return count > 0 , result.Error
}

func (repo *PermissionRoleRepository) ExistsByPermissionId(ctx context.Context, permissionId int) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.permission_roles").Where("permission_id = ?", permissionId).Count(&count)
	return count > 0 , result.Error
}

func (repo *PermissionRoleRepository) ExistsByRoleAndPermission(ctx context.Context, roleId int, permissionId int) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.permission_roles").Where("role_id = ? AND permission_id = ?", roleId, permissionId).Count(&count)
	return count > 0 , result.Error
}

func (repo *PermissionRoleRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.PermissionRoleData, error) {
	var permissionRole entities.PermissionRoleData
	result := repo.db.WithContext(ctx).Table("authentication.view_permission_role_data").Where("unique_id=?", uniqueId).First(&permissionRole)
	return permissionRole, result.Error
}
