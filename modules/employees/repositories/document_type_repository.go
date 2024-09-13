package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type DocumentTypeRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.DocumentType]
}

func NewDocumentTypeRepository(db *database.Database) *DocumentTypeRepository {
	return &DocumentTypeRepository{
		db:             db,
		BaseRepository: shared.NewBaseRepository[entities.DocumentType](db),
	}
}

func (repo *DocumentTypeRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.DocumentType, error) {
	var documentType entities.DocumentType
	result := repo.db.WithContext(ctx).Table("employees.document_types").Where("unique_id=?", uniqueId).First(&documentType)
	return documentType, result.Error
}

func (repo *DocumentTypeRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.DocumentType, error) {
	var documentTypes []entities.DocumentType
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM employees.document_types WHERE type_name LIKE ?", likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&documentTypes)
	return documentTypes, result.Error
}

func (repo *DocumentTypeRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM employees.document_types WHERE type_name LIKE ?", likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *DocumentTypeRepository) ExistsByName(ctx context.Context, documentTypeName string) (bool, error) {
	var documentType entities.DocumentType
	result := repo.db.WithContext(ctx).Where("type_name=?", documentTypeName).Find(&documentType)
	return documentType.TypeId != 0, result.Error
}
