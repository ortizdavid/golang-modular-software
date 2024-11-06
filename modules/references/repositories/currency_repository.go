package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type CurrencyRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Currency]
}

func NewCurrencyRepository(db *database.Database) *CurrencyRepository {
	return &CurrencyRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Currency](db),
	}
}

func (repo *CurrencyRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.Currency, error) {
	var currency entities.Currency
	result := repo.db.WithContext(ctx).Table("reference.currencies").Where("unique_id=?", uniqueId).First(&currency)
	return currency, result.Error
}

func (repo *CurrencyRepository) GetDataByName(ctx context.Context, name string) (entities.Currency, error) {
	var currency entities.Currency
	result := repo.db.WithContext(ctx).Table("reference.currencies").Where("currency_name=?", name).First(&currency)
	return currency, result.Error
}

func (repo *CurrencyRepository) GetDataByCode(ctx context.Context, code string) (entities.Currency, error) {
	var currency entities.Currency
	result := repo.db.WithContext(ctx).Table("reference.currencies").Where("code=?", code).First(&currency)
	return currency, result.Error
}

func (repo *CurrencyRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Currency, error) {
	var currency entities.Currency
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&currency)
	return currency, result.Error
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
