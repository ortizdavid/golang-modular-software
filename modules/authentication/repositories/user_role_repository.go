package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"

)

type UserRoleRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.UserRole]
}

func NewUserRoleRepository(db *database.Database) *UserRoleRepository {
	return &UserRoleRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.UserRole](db),
	}
}

func (repo *UserRoleRepository) FindAllDataByUserId(ctx context.Context, userId int64) ([]entities.UserRoleData, error) {
	var userRoles []entities.UserRoleData
	result := repo.db.WithContext(ctx).Table("authentication.view_user_role_data").Where("user_id = ?", userId).Find(&userRoles)
	return userRoles, result.Error
}

func (repo *UserRoleRepository) GetDataById(ctx context.Context, userRoleId int) (entities.UserRoleData, error) {
	var userRole entities.UserRoleData
	result := repo.db.WithContext(ctx).Table("authentication.view_user_role_data").Where("user_role_id=?", userRoleId).First(&userRole)
	return userRole, result.Error
}

func (repo *UserRoleRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.UserRoleData, error) {
	var userRole entities.UserRoleData
	result := repo.db.WithContext(ctx).Table("authentication.view_user_role_data").Where("unique_id=?", uniqueId).First(&userRole)
	return userRole, result.Error
}

func (repo *UserRoleRepository) ExistsByRoleCode(ctx context.Context, code string) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("view_user_role_data").Where("role_code = ?", code).Count(&count)
	return count > 0, result.Error
}

func (repo *UserRoleRepository) ExistsByRoleId(ctx context.Context, roleId int) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.user_roles").Where("role_id = ?", roleId).Count(&count)
	return count > 0, result.Error
}

func (repo *UserRoleRepository) ExistsByUserAndRole(ctx context.Context, userId int64, roleId int) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.user_roles").Where("user_id = ? AND role_id = ?", userId, roleId).Count(&count)
	return count > 0, result.Error
}
