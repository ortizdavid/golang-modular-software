package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
)

type OfficeRepository struct {
	db *database.Database
}

func NewOfficeRepository(db *database.Database) *OfficeRepository {
	return &OfficeRepository{
		db: db,
	}
}

func (repo *OfficeRepository) Create(ctx context.Context, office entities.Office) error {
	result := repo.db.WithContext(ctx).Create(&office)
	return result.Error
}

func (repo *OfficeRepository) Update(ctx context.Context, office entities.Office) error {
	result := repo.db.WithContext(ctx).Save(&office)
	return result.Error
}

func (repo *OfficeRepository) Delete(ctx context.Context, office entities.Office) error {
	result := repo.db.WithContext(ctx).Delete(&office)
	return result.Error
}

func (repo *OfficeRepository) FindAll(ctx context.Context) ([]entities.Office, error) {
	var offices []entities.Office
	result := repo.db.WithContext(ctx).Find(&offices)
	return offices, result.Error
}

func (repo *OfficeRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.OfficeData, error) {
	var offices []entities.OfficeData
	result := repo.db.WithContext(ctx).Table("company.view_office_data").Limit(limit).Offset(offset).Find(&offices)
	return offices, result.Error
}

func (repo *OfficeRepository) FindById(ctx context.Context, id int) (entities.Office, error) {
	var office entities.Office
	result := repo.db.WithContext(ctx).First(&office, id)
	return office, result.Error
}

func (repo *OfficeRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Office, error) {
	var office entities.Office
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&office)
	return office, result.Error
}

func (repo *OfficeRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.OfficeData, error) {
	var office entities.OfficeData
	result := repo.db.WithContext(ctx).Table("company.view_office_data").Where("unique_id=?", uniqueId).First(&office)
	return office, result.Error
}

func (repo *OfficeRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("company.offices").Count(&count)
	return count, result.Error
}

func (repo *OfficeRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.OfficeData, error) {
	var offices []entities.OfficeData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM company.view_office_data WHERE office_name LIKE ? OR email LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&offices)
	return offices, result.Error
}

func (repo *OfficeRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM company.view_office_data WHERE office_name LIKE ? OR email LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *OfficeRepository) ExistsByName(ctx context.Context, companyId int, officeName string) (bool, error) {
	var office entities.Office
	result := repo.db.WithContext(ctx).Where("company_id=? AND office_name=?", companyId, officeName).Find(&office)
	return office.OfficeId != 0, result.Error
}
