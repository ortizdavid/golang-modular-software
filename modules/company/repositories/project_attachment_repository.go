package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type ProjectAttachmentRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.ProjectAttachment]
}

func NewProjectAttachmentRepository(db *database.Database) *ProjectAttachmentRepository {
	return &ProjectAttachmentRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.ProjectAttachment](db),
	}
}

func (repo *ProjectAttachmentRepository) FindAllByProjectId(ctx context.Context, projectId int) ([]entities.ProjectAttachment, error) {
	var attachments []entities.ProjectAttachment
	result := repo.db.WithContext(ctx).Where("project_id = ?", projectId).Find(&attachments)
	return attachments, result.Error
}


