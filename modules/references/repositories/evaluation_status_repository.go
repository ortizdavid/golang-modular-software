package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type EvaluationStatusRepository struct {
	db *database.Database
}

func NewEvaluationStatusRepository(db *database.Database) *EvaluationStatusRepository {
	return &EvaluationStatusRepository{
		db: db,
	}
}

func (repo *EvaluationStatusRepository) Create(ctx context.Context, evaluationStatus entities.EvaluationStatus) error {
	result := repo.db.WithContext(ctx).Create(&evaluationStatus)
	return result.Error
}

func (repo *EvaluationStatusRepository) Update(ctx context.Context, evaluationStatus entities.EvaluationStatus) error {
	result := repo.db.WithContext(ctx).Save(&evaluationStatus)
	return result.Error
}

func (repo *EvaluationStatusRepository) Delete(ctx context.Context, evaluationStatus entities.EvaluationStatus) error {
	result := repo.db.WithContext(ctx).Delete(&evaluationStatus)
	return result.Error
}

func (repo *EvaluationStatusRepository) FindAll(ctx context.Context) ([]entities.EvaluationStatus, error) {
	var evaluationStatuses []entities.EvaluationStatus
	result := repo.db.WithContext(ctx).Find(&evaluationStatuses)
	return evaluationStatuses, result.Error
}

func (repo *EvaluationStatusRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.EvaluationStatus, error) {
	var evaluationStatuses []entities.EvaluationStatus
	result := repo.db.WithContext(ctx).Table("reference.evaluation_statuses").Limit(limit).Offset(offset).Find(&evaluationStatuses)
	return evaluationStatuses, result.Error
}

func (repo *EvaluationStatusRepository) FindById(ctx context.Context, id int) (entities.EvaluationStatus, error) {
	var evaluationStatus entities.EvaluationStatus
	result := repo.db.WithContext(ctx).First(&evaluationStatus, id)
	return evaluationStatus, result.Error
}

func (repo *EvaluationStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.EvaluationStatus, error) {
	var evaluationStatus entities.EvaluationStatus
	result := repo.db.WithContext(ctx).Table("reference.evaluation_statuses").Where("unique_id=?", uniqueId).First(&evaluationStatus)
	return evaluationStatus, result.Error
}

func (repo *EvaluationStatusRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.EvaluationStatus, error) {
	var evaluationStatus entities.EvaluationStatus
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&evaluationStatus)
	return evaluationStatus, result.Error
}

func (repo *EvaluationStatusRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("reference.evaluation_statuses").Count(&count)
	return count, result.Error
}

func (repo *EvaluationStatusRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.EvaluationStatus, error) {
	var evaluationStatuses []entities.EvaluationStatus
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.evaluation_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&evaluationStatuses)
	return evaluationStatuses, result.Error
}

func (repo *EvaluationStatusRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.evaluation_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *EvaluationStatusRepository) ExistsByName(ctx context.Context, evaluationStatusName string) (bool, error) {
	var evaluationStatus entities.EvaluationStatus
	result := repo.db.WithContext(ctx).Where("status_name=?", evaluationStatusName).Find(&evaluationStatus)
	return evaluationStatus.StatusId != 0, result.Error
}
