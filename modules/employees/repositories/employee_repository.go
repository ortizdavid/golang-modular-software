package repositories

import (
	"context"
	"sync"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *database.Database
	LastInsertId int64
	mu sync.Mutex
}

func NewEmployeeRepository(db *database.Database) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (repo *EmployeeRepository) BeginTransaction(ctx context.Context) (*gorm.DB, error){
	return repo.db.BeginTx(ctx)
}

func (repo *EmployeeRepository) Create(ctx context.Context, employee entities.Employee) error {
	result := repo.db.WithContext(ctx).Create(&employee)
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.LastInsertId = employee.EmployeeId
	return result.Error
}

func (repo *EmployeeRepository) CreateBatch(ctx context.Context, employees []entities.Employee) error {
	tx := repo.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	result := tx.WithContext(ctx).Create(&employees)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

func (repo *EmployeeRepository) Update(ctx context.Context, employee entities.Employee) error {
	result := repo.db.WithContext(ctx).Save(&employee)
	return result.Error
}

func (repo *EmployeeRepository) Delete(ctx context.Context, employee entities.Employee) error {
	result := repo.db.WithContext(ctx).Delete(&employee)
	return result.Error
}

func (repo *EmployeeRepository) FindAll(ctx context.Context) ([]entities.Employee, error) {
	var employees []entities.Employee
	result := repo.db.WithContext(ctx).Find(&employees)
	return employees, result.Error
}

func (repo *EmployeeRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.EmployeeData, error) {
	var employees []entities.EmployeeData
	result := repo.db.WithContext(ctx).
		Table("employees.view_employee_data").
		Limit(limit).
		Offset(offset).Find(&employees)
	return employees, result.Error
}

func (repo *EmployeeRepository) FindById(ctx context.Context, id int64) (entities.Employee, error) {
	var employee entities.Employee
	result := repo.db.WithContext(ctx).First(&employee, id)
	return employee, result.Error
}

func (repo *EmployeeRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Employee, error) {
	var employee entities.Employee
	result := repo.db.WithContext(ctx).First(&employee, "unique_id=?", uniqueId)
	return employee, result.Error
}

func (repo *EmployeeRepository) FindByName(ctx context.Context, employeeName string) (entities.Employee, error) {
	var employee entities.Employee
	result := repo.db.WithContext(ctx).First(&employee, "first_name=?", employeeName)
	return employee, result.Error
}

func (repo *EmployeeRepository) FindByIdentificationNumber(ctx context.Context, identNumber string) (entities.Employee, error) {
	var employee entities.Employee
	result := repo.db.WithContext(ctx).First(&employee, "identification_number=?", identNumber)
	return employee, result.Error
}

func (repo *EmployeeRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("employees.employees").Count(&count)
	return count, result.Error
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

func (repo *EmployeeRepository) FindAllData(ctx context.Context) ([]entities.EmployeeData, error) {
	var employees []entities.EmployeeData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM employees.view_employee_data").Scan(&employees)
	return employees, result.Error
}

func (repo *EmployeeRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.EmployeeData, error) {
	var employees []entities.EmployeeData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM employees.view_employee_data LIMIT ? OFFSET ?", limit, offset).Scan(&employees)
	return employees, result.Error
}

func (repo *EmployeeRepository) ExistsByIdentNumber(ctx context.Context, identNumber string) (bool, error) {
	var employee entities.Employee
	result := repo.db.WithContext(ctx).Where("identification_number=?", identNumber).Find(&employee)
	return employee.EmployeeId !=0 , result.Error
}

func (repo *EmployeeRepository) GetDataByIdentificationNumber(ctx context.Context, identNumber string) (entities.EmployeeData, error) {
	var employeeData entities.EmployeeData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM employees.view_employee_data WHERE identification_number=?", identNumber).Scan(&employeeData)
	return employeeData, result.Error
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
