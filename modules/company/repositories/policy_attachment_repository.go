package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type PolicyAttachmentRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.PolicyAttachment]
}

func NewPolicyAttachmentRepository(db *database.Database) *PolicyAttachmentRepository {
	return &PolicyAttachmentRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.PolicyAttachment](db),
	}
}

func (repo *PolicyAttachmentRepository) FindAllByPolicyId(ctx context.Context, policyId int) ([]entities.PolicyAttachment, error) {
	var policies []entities.PolicyAttachment
	result := repo.db.WithContext(ctx).Where("policy_id = ?", policyId).Find(&policies)
	return policies, result.Error
}


