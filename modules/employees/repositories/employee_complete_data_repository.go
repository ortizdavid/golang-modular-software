package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

type EmployeeCompleteDataRepository struct {
	db *database.Database
}

func NewEmployeeCompleteDataRepository(db *database.Database) *EmployeeCompleteDataRepository {
	return &EmployeeCompleteDataRepository{
		db: db,
	}
}

func (repo *EmployeeCompleteDataRepository) GetByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeeCompleteData, error) {
	var completeData entities.EmployeeCompleteData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM ? WHERE unique_id=?", uniqueId).Scan(&completeData)
	return completeData, result.Error
}

func (repo *EmployeeCompleteDataRepository) GetByIdentificationNumber(ctx context.Context, identNumber string) (entities.EmployeeCompleteData, error) {
	var completeData entities.EmployeeCompleteData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM ? WHERE identification_number=?", identNumber).Scan(&completeData)
	return completeData, result.Error
}