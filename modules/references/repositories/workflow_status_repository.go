package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type WorkflowStatusRepository struct {
	db *database.Database
}

func NewWorkflowStatusRepository(db *database.Database) *WorkflowStatusRepository {
	return &WorkflowStatusRepository{
		db: db,
	}
}

func (repo *WorkflowStatusRepository) Create(ctx context.Context, workflowStatus entities.WorkflowStatus) error {
	result := repo.db.WithContext(ctx).Create(&workflowStatus)
	return result.Error
}

func (repo *WorkflowStatusRepository) Update(ctx context.Context, workflowStatus entities.WorkflowStatus) error {
	result := repo.db.WithContext(ctx).Save(&workflowStatus)
	return result.Error
}

func (repo *WorkflowStatusRepository) Delete(ctx context.Context, workflowStatus entities.WorkflowStatus) error {
	result := repo.db.WithContext(ctx).Delete(&workflowStatus)
	return result.Error
}

func (repo *WorkflowStatusRepository) FindAll(ctx context.Context) ([]entities.WorkflowStatus, error) {
	var workflowStatuses []entities.WorkflowStatus
	result := repo.db.WithContext(ctx).Find(&workflowStatuses)
	return workflowStatuses, result.Error
}

func (repo *WorkflowStatusRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.WorkflowStatus, error) {
	var workflowStatuses []entities.WorkflowStatus
	result := repo.db.WithContext(ctx).Table("reference.workflow_statuses").Limit(limit).Offset(offset).Find(&workflowStatuses)
	return workflowStatuses, result.Error
}

func (repo *WorkflowStatusRepository) FindById(ctx context.Context, id int) (entities.WorkflowStatus, error) {
	var workflowStatus entities.WorkflowStatus
	result := repo.db.WithContext(ctx).First(&workflowStatus, id)
	return workflowStatus, result.Error
}

func (repo *WorkflowStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.WorkflowStatus, error) {
	var workflowStatus entities.WorkflowStatus
	result := repo.db.WithContext(ctx).Table("reference.workflow_statuses").Where("unique_id=?", uniqueId).First(&workflowStatus)
	return workflowStatus, result.Error
}

func (repo *WorkflowStatusRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.WorkflowStatus, error) {
	var workflowStatus entities.WorkflowStatus
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&workflowStatus)
	return workflowStatus, result.Error
}

func (repo *WorkflowStatusRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("reference.workflow_statuses").Count(&count)
	return count, result.Error
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
