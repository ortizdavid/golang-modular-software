package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type ApprovalStatusRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.ApprovalStatus]
}

func NewApprovalStatusRepository(db *database.Database) *ApprovalStatusRepository {
	return &ApprovalStatusRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.ApprovalStatus](db),
	}
}

func (repo *ApprovalStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.ApprovalStatus, error) {
	var approvalStatus entities.ApprovalStatus
	result := repo.db.WithContext(ctx).Table("reference.approval_statuses").Where("unique_id=?", uniqueId).First(&approvalStatus)
	return approvalStatus, result.Error
}

func (repo *ApprovalStatusRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.ApprovalStatus, error) {
	var approvalStatus entities.ApprovalStatus
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&approvalStatus)
	return approvalStatus, result.Error
}

func (repo *ApprovalStatusRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.ApprovalStatus, error) {
	var approval_statuses []entities.ApprovalStatus
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.approval_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&approval_statuses)
	return approval_statuses, result.Error
}

func (repo *ApprovalStatusRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.approval_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *ApprovalStatusRepository) ExistsByName(ctx context.Context, approvalStatusName string) (bool, error) {
	var approvalStatus entities.ApprovalStatus
	result := repo.db.WithContext(ctx).Where("status_name=?", approvalStatusName).Find(&approvalStatus)
	return approvalStatus.StatusId != 0, result.Error
}
