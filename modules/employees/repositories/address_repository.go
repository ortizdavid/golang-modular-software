package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type AddressRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Address]
}

func NewAddressRepository(db *database.Database) *AddressRepository {
	return &AddressRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Address](db),
	}
}

func (repo *AddressRepository) FindAllByEmployeeId(ctx context.Context, employeeId int64) ([]entities.AddressData, error) {
	var addresses []entities.AddressData
	result := repo.db.WithContext(ctx).Table("employees.view_address_data").Where("employee_id=?", employeeId).Find(&addresses)
	return addresses, result.Error
}

func (repo *AddressRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.AddressData, error) {
	var address entities.AddressData
	result := repo.db.WithContext(ctx).Table("employees.view_address_data").Where("unique_id=?", uniqueId).First(&address)
	return address, result.Error
}

func (repo *AddressRepository) CountByEmployee(ctx context.Context, employeeId int64) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("employees.address").Where("employee_id=?", employeeId).Count(&count)
	return count, result.Error
}

func (repo *AddressRepository) Exists(ctx context.Context, request entities.CreateAddressRequest) (bool, error) {
	var address entities.Address
	result := repo.db.WithContext(ctx).Where("state=? AND employee_id=?", request.State, request.EmployeeId).Find(&address)
	return address.AddressId != 0, result.Error
}

func (repo *AddressRepository) UpdateCurrent(ctx context.Context, employeeId int64) error {
	result := repo.db.Raw("UPDATE employees.address SET is_current = false WHERE employee_id = ?", employeeId)
	return result.Error
}