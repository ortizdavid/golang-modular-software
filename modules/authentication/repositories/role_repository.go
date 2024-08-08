package repositories

import (
	"context"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

type RoleRepository struct {
	db *database.Database
}

func NewRoleRepository(db *database.Database) *RoleRepository {
	return &RoleRepository {
		db: db,
	}
}

func (repo *RoleRepository) Create(ctx context.Context, role entities.Role) error {
	result := repo.db.WithContext(ctx).Create(&role)
	return result.Error
}

func (repo *RoleRepository) Update(ctx context.Context, role entities.Role) error {
	result := repo.db.WithContext(ctx).Save(&role)
	return result.Error
}

func (repo *RoleRepository) Delete(ctx context.Context, role entities.Role) error {
	result := repo.db.WithContext(ctx).Delete(&role)
	return result.Error
}

func (repo *RoleRepository) FindAll(ctx context.Context) ([]entities.Role, error) {
	var roles []entities.Role
	result := repo.db.WithContext(ctx).Find(&roles)
	return roles, result.Error
}

func (repo *RoleRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.Role, error) {
	var roles []entities.Role
	result := repo.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&roles)
	return roles, result.Error
}

func (repo *RoleRepository) FindById(ctx context.Context, id int) (entities.Role, error) {
	var role entities.Role
	result := repo.db.WithContext(ctx).First(&role, id)
	return role, result.Error
}

func (repo *RoleRepository) FindByName(ctx context.Context, name string) (entities.Role, error) {
	var role entities.Role
	result := repo.db.WithContext(ctx).Where("role_name=?", name).First(&role)
	return role, result.Error
}

func (repo *RoleRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.roles").Count(&count)
	return count, result.Error
}

func (repo *RoleRepository) ExistsByCode(ctx context.Context, code string) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).WithContext(ctx).Where("role_code = ?", code).Count(&count)
	return count > 0 , result.Error
}
