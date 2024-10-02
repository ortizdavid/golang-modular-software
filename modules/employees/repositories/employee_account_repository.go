package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

type EmployeeAccountRepository struct {
	db *database.Database
}

func NewEmployeeAccountRepository(db *database.Database) *EmployeeAccountRepository {
	return &EmployeeAccountRepository{
		db: db,
	}
}

func (repo *EmployeeAccountRepository) GetDataByEmployeeId(ctx context.Context, id int64) (entities.EmployeeAccountData, error) {
	var accountData entities.EmployeeAccountData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM employees.view_employee_account_data WHERE employee_id=?", id).Scan(&accountData)
	return accountData, result.Error
}

func (repo *EmployeeAccountRepository) GetDataByIdentificationNumber(ctx context.Context, identNumber string) (entities.EmployeeAccountData, error) {
	var accountData entities.EmployeeAccountData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM employees.view_employee_account_data WHERE identification_number=?", identNumber).Scan(&accountData)
	return accountData, result.Error
}

func (repo *EmployeeAccountRepository) GetDataByEmployeeUniqueId(ctx context.Context, uniqueId string) (entities.EmployeeAccountData, error) {
	var accountData entities.EmployeeAccountData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM employees.view_employee_account_data WHERE employee_unique_id=?", uniqueId).Scan(&accountData)
	return accountData, result.Error
}