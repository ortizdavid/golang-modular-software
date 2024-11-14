package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type BranchRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Branch]
}

func NewBranchRepository(db *database.Database) *BranchRepository {
	return &BranchRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Branch](db),
	}
}

func (repo *BranchRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.BranchData, error) {
	var branches []entities.BranchData
	result := repo.db.WithContext(ctx).Table("company.view_branch_data").Limit(limit).Offset(offset).Find(&branches)
	return branches, result.Error
}

func (repo *BranchRepository) FindAllData(ctx context.Context) ([]entities.BranchData, error) {
	var branches []entities.BranchData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM company.view_branch_data").Find(&branches)
	return branches, result.Error
}

func (repo *BranchRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.BranchData, error) {
	var branch entities.BranchData
	result := repo.db.WithContext(ctx).Table("company.view_branch_data").Where("unique_id=?", uniqueId).First(&branch)
	return branch, result.Error
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
