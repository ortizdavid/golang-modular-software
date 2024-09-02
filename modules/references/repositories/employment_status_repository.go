package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type EmploymentStatusRepository struct {
	db *database.Database
}

func NewEmploymentStatusRepository(db *database.Database) *EmploymentStatusRepository {
	return &EmploymentStatusRepository{
		db: db,
	}
}

func (repo *EmploymentStatusRepository) Create(ctx context.Context, employmentStatus entities.EmploymentStatus) error {
	result := repo.db.WithContext(ctx).Create(&employmentStatus)
	return result.Error
}

func (repo *EmploymentStatusRepository) Update(ctx context.Context, employmentStatus entities.EmploymentStatus) error {
	result := repo.db.WithContext(ctx).Save(&employmentStatus)
	return result.Error
}

func (repo *EmploymentStatusRepository) Delete(ctx context.Context, employmentStatus entities.EmploymentStatus) error {
	result := repo.db.WithContext(ctx).Delete(&employmentStatus)
	return result.Error
}

func (repo *EmploymentStatusRepository) FindAll(ctx context.Context) ([]entities.EmploymentStatus, error) {
	var employmentStatuses []entities.EmploymentStatus
	result := repo.db.WithContext(ctx).Find(&employmentStatuses)
	return employmentStatuses, result.Error
}

func (repo *EmploymentStatusRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.EmploymentStatus, error) {
	var employmentStatuses []entities.EmploymentStatus
	result := repo.db.WithContext(ctx).Table("reference.employment_statuses").Limit(limit).Offset(offset).Find(&employmentStatuses)
	return employmentStatuses, result.Error
}

func (repo *EmploymentStatusRepository) FindById(ctx context.Context, id int) (entities.EmploymentStatus, error) {
	var employmentStatus entities.EmploymentStatus
	result := repo.db.WithContext(ctx).First(&employmentStatus, id)
	return employmentStatus, result.Error
}

func (repo *EmploymentStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.EmploymentStatus, error) {
	var employmentStatus entities.EmploymentStatus
	result := repo.db.WithContext(ctx).Table("reference.employment_statuses").Where("unique_id=?", uniqueId).First(&employmentStatus)
	return employmentStatus, result.Error
}

func (repo *EmploymentStatusRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.EmploymentStatus, error) {
	var employmentStatus entities.EmploymentStatus
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&employmentStatus)
	return employmentStatus, result.Error
}

func (repo *EmploymentStatusRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("reference.employment_statuses").Count(&count)
	return count, result.Error
}

func (repo *EmploymentStatusRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.EmploymentStatus, error) {
	var employmentStatuses []entities.EmploymentStatus
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.employment_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&employmentStatuses)
	return employmentStatuses, result.Error
}

func (repo *EmploymentStatusRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.employment_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *EmploymentStatusRepository) ExistsByName(ctx context.Context, employmentStatusName string) (bool, error) {
	var employmentStatus entities.EmploymentStatus
	result := repo.db.WithContext(ctx).Where("status_name=?", employmentStatusName).Find(&employmentStatus)
	return employmentStatus.StatusId != 0, result.Error
}
