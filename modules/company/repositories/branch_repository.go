package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
)

type BranchRepository struct {
	db *database.Database
}

func NewBranchRepository(db *database.Database) *BranchRepository {
	return &BranchRepository{
		db: db,
	}
}

func (repo *BranchRepository) Create(ctx context.Context, company entities.Branch) error {
	result := repo.db.WithContext(ctx).Create(&company)
	return result.Error
}

func (repo *BranchRepository) Update(ctx context.Context, company entities.Branch) error {
	result := repo.db.WithContext(ctx).Save(&company)
	return result.Error
}

func (repo *BranchRepository) Delete(ctx context.Context, company entities.Branch) error {
	result := repo.db.WithContext(ctx).Delete(&company)
	return result.Error
}

func (repo *BranchRepository) FindAll(ctx context.Context) ([]entities.Branch, error) {
	var branches []entities.Branch
	result := repo.db.WithContext(ctx).Find(&branches)
	return branches, result.Error
}

func (repo *BranchRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.BranchData, error) {
	var branches []entities.BranchData
	result := repo.db.WithContext(ctx).Table("company.view_branch_data").Limit(limit).Offset(offset).Find(&branches)
	return branches, result.Error
}

func (repo *BranchRepository) FindById(ctx context.Context, id int) (entities.Branch, error) {
	var company entities.Branch
	result := repo.db.WithContext(ctx).First(&company, id)
	return company, result.Error
}

func (repo *BranchRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Branch, error) {
	var company entities.Branch
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *BranchRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.BranchData, error) {
	var company entities.BranchData
	result := repo.db.WithContext(ctx).Table("company.view_branch_data").Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *BranchRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("company.branches").Count(&count)
	return count, result.Error
}

func (repo *BranchRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.BranchData, error) {
	var branches []entities.BranchData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM company.view_branch_data WHERE branch_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&branches)
	return branches, result.Error
}

func (repo *BranchRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM company.view_branch_data WHERE branch_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *BranchRepository) ExistsByName(ctx context.Context, companyId int, branchName string) (bool, error) {
	var branch entities.Branch
	result := repo.db.WithContext(ctx).Where("company_id=? AND branch_name=?", companyId, branchName).Find(&branch)
	return branch.BranchId != 0, result.Error
}
