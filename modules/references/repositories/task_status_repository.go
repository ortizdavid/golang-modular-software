package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type TaskStatusRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.TaskStatus]
}

func NewTaskStatusRepository(db *database.Database) *TaskStatusRepository {
	return &TaskStatusRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.TaskStatus](db),
	}
}

func (repo *TaskStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.TaskStatus, error) {
	var taskStatus entities.TaskStatus
	result := repo.db.WithContext(ctx).Table("reference.task_statuses").Where("unique_id=?", uniqueId).First(&taskStatus)
	return taskStatus, result.Error
}

func (repo *TaskStatusRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.TaskStatus, error) {
	var taskStatuses []entities.TaskStatus
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.task_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&taskStatuses)
	return taskStatuses, result.Error
}

func (repo *TaskStatusRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.task_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *TaskStatusRepository) ExistsByName(ctx context.Context, taskStatusName string) (bool, error) {
	var taskStatus entities.TaskStatus
	result := repo.db.WithContext(ctx).Where("status_name=?", taskStatusName).Find(&taskStatus)
	return taskStatus.StatusId != 0, result.Error
}
