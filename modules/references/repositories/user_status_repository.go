package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type UserStatusRepository struct {
	db *database.Database
}

func NewUserStatusRepository(db *database.Database) *UserStatusRepository {
	return &UserStatusRepository{
		db: db,
	}
}

func (repo *UserStatusRepository) Create(ctx context.Context, company entities.UserStatus) error {
	result := repo.db.WithContext(ctx).Create(&company)
	return result.Error
}

func (repo *UserStatusRepository) Update(ctx context.Context, company entities.UserStatus) error {
	result := repo.db.WithContext(ctx).Save(&company)
	return result.Error
}

func (repo *UserStatusRepository) Delete(ctx context.Context, company entities.UserStatus) error {
	result := repo.db.WithContext(ctx).Delete(&company)
	return result.Error
}

func (repo *UserStatusRepository) FindAll(ctx context.Context) ([]entities.UserStatus, error) {
	var userStatuses []entities.UserStatus
	result := repo.db.WithContext(ctx).Find(&userStatuses)
	return userStatuses, result.Error
}

func (repo *UserStatusRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.UserStatus, error) {
	var userStatuses []entities.UserStatus
	result := repo.db.WithContext(ctx).Table("reference.user_statuses").Limit(limit).Offset(offset).Find(&userStatuses)
	return userStatuses, result.Error
}

func (repo *UserStatusRepository) FindById(ctx context.Context, id int) (entities.UserStatus, error) {
	var company entities.UserStatus
	result := repo.db.WithContext(ctx).First(&company, id)
	return company, result.Error
}

func (repo *UserStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.UserStatus, error) {
	var userStatus entities.UserStatus
	result := repo.db.WithContext(ctx).Table("reference.user_statuses").Where("unique_id=?", uniqueId).First(&userStatus)
	return userStatus, result.Error
}

func (repo *UserStatusRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.UserStatus, error) {
	var company entities.UserStatus
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *UserStatusRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("reference.user_statuses").Count(&count)
	return count, result.Error
}

func (repo *UserStatusRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.UserStatus, error) {
	var userStatuses []entities.UserStatus
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.user_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&userStatuses)
	return userStatuses, result.Error
}

func (repo *UserStatusRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.user_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *UserStatusRepository) ExistsByName(ctx context.Context, userStatusName string) (bool, error) {
	var userStatus entities.UserStatus
	result := repo.db.WithContext(ctx).Where("status_name=?", userStatusName).Find(&userStatus)
	return userStatus.StatusId != 0, result.Error
}
