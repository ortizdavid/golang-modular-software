package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type PolicyRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Policy]
}

func NewPolicyRepository(db *database.Database) *PolicyRepository {
	return &PolicyRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Policy](db),
	}
}

func (repo *PolicyRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.PolicyData, error) {
	var policies []entities.PolicyData
	result := repo.db.WithContext(ctx).Table("company.view_policy_data").Limit(limit).Offset(offset).Find(&policies)
	return policies, result.Error
}

func (repo *PolicyRepository) FindAllData(ctx context.Context) ([]entities.PolicyData, error) {
	var policies []entities.PolicyData
	result := repo.db.WithContext(ctx).Table("company.view_policy_data").Find(&policies)
	return policies, result.Error
}

func (repo *PolicyRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.PolicyData, error) {
	var policy entities.PolicyData
	result := repo.db.WithContext(ctx).Table("company.view_policy_data").Where("unique_id=?", uniqueId).First(&policy)
	return policy, result.Error
}

func (repo *PolicyRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.PolicyData, error) {
	var policies []entities.PolicyData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM company.view_policy_data WHERE policy_name LIKE ? OR company_name LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&policies)
	return policies, result.Error
}

func (repo *PolicyRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM company.view_policy_data WHERE policy_name LIKE ? OR company_name LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *PolicyRepository) ExistsByName(ctx context.Context, companyId int, policyName string) (bool, error) {
	var policy entities.Policy
	result := repo.db.WithContext(ctx).Where("company_id=? AND policy_name=?", companyId, policyName).Find(&policy)
	return policy.PolicyId != 0, result.Error
}
