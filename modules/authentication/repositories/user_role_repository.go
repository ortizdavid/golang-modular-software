package repositories

import (
	"context"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

type UserRoleRepository struct {
	db *database.Database
}

func NewUserRoleRepository(db *database.Database) *UserRoleRepository {
	return &UserRoleRepository {
		db: db,
	}
}

func (repo *UserRoleRepository) Create(ctx context.Context, userRole entities.UserRole) error {
	result := repo.db.WithContext(ctx).Create(&userRole)
	return result.Error
}

func (repo *UserRoleRepository) Update(ctx context.Context, userRole entities.UserRole) error {
	result := repo.db.WithContext(ctx).Save(&userRole)
	return result.Error
}

func (repo *UserRoleRepository) Delete(ctx context.Context, userRole entities.UserRole) error {
	result := repo.db.WithContext(ctx).Delete(&userRole)
	return result.Error
}

func (repo *UserRoleRepository) FindAll(ctx context.Context) ([]entities.UserRole, error) {
	var roles []entities.UserRole
	result := repo.db.WithContext(ctx).Find(&roles)
	return roles, result.Error
}

func (repo *UserRoleRepository) FindAllByUserId(ctx context.Context, userId int64) ([]entities.UserRoleData, error) {
	var roles []entities.UserRoleData
	result := repo.db.WithContext(ctx).Table("authentication.view_user_role_data").Where("user_id = ?", userId).Find(&roles)
	return roles, result.Error
}

func (repo *UserRoleRepository) FindAllDataByUserId(ctx context.Context, userId int64) ([]entities.UserRoleData, error) {
	var roles []entities.UserRoleData
	result := repo.db.WithContext(ctx).Table("view_user_role_data").Where("user_id = ?", userId).Find(&roles)
	return roles, result.Error
}

func (repo *UserRoleRepository) FindById(ctx context.Context, id int) (entities.UserRoleData, error) {
	var userRole entities.UserRoleData
	result := repo.db.WithContext(ctx).First(&userRole, id)
	return userRole, result.Error
}

func (repo *UserRoleRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.UserRole, error) {
	var userRole entities.UserRole
	result := repo.db.WithContext(ctx).Where("unique_id = ?", uniqueId).First(&userRole)
	return userRole, result.Error
}

func (repo *UserRoleRepository) FindByRoleId(ctx context.Context, roleId int) (entities.UserRole, error) {
	var role entities.UserRole
	result := repo.db.WithContext(ctx).Where("role_id=?", roleId).First(&role)
	return role, result.Error
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


func (repo *UserRoleRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.user_roles").Count(&count)
	return count, result.Error
}

func (repo *UserRoleRepository) ExistsByRoleCode(ctx context.Context, code string) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("view_user_role_data").Where("role_code = ?", code).Count(&count)
	return count > 0 , result.Error
}

func (repo *UserRoleRepository) ExistsByRoleId(ctx context.Context, roleId int) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.user_roles").Where("role_id = ?", roleId).Count(&count)
	return count > 0 , result.Error
}

func (repo *UserRoleRepository) ExistsByUserAndRole(ctx context.Context, userId int64, roleId int) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.user_roles").Where("user_id = ? AND role_id = ?", userId, roleId).Count(&count)
	return count > 0 , result.Error
}