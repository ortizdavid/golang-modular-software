package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"

)

type RoleRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Role]
}

func NewRoleRepository(db *database.Database) *RoleRepository {
	return &RoleRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Role](db),
	}
}

func (repo *RoleRepository) FindAllEnabled(ctx context.Context) ([]entities.Role, error) {
	var roles []entities.Role
	result := repo.db.WithContext(ctx).Where("status='Enabled'").Find(&roles)
	return roles, result.Error
}

func (repo *RoleRepository) FindAllEnabledNotIn(ctx context.Context, values []interface{}) ([]entities.Role, error) {
	var roles []entities.Role
	result := repo.db.WithContext(ctx).Where("status='Enabled' AND code NOT IN (?)", values).Find(&roles)
	return roles, result.Error
}

/*func (repo *RoleRepository) FindAllEnabledNotIn(ctx context.Context, values []interface{}) ([]entities.Role, error) {
	var roles []entities.Role
	result := repo.db.WithContext(ctx).Where("status = ? AND code NOT IN (?)", "Enabled", values).Find(&roles)
	return roles, result.Error
}*/


func (repo *RoleRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.RoleData, error) {
	var roles []entities.RoleData
	result := repo.db.WithContext(ctx).Table("authentication.view_role_data").Limit(limit).Offset(offset).Find(&roles)
	return roles, result.Error
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

func (repo *RoleRepository) ExistsByCode(ctx context.Context, code string) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.roles").WithContext(ctx).Where("code = ?", code).Count(&count)
	return count > 0, result.Error
}

func (repo *RoleRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.roles").Where("role_name = ?", name).Count(&count)
	return count > 0, result.Error
}

func (repo *RoleRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.RoleData, error) {
	var users []entities.RoleData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM authentication.view_user_role_data WHERE role_name LIKE ? OR role_code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&users)
	return users, result.Error
}

func (repo *RoleRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM authentication.view_user_role_data WHERE role_name LIKE ? OR role_code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *RoleRepository) FindUnassignedRolesByUser(ctx context.Context, userId int64) ([]entities.Role, error) {
	var roles []entities.Role
	result := repo.db.WithContext(ctx).
		Table("authentication.roles").
		Where("role_id NOT IN(SELECT role_id FROM authentication.user_roles WHERE user_id = ?) AND status='Enabled'", userId).
		Find(&roles)
	return roles, result.Error
}
