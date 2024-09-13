package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type DocumentStatusRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.DocumentStatus]
}

func NewDocumentStatusRepository(db *database.Database) *DocumentStatusRepository {
	return &DocumentStatusRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.DocumentStatus](db),
	}
}

func (repo *DocumentStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.DocumentStatus, error) {
	var documentStatus entities.DocumentStatus
	result := repo.db.WithContext(ctx).Table("reference.document_statuses").Where("unique_id=?", uniqueId).First(&documentStatus)
	return documentStatus, result.Error
}

func (repo *DocumentStatusRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.DocumentStatus, error) {
	var documentStatuses []entities.DocumentStatus
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.document_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&documentStatuses)
	return documentStatuses, result.Error
}

func (repo *DocumentStatusRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.document_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *DocumentStatusRepository) ExistsByName(ctx context.Context, documentStatusName string) (bool, error) {
	var documentStatus entities.DocumentStatus
	result := repo.db.WithContext(ctx).Where("status_name=?", documentStatusName).Find(&documentStatus)
	return documentStatus.StatusId != 0, result.Error
}
