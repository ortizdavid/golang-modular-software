package repositories

import (
	"context"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
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

func (repo *RoleRepository) FindAll(ctx context.Context) ([]entities.Role, error) {
	var roles []entities.Role
	result := repo.db.WithContext(ctx).Find(&roles)
	return roles, result.Error
}

func (repo *RoleRepository) FindById(ctx context.Context, id int) (entities.Role, error) {
	var role entities.Role
	result := repo.db.WithContext(ctx).First(&role, id)
	return role, result.Error
}

func (repo *RoleRepository) FindByName(name string) (entities.Role, error) {
	var role entities.Role
	result := repo.db.Where("role_name=?", name).First(&role)
	return role, result.Error
}

func (repo *RoleRepository) Count() (int64, error) {
	var count int64
	result := repo.db.Table("authentication.roles").Count(&count)
	return count, result.Error
}