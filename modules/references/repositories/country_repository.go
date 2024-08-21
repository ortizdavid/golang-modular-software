package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type CountryRepository struct {
	db *database.Database
}

func NewCountryRepository(db *database.Database) *CountryRepository {
	return &CountryRepository{
		db: db,
	}
}

func (repo *CountryRepository) Create(ctx context.Context, company entities.Country) error {
	result := repo.db.WithContext(ctx).Create(&company)
	return result.Error
}

func (repo *CountryRepository) Update(ctx context.Context, company entities.Country) error {
	result := repo.db.WithContext(ctx).Save(&company)
	return result.Error
}

func (repo *CountryRepository) Delete(ctx context.Context, company entities.Country) error {
	result := repo.db.WithContext(ctx).Delete(&company)
	return result.Error
}

func (repo *CountryRepository) FindAll(ctx context.Context) ([]entities.CountryData, error) {
	var countries []entities.CountryData
	result := repo.db.WithContext(ctx).Find(&countries)
	return countries, result.Error
}

func (repo *CountryRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.CountryData, error) {
	var countries []entities.CountryData
	result := repo.db.WithContext(ctx).Table("reference.view_country_data").Limit(limit).Offset(offset).Find(&countries)
	return countries, result.Error
}

func (repo *CountryRepository) FindById(ctx context.Context, id int) (entities.Country, error) {
	var company entities.Country
	result := repo.db.WithContext(ctx).First(&company, id)
	return company, result.Error
}

func (repo *CountryRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.CountryData, error) {
	var country entities.CountryData
	result := repo.db.WithContext(ctx).Table("reference.view_country_data").Where("unique_id=?", uniqueId).First(&country)
	return country, result.Error
}

func (repo *CountryRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Country, error) {
	var company entities.Country
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *CountryRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("reference.countries").Count(&count)
	return count, result.Error
}

func (repo *CountryRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.CountryData, error) {
	var countries []entities.CountryData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.view_country_data WHERE country_name LIKE ? OR iso_code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&countries)
	return countries, result.Error
}

func (repo *CountryRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.view_country_data WHERE country_name LIKE ? OR iso_code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *CountryRepository) ExistsByName(ctx context.Context, countryName string) (bool, error) {
	var country entities.Country
	result := repo.db.WithContext(ctx).Where("country_name=?", countryName).Find(&country)
	return country.CountryId != 0, result.Error
}
