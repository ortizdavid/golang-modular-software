package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type CountryRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Country]
}

func NewCountryRepository(db *database.Database) *CountryRepository {
	return &CountryRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Country](db),
	}
}

func (repo *CountryRepository) FindAllData(ctx context.Context) ([]entities.CountryData, error) {
	var countries []entities.CountryData
	result := repo.db.WithContext(ctx).Table("reference.view_country_data").Find(&countries)
	return countries, result.Error
}

func (repo *CountryRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.CountryData, error) {
	var countries []entities.CountryData
	result := repo.db.WithContext(ctx).Table("reference.view_country_data").Limit(limit).Offset(offset).Find(&countries)
	return countries, result.Error
}

func (repo *CountryRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.CountryData, error) {
	var country entities.CountryData
	result := repo.db.WithContext(ctx).Table("reference.view_country_data").Where("unique_id=?", uniqueId).First(&country)
	return country, result.Error
}

func (repo *CountryRepository) GetDataByName(ctx context.Context, name string) (entities.CountryData, error) {
	var country entities.CountryData
	result := repo.db.WithContext(ctx).Table("reference.view_country_data").Where("country_name=?", name).First(&country)
	return country, result.Error
}

func (repo *CountryRepository) GetDataByIsoCode(ctx context.Context, isoCode string) (entities.CountryData, error) {
	var country entities.CountryData
	result := repo.db.WithContext(ctx).Table("reference.view_country_data").Where("iso_code=? OR iso_code_lower=?", isoCode, isoCode).First(&country)
	return country, result.Error
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
