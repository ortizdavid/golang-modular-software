package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
)

type DocumentStatusRepository struct {
	db *database.Database
}

func NewDocumentStatusRepository(db *database.Database) *DocumentStatusRepository {
	return &DocumentStatusRepository{
		db: db,
	}
}

func (repo *DocumentStatusRepository) Create(ctx context.Context, documentStatus entities.DocumentStatus) error {
	result := repo.db.WithContext(ctx).Create(&documentStatus)
	return result.Error
}

func (repo *DocumentStatusRepository) Update(ctx context.Context, documentStatus entities.DocumentStatus) error {
	result := repo.db.WithContext(ctx).Save(&documentStatus)
	return result.Error
}

func (repo *DocumentStatusRepository) Delete(ctx context.Context, documentStatus entities.DocumentStatus) error {
	result := repo.db.WithContext(ctx).Delete(&documentStatus)
	return result.Error
}

func (repo *DocumentStatusRepository) FindAll(ctx context.Context) ([]entities.DocumentStatus, error) {
	var documentStatuses []entities.DocumentStatus
	result := repo.db.WithContext(ctx).Find(&documentStatuses)
	return documentStatuses, result.Error
}

func (repo *DocumentStatusRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.DocumentStatus, error) {
	var documentStatuses []entities.DocumentStatus
	result := repo.db.WithContext(ctx).Table("reference.document_statuses").Limit(limit).Offset(offset).Find(&documentStatuses)
	return documentStatuses, result.Error
}

func (repo *DocumentStatusRepository) FindById(ctx context.Context, id int) (entities.DocumentStatus, error) {
	var documentStatus entities.DocumentStatus
	result := repo.db.WithContext(ctx).First(&documentStatus, id)
	return documentStatus, result.Error
}

func (repo *DocumentStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.DocumentStatus, error) {
	var documentStatus entities.DocumentStatus
	result := repo.db.WithContext(ctx).Table("reference.document_statuses").Where("unique_id=?", uniqueId).First(&documentStatus)
	return documentStatus, result.Error
}

func (repo *DocumentStatusRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.DocumentStatus, error) {
	var documentStatus entities.DocumentStatus
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&documentStatus)
	return documentStatus, result.Error
}

func (repo *DocumentStatusRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("reference.document_statuses").Count(&count)
	return count, result.Error
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
