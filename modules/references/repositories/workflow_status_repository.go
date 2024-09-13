package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type WorkflowStatusRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.WorkflowStatus]
}

func NewWorkflowStatusRepository(db *database.Database) *WorkflowStatusRepository {
	return &WorkflowStatusRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.WorkflowStatus](db),
	}
}

func (repo *WorkflowStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.WorkflowStatus, error) {
	var workflowStatus entities.WorkflowStatus
	result := repo.db.WithContext(ctx).Table("reference.workflow_statuses").Where("unique_id=?", uniqueId).First(&workflowStatus)
	return workflowStatus, result.Error
}

func (repo *WorkflowStatusRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.WorkflowStatus, error) {
	var workflowStatuses []entities.WorkflowStatus
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.workflow_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&workflowStatuses)
	return workflowStatuses, result.Error
}

func (repo *WorkflowStatusRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.workflow_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *WorkflowStatusRepository) ExistsByName(ctx context.Context, workflowStatusName string) (bool, error) {
	var workflowStatus entities.WorkflowStatus
	result := repo.db.WithContext(ctx).Where("status_name=?", workflowStatusName).Find(&workflowStatus)
	return workflowStatus.StatusId != 0, result.Error
}
