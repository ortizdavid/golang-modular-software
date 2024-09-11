package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

type EmployeeEmailRepository struct {
	db *database.Database
}

func NewEmployeeEmailRepository(db *database.Database) *EmployeeEmailRepository {
	return &EmployeeEmailRepository{
		db: db,
	}
}

func (repo *EmployeeEmailRepository) Create(ctx context.Context, employeeEmail entities.EmployeeEmail) error {
	result := repo.db.WithContext(ctx).Create(&employeeEmail)
	return result.Error
}

func (repo *EmployeeEmailRepository) Update(ctx context.Context, employeeEmail entities.EmployeeEmail) error {
	result := repo.db.WithContext(ctx).Save(&employeeEmail)
	return result.Error
}

func (repo *EmployeeEmailRepository) Delete(ctx context.Context, employeeEmail entities.EmployeeEmail) error {
	result := repo.db.WithContext(ctx).Delete(&employeeEmail)
	return result.Error
}

func (repo *EmployeeEmailRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.EmployeeEmailData, error) {
	var employeeEmails []entities.EmployeeEmailData
	result := repo.db.WithContext(ctx).Table("employees.employee_emails").Limit(limit).Offset(offset).Find(&employeeEmails)
	return employeeEmails, result.Error
}

func (repo *EmployeeEmailRepository) FindAllByEmployeeIdLimit(ctx context.Context, limit int, offset int, employeeId int64) ([]entities.EmployeeEmailData, error) {
	var employeeEmails []entities.EmployeeEmailData
	result := repo.db.WithContext(ctx).Table("employees.view_employee_email_data").
		Where("employee_id=?", employeeId).
		Limit(limit).Offset(offset).
		Find(&employeeEmails)
	return employeeEmails, result.Error
}

func (repo *EmployeeEmailRepository) FindAllByEmployeeId(ctx context.Context, employeeId int64) ([]entities.EmployeeEmailData, error) {
	var employeeEmails []entities.EmployeeEmailData
	result := repo.db.WithContext(ctx).Table("employees.view_employee_email_data").Where("employee_id=?", employeeId).Find(&employeeEmails)
	return employeeEmails, result.Error
}

func (repo *EmployeeEmailRepository) FindById(ctx context.Context, id int64) (entities.EmployeeEmail, error) {
	var employeeEmail entities.EmployeeEmail
	result := repo.db.WithContext(ctx).First(&employeeEmail, id)
	return employeeEmail, result.Error
}

func (repo *EmployeeEmailRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeeEmailData, error) {
	var employeeEmail entities.EmployeeEmailData
	result := repo.db.WithContext(ctx).Table("employees.view_employee_email_data").Where("unique_id=?", uniqueId).First(&employeeEmail)
	return employeeEmail, result.Error
}

func (repo *EmployeeEmailRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeeEmail, error) {
	var employeeEmail entities.EmployeeEmail
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&employeeEmail)
	return employeeEmail, result.Error
}

func (repo *EmployeeEmailRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("employees.employee_emails").Count(&count)
	return count, result.Error
}

func (repo *EmployeeEmailRepository) CountByEmployee(ctx context.Context, employeeId int64) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("employees.employee_emails").Where("employee_id=?", employeeId).Count(&count)
	return count, result.Error
}

func (repo *EmployeeEmailRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.EmployeeEmail, error) {
	var employeeEmails []entities.EmployeeEmail
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM employees.employee_emails WHERE first_name LIKE ?", likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&employeeEmails)
	return employeeEmails, result.Error
}

func (repo *EmployeeEmailRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM employees.employee_emails WHERE first_name LIKE ?", likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *EmployeeEmailRepository) Exists(ctx context.Context, email string) (bool, error) {
	var employeeEmail entities.EmployeeEmail
	result := repo.db.WithContext(ctx).Where("email_address=?", email).Find(&employeeEmail)
	return employeeEmail.EmailId != 0, result.Error
}
