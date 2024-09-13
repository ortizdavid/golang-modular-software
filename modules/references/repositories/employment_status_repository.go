package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type EmploymentStatusRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.EmploymentStatus]
}

func NewEmploymentStatusRepository(db *database.Database) *EmploymentStatusRepository {
	return &EmploymentStatusRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.EmploymentStatus](db),
	}
}

func (repo *EmploymentStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.EmploymentStatus, error) {
	var employmentStatus entities.EmploymentStatus
	result := repo.db.WithContext(ctx).Table("reference.employment_statuses").Where("unique_id=?", uniqueId).First(&employmentStatus)
	return employmentStatus, result.Error
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
