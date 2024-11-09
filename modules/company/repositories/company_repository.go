package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type CompanyRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Company]
}

func NewCompanyRepository(db *database.Database) *CompanyRepository {
	return &CompanyRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Company](db),
	}
}

func (repo *CompanyRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.CompanyData, error) {
	var companies []entities.CompanyData
	result := repo.db.WithContext(ctx).Table("company.view_company_data").Limit(limit).Offset(offset).Find(&companies)
	return companies, result.Error
}

func (repo *CompanyRepository) GetCurrentData(ctx context.Context) (entities.CompanyData, error) {
	var company entities.CompanyData
	result := repo.db.WithContext(ctx).Table("company.view_company_data").First(&company)
	return company, result.Error
}

func (repo *CompanyRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.CompanyData, error) {
	var company entities.CompanyData
	result := repo.db.WithContext(ctx).Table("company.view_company_data").Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *CompanyRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.CompanyData, error) {
	var companies []entities.CompanyData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM company.view_company_data WHERE company_name LIKE ? OR company_acronym LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&companies)
	return companies, result.Error
}

func (repo *CompanyRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM company.view_company_data WHERE company_name LIKE ? OR company_acronym LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}
