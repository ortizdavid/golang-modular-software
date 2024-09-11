package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

type EmployeePhoneRepository struct {
	db *database.Database
}

func NewEmployeePhoneRepository(db *database.Database) *EmployeePhoneRepository {
	return &EmployeePhoneRepository{
		db: db,
	}
}

func (repo *EmployeePhoneRepository) Create(ctx context.Context, employeePhone entities.EmployeePhone) error {
	result := repo.db.WithContext(ctx).Create(&employeePhone)
	return result.Error
}

func (repo *EmployeePhoneRepository) Update(ctx context.Context, employeePhone entities.EmployeePhone) error {
	result := repo.db.WithContext(ctx).Save(&employeePhone)
	return result.Error
}

func (repo *EmployeePhoneRepository) Delete(ctx context.Context, employeePhone entities.EmployeePhone) error {
	result := repo.db.WithContext(ctx).Delete(&employeePhone)
	return result.Error
}

func (repo *EmployeePhoneRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.EmployeePhoneData, error) {
	var employeePhones []entities.EmployeePhoneData
	result := repo.db.WithContext(ctx).Table("employees.employee_phones").Limit(limit).Offset(offset).Find(&employeePhones)
	return employeePhones, result.Error
}

func (repo *EmployeePhoneRepository) FindAllByEmployeeIdLimit(ctx context.Context, limit int, offset int, employeeId int64) ([]entities.EmployeePhoneData, error) {
	var employeePhones []entities.EmployeePhoneData
	result := repo.db.WithContext(ctx).Table("employees.view_employee_phone_data").
		Where("employee_id=?", employeeId).
		Limit(limit).Offset(offset).
		Find(&employeePhones)
	return employeePhones, result.Error
}

func (repo *EmployeePhoneRepository) FindAllByEmployeeId(ctx context.Context, employeeId int64) ([]entities.EmployeePhoneData, error) {
	var employeePhones []entities.EmployeePhoneData
	result := repo.db.WithContext(ctx).Table("employees.view_employee_phone_data").Where("employee_id=?", employeeId).Find(&employeePhones)
	return employeePhones, result.Error
}

func (repo *EmployeePhoneRepository) FindById(ctx context.Context, id int64) (entities.EmployeePhone, error) {
	var employeePhone entities.EmployeePhone
	result := repo.db.WithContext(ctx).First(&employeePhone, id)
	return employeePhone, result.Error
}

func (repo *EmployeePhoneRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeePhoneData, error) {
	var employeePhone entities.EmployeePhoneData
	result := repo.db.WithContext(ctx).Table("employees.view_employee_phone_data").Where("unique_id=?", uniqueId).First(&employeePhone)
	return employeePhone, result.Error
}

func (repo *EmployeePhoneRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeePhone, error) {
	var employeePhone entities.EmployeePhone
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&employeePhone)
	return employeePhone, result.Error
}

func (repo *EmployeePhoneRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("employees.employee_phones").Count(&count)
	return count, result.Error
}

func (repo *EmployeePhoneRepository) CountByEmployee(ctx context.Context, employeeId int64) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("employees.employee_phones").Where("employee_id=?", employeeId).Count(&count)
	return count, result.Error
}

func (repo *EmployeePhoneRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.EmployeePhone, error) {
	var employeePhones []entities.EmployeePhone
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM employees.employee_phones WHERE first_name LIKE ?", likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&employeePhones)
	return employeePhones, result.Error
}

func (repo *EmployeePhoneRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM employees.employee_phones WHERE first_name LIKE ?", likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *EmployeePhoneRepository) Exists(ctx context.Context, phone string) (bool, error) {
	var employeePhone entities.EmployeePhone
	result := repo.db.WithContext(ctx).Where("phone_number=?", phone).Find(&employeePhone)
	return employeePhone.PhoneId != 0, result.Error
}