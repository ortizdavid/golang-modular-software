package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type ProjectRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Project]
}

func NewProjectRepository(db *database.Database) *ProjectRepository {
	return &ProjectRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Project](db),
	}
}

func (repo *ProjectRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.ProjectData, error) {
	var projects []entities.ProjectData
	result := repo.db.WithContext(ctx).Table("company.view_project_data").Limit(limit).Offset(offset).Find(&projects)
	return projects, result.Error
}

func (repo *ProjectRepository) FindAllData(ctx context.Context) ([]entities.ProjectData, error) {
	var projects []entities.ProjectData
	result := repo.db.WithContext(ctx).Table("company.view_project_data").Find(&projects)
	return projects, result.Error
}

func (repo *ProjectRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.ProjectData, error) {
	var project entities.ProjectData
	result := repo.db.WithContext(ctx).Table("company.view_project_data").Where("unique_id=?", uniqueId).First(&project)
	return project, result.Error
}

func (repo *ProjectRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.ProjectData, error) {
	var projects []entities.ProjectData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM company.view_project_data WHERE project_name LIKE ? OR status LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&projects)
	return projects, result.Error
}

func (repo *ProjectRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM company.view_project_data WHERE project_name LIKE ? OR status LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *ProjectRepository) ExistsByName(ctx context.Context, companyId int, projectName string) (bool, error) {
	var project entities.Project
	result := repo.db.WithContext(ctx).Where("company_id=? AND project_name=?", companyId, projectName).Find(&project)
	return project.ProjectId != 0, result.Error
}
