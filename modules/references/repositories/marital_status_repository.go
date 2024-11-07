package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type MaritalStatusRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.MaritalStatus]
}

func NewMaritalStatusRepository(db *database.Database) *MaritalStatusRepository {
	return &MaritalStatusRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.MaritalStatus](db),
	}
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
