package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
)

type ProjectRepository struct {
	db *database.Database
}

func NewProjectRepository(db *database.Database) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (repo *ProjectRepository) Create(ctx context.Context, company entities.Project) error {
	result := repo.db.WithContext(ctx).Create(&company)
	return result.Error
}

func (repo *ProjectRepository) Update(ctx context.Context, company entities.Project) error {
	result := repo.db.WithContext(ctx).Save(&company)
	return result.Error
}

func (repo *ProjectRepository) Delete(ctx context.Context, company entities.Project) error {
	result := repo.db.WithContext(ctx).Delete(&company)
	return result.Error
}

func (repo *ProjectRepository) FindAll(ctx context.Context) ([]entities.Project, error) {
	var projects []entities.Project
	result := repo.db.WithContext(ctx).Find(&projects)
	return projects, result.Error
}

func (repo *ProjectRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.ProjectData, error) {
	var projects []entities.ProjectData
	result := repo.db.WithContext(ctx).Table("company.view_project_data").Limit(limit).Offset(offset).Find(&projects)
	return projects, result.Error
}

func (repo *ProjectRepository) FindById(ctx context.Context, id int) (entities.Project, error) {
	var company entities.Project
	result := repo.db.WithContext(ctx).First(&company, id)
	return company, result.Error
}

func (repo *ProjectRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Project, error) {
	var company entities.Project
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *ProjectRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.ProjectData, error) {
	var company entities.ProjectData
	result := repo.db.WithContext(ctx).Table("company.view_project_data").Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *ProjectRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("company.projects").Count(&count)
	return count, result.Error
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