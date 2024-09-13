package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type JobTitleRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.JobTitle]
}

func NewJobTitleRepository(db *database.Database) *JobTitleRepository {
	return &JobTitleRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.JobTitle](db),
	}
}

func (repo *JobTitleRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.JobTitle, error) {
	var jobTitle entities.JobTitle
	result := repo.db.WithContext(ctx).Table("employees.job_titles").Where("unique_id=?", uniqueId).First(&jobTitle)
	return jobTitle, result.Error
}

func (repo *JobTitleRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.JobTitle, error) {
	var jobTitles []entities.JobTitle
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM employees.job_titles WHERE title_name LIKE ?", likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&jobTitles)
	return jobTitles, result.Error
}

func (repo *JobTitleRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM employees.job_titles WHERE title_name LIKE ?", likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *JobTitleRepository) ExistsByName(ctx context.Context, jobTitleName string) (bool, error) {
	var jobTitle entities.JobTitle
	result := repo.db.WithContext(ctx).Where("title_name=?", jobTitleName).Find(&jobTitle)
	return jobTitle.JobTitleId != 0, result.Error
}
