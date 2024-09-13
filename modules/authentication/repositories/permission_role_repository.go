package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"

)

type PermissionRoleRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.PermissionRole]
}

func NewPermissionRoleRepository(db *database.Database) *PermissionRoleRepository {
	return &PermissionRoleRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.PermissionRole](db),
	}
}

func (repo *PermissionRoleRepository) FindAllDataByRoleId(ctx context.Context, roleId int) ([]entities.PermissionRoleData, error) {
	var roles []entities.PermissionRoleData
	result := repo.db.WithContext(ctx).Table("authentication.view_permission_role_data").Where("role_id = ?", roleId).Find(&roles)
	return roles, result.Error
}

func (repo *PermissionRoleRepository) FindAllDataByUserId(ctx context.Context, userId int64) ([]entities.PermissionRoleData, error) {
	var roles []entities.PermissionRoleData
	result := repo.db.WithContext(ctx).Table("view_user_role_data").Where("user_id = ?", userId).Find(&roles)
	return roles, result.Error
}

func (repo *PermissionRoleRepository) FindByRoleId(ctx context.Context, roleId int) (entities.PermissionRole, error) {
	var permissionRole entities.PermissionRole
	result := repo.db.WithContext(ctx).Where("role_id=?", roleId).First(&permissionRole)
	return permissionRole, result.Error
}

func (repo *PermissionRoleRepository) ExistsByRoleCode(ctx context.Context, code string) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("view_permission_role_data").Where("role_code = ?", code).Count(&count)
	return count > 0, result.Error
}

func (repo *PermissionRoleRepository) ExistsByPermissionId(ctx context.Context, permissionId int) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.permission_roles").Where("permission_id = ?", permissionId).Count(&count)
	return count > 0, result.Error
}

func (repo *PermissionRoleRepository) ExistsByRoleAndPermission(ctx context.Context, roleId int, permissionId int) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.permission_roles").Where("role_id = ? AND permission_id = ?", roleId, permissionId).Count(&count)
	return count > 0, result.Error
}

func (repo *PermissionRoleRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.PermissionRoleData, error) {
	var permissionRole entities.PermissionRoleData
	result := repo.db.WithContext(ctx).Table("authentication.view_permission_role_data").Where("unique_id=?", uniqueId).First(&permissionRole)
	return permissionRole, result.Error
}
