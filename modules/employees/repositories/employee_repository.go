package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type EmployeeRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Employee]
}

func NewEmployeeRepository(db *database.Database) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Employee](db),
	}
}

func (repo *EmployeeRepository) Create(ctx context.Context, employee entities.Employee) error {
	result := repo.db.WithContext(ctx).Create(&employee)
	repo.SetLastInsertId(employee.EmployeeId)
	return result.Error
}

func (repo *EmployeeRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.EmployeeData, error) {
	var employees []entities.EmployeeData
	result := repo.db.WithContext(ctx).
		Table("employees.view_employee_data").
		Limit(limit).
		Offset(offset).Find(&employees)
	return employees, result.Error
}

func (repo *EmployeeRepository) FindByIdentificationNumber(ctx context.Context, identNumber string) (entities.Employee, error) {
	var employee entities.Employee
	result := repo.db.WithContext(ctx).First(&employee, "identification_number=?", identNumber)
	return employee, result.Error
}

func (repo *EmployeeRepository) FindAllOrdered(ctx context.Context) ([]entities.Employee, error) {
	var employees []entities.Employee
	result := repo.db.WithContext(ctx).Order("first_name ASC").Find(&employees)
	return employees, result.Error
}

func (repo *EmployeeRepository) GetDataById(ctx context.Context, id int64) (entities.EmployeeData, error) {
	var employeeData entities.EmployeeData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM employees.view_employee_data WHERE employee_id=?", id).Scan(&employeeData)
	return employeeData, result.Error
}

func (repo *EmployeeRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeeData, error) {
	var employeeData entities.EmployeeData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM employees.view_employee_data WHERE unique_id=?", uniqueId).Scan(&employeeData)
	return employeeData, result.Error
}

func (repo *EmployeeRepository) GetDataByIndentificationNumber(ctx context.Context, identNumber string) (entities.EmployeeData, error) {
	var employeeData entities.EmployeeData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM employees.view_employee_data WHERE identification_number=?", identNumber).Scan(&employeeData)
	return employeeData, result.Error
}

func (repo *EmployeeRepository) FindAllData(ctx context.Context) ([]entities.EmployeeData, error) {
	var employees []entities.EmployeeData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM employees.view_employee_data").Scan(&employees)
	return employees, result.Error
}

func (repo *EmployeeRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.EmployeeData, error) {
	var employees []entities.EmployeeData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM employees.view_employee_data WHERE first_name LIKE ? OR last_name LIKE ? OR identification_number LIKE ?", likeParam, likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&employees)
	return employees, result.Error
}

func (repo *EmployeeRepository) CountByParam(ctx context.Context, param string) (int64, error) {
    var count int64
	likeParam := "%" + param + "%"
    result := repo.db.WithContext(ctx).
        Raw("SELECT COUNT(*) FROM employees.view_employee_data WHERE first_name LIKE ? OR last_name LIKE ? OR identification_number LIKE ?", likeParam, likeParam, likeParam).
        Scan(&count)
    return count, result.Error
}

func (repo *EmployeeRepository) ExistsByIdentificationNumber(ctx context.Context, identNumber string) (bool, error) {
	return repo.ExistsField(ctx, "identification_number", identNumber)
}

