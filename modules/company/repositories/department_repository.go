package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type DepartmentRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Department]
}

func NewDepartmentRepository(db *database.Database) *DepartmentRepository {
	return &DepartmentRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Department](db),
	}
}

func (repo *DepartmentRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.DepartmentData, error) {
	var departments []entities.DepartmentData
	result := repo.db.WithContext(ctx).Table("company.view_department_data").Limit(limit).Offset(offset).Find(&departments)
	return departments, result.Error
}

func (repo *DepartmentRepository) FindAllData(ctx context.Context) ([]entities.DepartmentData, error) {
	var departments []entities.DepartmentData
	result := repo.db.WithContext(ctx).Table("company.view_department_data").Find(&departments)
	return departments, result.Error
}

func (repo *DepartmentRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.DepartmentData, error) {
	var department entities.DepartmentData
	result := repo.db.WithContext(ctx).Table("company.view_department_data").Where("unique_id=?", uniqueId).First(&department)
	return department, result.Error
}

func (repo *DepartmentRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.DepartmentData, error) {
	var departments []entities.DepartmentData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM company.view_department_data WHERE department_name LIKE ? OR acronym LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&departments)
	return departments, result.Error
}

func (repo *DepartmentRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM company.view_department_data WHERE department_name LIKE ? OR acronym LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *DepartmentRepository) ExistsByName(ctx context.Context, companyId int, departmentName string) (bool, error) {
	var department entities.Department
	result := repo.db.WithContext(ctx).Where("company_id=? AND department_name=?", companyId, departmentName).Find(&department)
	return department.DepartmentId != 0, result.Error
}
