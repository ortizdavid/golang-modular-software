package repositories

import (
	"context"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
)

type CompanyRepository struct {
	db *database.Database
}

func NewCompanyRepository(db *database.Database) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (repo *CompanyRepository) Create(ctx context.Context, company entities.Company) error {
	result := repo.db.WithContext(ctx).Create(&company)
	return result.Error
}

func (repo *CompanyRepository) Update(ctx context.Context, company entities.Company) error {
	result := repo.db.WithContext(ctx).Save(&company)
	return result.Error
}

func (repo *CompanyRepository) Delete(ctx context.Context, company entities.Company) error {
	result := repo.db.WithContext(ctx).Delete(&company)
	return result.Error
}

func (repo *CompanyRepository) FindAll(ctx context.Context) ([]entities.Company, error) {
	var companies []entities.Company
	result := repo.db.WithContext(ctx).Find(&companies)
	return companies, result.Error
}

func (repo *CompanyRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.CompanyData, error) {
	var companies []entities.CompanyData
	result := repo.db.WithContext(ctx).Table("company.view_company_data").Limit(limit).Offset(offset).Find(&companies)
	return companies, result.Error
}

func (repo *CompanyRepository) FindById(ctx context.Context, id int) (entities.Company, error) {
	var company entities.Company
	result := repo.db.WithContext(ctx).First(&company, id)
	return company, result.Error
}

func (repo *CompanyRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Company, error) {
	var company entities.Company
	result := repo.db.WithContext(ctx).Where("unqiue_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *CompanyRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.CompanyData, error) {
	var company entities.CompanyData
	result := repo.db.WithContext(ctx).Table("company.view_company_data").Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *CompanyRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("company.companies").Count(&count)
	return count, result.Error
}

func (repo *CompanyRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.CompanyData, error) {
	var companies []entities.CompanyData
	likeParam := "%"+param+"%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM company.view_company_data WHERE company_name LIKE ? OR company_acronym LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&companies)
	return companies, result.Error
}

func (repo *CompanyRepository) CountByParam(ctx context.Context, param string) (int64, error) {
    var count int64
	likeParam := "%"+param+"%"
    result := repo.db.WithContext(ctx).
        Raw("SELECT COUNT(*) FROM company.view_company_data WHERE company_name LIKE ? OR company_acronym LIKE ?", likeParam, likeParam).
        Scan(&count)
    return count, result.Error
}

func (repo *BranchRepository) ExistsByName(ctx context.Context, companyId int, branchName string) (bool, error) {
	var branch entities.Branch
	result := repo.db.WithContext(ctx).Where("company_id=? AND branch_name=?", companyId, branchName).Find(&branch)
	return branch.BranchId !=0 , result.Error
}