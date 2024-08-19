package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
)

type DepartmentRepository struct {
	db *database.Database
}

func NewDepartmentRepository(db *database.Database) *DepartmentRepository {
	return &DepartmentRepository{
		db: db,
	}
}


func (repo *DepartmentRepository) Create(ctx context.Context, company entities.Department) error {
	result := repo.db.WithContext(ctx).Create(&company)
	return result.Error
}

func (repo *DepartmentRepository) Update(ctx context.Context, company entities.Department) error {
	result := repo.db.WithContext(ctx).Save(&company)
	return result.Error
}

func (repo *DepartmentRepository) Delete(ctx context.Context, company entities.Department) error {
	result := repo.db.WithContext(ctx).Delete(&company)
	return result.Error
}

func (repo *DepartmentRepository) FindAll(ctx context.Context) ([]entities.Department, error) {
	var departments []entities.Department
	result := repo.db.WithContext(ctx).Find(&departments)
	return departments, result.Error
}

func (repo *DepartmentRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.DepartmentData, error) {
	var departments []entities.DepartmentData
	result := repo.db.WithContext(ctx).Table("company.view_department_data").Limit(limit).Offset(offset).Find(&departments)
	return departments, result.Error
}

func (repo *DepartmentRepository) FindById(ctx context.Context, id int) (entities.Department, error) {
	var company entities.Department
	result := repo.db.WithContext(ctx).First(&company, id)
	return company, result.Error
}

func (repo *DepartmentRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Department, error) {
	var company entities.Department
	result := repo.db.WithContext(ctx).Where("unqiue_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *DepartmentRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.DepartmentData, error) {
	var company entities.DepartmentData
	result := repo.db.WithContext(ctx).Table("company.view_department_data").Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *DepartmentRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("company.departments").Count(&count)
	return count, result.Error
}

func (repo *DepartmentRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.DepartmentData, error) {
	var departments []entities.DepartmentData
	likeParam := "%"+param+"%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM company.view_department_data WHERE department_name LIKE ? OR acronym LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&departments)
	return departments, result.Error
}

func (repo *DepartmentRepository) CountByParam(ctx context.Context, param string) (int64, error) {
    var count int64
	likeParam := "%"+param+"%"
    result := repo.db.WithContext(ctx).
        Raw("SELECT COUNT(*) FROM company.view_department_data WHERE department_name LIKE ? OR acronym LIKE ?", likeParam, likeParam).
        Scan(&count)
    return count, result.Error
}

func (repo *DepartmentRepository) ExistsByName(ctx context.Context, companyId int, departmentName string) (bool, error) {
	var department entities.Department
	result := repo.db.WithContext(ctx).Where("company_id=? AND department_name=?", companyId, departmentName).Find(&department)
	return department.DepartmentId !=0 , result.Error
}