package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type EvaluationStatusRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.EvaluationStatus]
}

func NewEvaluationStatusRepository(db *database.Database) *EvaluationStatusRepository {
	return &EvaluationStatusRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.EvaluationStatus](db),
	}
}

func (repo *EvaluationStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.EvaluationStatus, error) {
	var evaluationStatus entities.EvaluationStatus
	result := repo.db.WithContext(ctx).Table("reference.evaluation_statuses").Where("unique_id=?", uniqueId).First(&evaluationStatus)
	return evaluationStatus, result.Error
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
