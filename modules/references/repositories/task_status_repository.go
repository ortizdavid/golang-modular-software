package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type TaskStatusRepository struct {
	db *database.Database
}

func NewTaskStatusRepository(db *database.Database) *TaskStatusRepository {
	return &TaskStatusRepository{
		db: db,
	}
}

func (repo *TaskStatusRepository) Create(ctx context.Context, company entities.TaskStatus) error {
	result := repo.db.WithContext(ctx).Create(&company)
	return result.Error
}

func (repo *TaskStatusRepository) Update(ctx context.Context, company entities.TaskStatus) error {
	result := repo.db.WithContext(ctx).Save(&company)
	return result.Error
}

func (repo *TaskStatusRepository) Delete(ctx context.Context, company entities.TaskStatus) error {
	result := repo.db.WithContext(ctx).Delete(&company)
	return result.Error
}

func (repo *TaskStatusRepository) FindAll(ctx context.Context) ([]entities.TaskStatus, error) {
	var taskStatuses []entities.TaskStatus
	result := repo.db.WithContext(ctx).Find(&taskStatuses)
	return taskStatuses, result.Error
}

func (repo *TaskStatusRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.TaskStatus, error) {
	var taskStatuses []entities.TaskStatus
	result := repo.db.WithContext(ctx).Table("reference.task_statuses").Limit(limit).Offset(offset).Find(&taskStatuses)
	return taskStatuses, result.Error
}

func (repo *TaskStatusRepository) FindById(ctx context.Context, id int) (entities.TaskStatus, error) {
	var company entities.TaskStatus
	result := repo.db.WithContext(ctx).First(&company, id)
	return company, result.Error
}

func (repo *TaskStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.TaskStatus, error) {
	var taskStatus entities.TaskStatus
	result := repo.db.WithContext(ctx).Table("reference.task_statuses").Where("unique_id=?", uniqueId).First(&taskStatus)
	return taskStatus, result.Error
}

func (repo *TaskStatusRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.TaskStatus, error) {
	var company entities.TaskStatus
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *TaskStatusRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("reference.task_statuses").Count(&count)
	return count, result.Error
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
