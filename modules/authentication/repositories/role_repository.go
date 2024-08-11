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
	result := repo.db.WithContext(ctx).Table("authentication.view_role_data").Limit(limit).Offset(offset).Find(&roles)
	return roles, result.Error
}

func (repo *RoleRepository) FindById(ctx context.Context, id int) (entities.Role, error) {
	var role entities.Role
	result := repo.db.WithContext(ctx).First(&role, id)
	return role, result.Error
}

func (repo *RoleRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.RoleData, error) {
	var role entities.RoleData
	result := repo.db.WithContext(ctx).Table("authentication.view_role_data").Where("unique_id=?", uniqueId).First(&role)
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
	result := repo.db.WithContext(ctx).Table("authentication.roles").WithContext(ctx).Where("code = ?", code).Count(&count)
	return count > 0 , result.Error
}

func (repo *RoleRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.roles").Where("role_name = ?", name).Count(&count)
	return count > 0 , result.Error
}

func (repo *RoleRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.RoleData, error) {
	var users []entities.RoleData
	likeParam := "%"+param+"%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM authentication.view_user_role_data WHERE role_name LIKE ? OR role_code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&users)
	return users, result.Error
}

func (repo *RoleRepository) CountByParam(ctx context.Context, param string) (int64, error) {
    var count int64
	likeParam := "%"+param+"%"
    result := repo.db.WithContext(ctx).
        Raw("SELECT COUNT(*) FROM authentication.view_user_role_data WHERE role_name LIKE ? OR role_code LIKE ?", likeParam, likeParam).
        Scan(&count)
    return count, result.Error
}
