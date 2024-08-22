package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type MaritalStatusRepository struct {
	db *database.Database
}

func NewMaritalStatusRepository(db *database.Database) *MaritalStatusRepository {
	return &MaritalStatusRepository{
		db: db,
	}
}

func (repo *MaritalStatusRepository) Create(ctx context.Context, maritalStatus entities.MaritalStatus) error {
	result := repo.db.WithContext(ctx).Create(&maritalStatus)
	return result.Error
}

func (repo *MaritalStatusRepository) Update(ctx context.Context, maritalStatus entities.MaritalStatus) error {
	result := repo.db.WithContext(ctx).Save(&maritalStatus)
	return result.Error
}

func (repo *MaritalStatusRepository) Delete(ctx context.Context, maritalStatus entities.MaritalStatus) error {
	result := repo.db.WithContext(ctx).Delete(&maritalStatus)
	return result.Error
}

func (repo *MaritalStatusRepository) FindAll(ctx context.Context) ([]entities.MaritalStatus, error) {
	var maritalStatuses []entities.MaritalStatus
	result := repo.db.WithContext(ctx).Find(&maritalStatuses)
	return maritalStatuses, result.Error
}

func (repo *MaritalStatusRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.MaritalStatus, error) {
	var maritalStatuses []entities.MaritalStatus
	result := repo.db.WithContext(ctx).Table("reference.marital_statuses").Limit(limit).Offset(offset).Find(&maritalStatuses)
	return maritalStatuses, result.Error
}

func (repo *MaritalStatusRepository) FindById(ctx context.Context, id int) (entities.MaritalStatus, error) {
	var maritalStatus entities.MaritalStatus
	result := repo.db.WithContext(ctx).First(&maritalStatus, id)
	return maritalStatus, result.Error
}

func (repo *MaritalStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.MaritalStatus, error) {
	var maritalStatus entities.MaritalStatus
	result := repo.db.WithContext(ctx).Table("reference.marital_statuses").Where("unique_id=?", uniqueId).First(&maritalStatus)
	return maritalStatus, result.Error
}

func (repo *MaritalStatusRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.MaritalStatus, error) {
	var maritalStatus entities.MaritalStatus
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&maritalStatus)
	return maritalStatus, result.Error
}

func (repo *MaritalStatusRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("reference.marital_statuses").Count(&count)
	return count, result.Error
}

func (repo *MaritalStatusRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.MaritalStatus, error) {
	var maritalStatuses []entities.MaritalStatus
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.marital_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&maritalStatuses)
	return maritalStatuses, result.Error
}

func (repo *MaritalStatusRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.marital_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *MaritalStatusRepository) ExistsByName(ctx context.Context, maritalStatusName string) (bool, error) {
	var maritalStatus entities.MaritalStatus
	result := repo.db.WithContext(ctx).Where("status_name=?", maritalStatusName).Find(&maritalStatus)
	return maritalStatus.StatusId != 0, result.Error
}
