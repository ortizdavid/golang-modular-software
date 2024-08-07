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

func (repo *UserRoleRepository) FindAll(ctx context.Context) ([]entities.UserRole, error) {
	var roles []entities.UserRole
	result := repo.db.WithContext(ctx).Find(&roles)
	return roles, result.Error
}

func (repo *UserRoleRepository) FindAllByUserId(ctx context.Context, userId int64) ([]entities.UserRole, error) {
	var roles []entities.UserRole
	result := repo.db.WithContext(ctx).Where("user_id = ?", userId).Find(&roles)
	return roles, result.Error
}

func (repo *UserRoleRepository) FindAllDataByUserId(ctx context.Context, userId int64) ([]entities.UserRoleData, error) {
	var roles []entities.UserRoleData
	result := repo.db.WithContext(ctx).Table("view_user_role_data").Where("user_id = ?", userId).Find(&roles)
	return roles, result.Error
}

func (repo *UserRoleRepository) FindById(ctx context.Context, id int) (entities.UserRole, error) {
	var role entities.UserRole
	result := repo.db.WithContext(ctx).First(&role, id)
	return role, result.Error
}

func (repo *UserRoleRepository) FindByRoleId(roleId int) (entities.UserRole, error) {
	var role entities.UserRole
	result := repo.db.Where("role_id=?", roleId).First(&role)
	return role, result.Error
}

func (repo *UserRoleRepository) FindByUserId(userId int) (entities.UserRole, error) {
	var role entities.UserRole
	result := repo.db.Where("user_id=?", userId).First(&role)
	return role, result.Error
}

func (repo *UserRoleRepository) Count() (int64, error) {
	var count int64
	result := repo.db.Table("authentication.user_roles").Count(&count)
	return count, result.Error
}

func (repo *UserRoleRepository) ExistsByCode(ctx context.Context, code string) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("view_user_role_data").Where("role_code = ?", code).Count(&count)
	return count > 0 , result.Error
}

func (repo *UserRoleRepository) ExistsByUserId(ctx context.Context, userId int64, roleId int) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.user_roles").Where("user_id = ? AND role_id = ?", userId, roleId).Count(&count)
	return count > 0 , result.Error
}