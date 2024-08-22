package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type CurrencyRepository struct {
	db *database.Database
}

func NewCurrencyRepository(db *database.Database) *CurrencyRepository {
	return &CurrencyRepository{
		db: db,
	}
}

func (repo *CurrencyRepository) Create(ctx context.Context, company entities.Currency) error {
	result := repo.db.WithContext(ctx).Create(&company)
	return result.Error
}

func (repo *CurrencyRepository) Update(ctx context.Context, company entities.Currency) error {
	result := repo.db.WithContext(ctx).Save(&company)
	return result.Error
}

func (repo *CurrencyRepository) Delete(ctx context.Context, company entities.Currency) error {
	result := repo.db.WithContext(ctx).Delete(&company)
	return result.Error
}

func (repo *CurrencyRepository) FindAll(ctx context.Context) ([]entities.Currency, error) {
	var currencies []entities.Currency
	result := repo.db.WithContext(ctx).Find(&currencies)
	return currencies, result.Error
}

func (repo *CurrencyRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.Currency, error) {
	var currencies []entities.Currency
	result := repo.db.WithContext(ctx).Table("reference.currencies").Limit(limit).Offset(offset).Find(&currencies)
	return currencies, result.Error
}

func (repo *CurrencyRepository) FindById(ctx context.Context, id int) (entities.Currency, error) {
	var company entities.Currency
	result := repo.db.WithContext(ctx).First(&company, id)
	return company, result.Error
}

func (repo *CurrencyRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.Currency, error) {
	var currency entities.Currency
	result := repo.db.WithContext(ctx).Table("reference.currencies").Where("unique_id=?", uniqueId).First(&currency)
	return currency, result.Error
}

func (repo *CurrencyRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Currency, error) {
	var company entities.Currency
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *CurrencyRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("reference.currencies").Count(&count)
	return count, result.Error
}

func (repo *CurrencyRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.Currency, error) {
	var currencies []entities.Currency
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.currencies WHERE currency_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&currencies)
	return currencies, result.Error
}

func (repo *CurrencyRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.currencies WHERE currency_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *CurrencyRepository) ExistsByName(ctx context.Context, currencyName string) (bool, error) {
	var currency entities.Currency
	result := repo.db.WithContext(ctx).Where("currency_name=?", currencyName).Find(&currency)
	return currency.CurrencyId != 0, result.Error
}
