package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

type DocumentTypeRepository struct {
	db *database.Database
}

func NewDocumentTypeRepository(db *database.Database) *DocumentTypeRepository {
	return &DocumentTypeRepository{
		db: db,
	}
}

func (repo *DocumentTypeRepository) Create(ctx context.Context, documentType entities.DocumentType) error {
	result := repo.db.WithContext(ctx).Create(&documentType)
	return result.Error
}

func (repo *DocumentTypeRepository) Update(ctx context.Context, documentType entities.DocumentType) error {
	result := repo.db.WithContext(ctx).Save(&documentType)
	return result.Error
}

func (repo *DocumentTypeRepository) Delete(ctx context.Context, documentType entities.DocumentType) error {
	result := repo.db.WithContext(ctx).Delete(&documentType)
	return result.Error
}

func (repo *DocumentTypeRepository) FindAll(ctx context.Context) ([]entities.DocumentType, error) {
	var documentTypes []entities.DocumentType
	result := repo.db.WithContext(ctx).Find(&documentTypes)
	return documentTypes, result.Error
}

func (repo *DocumentTypeRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.DocumentType, error) {
	var documentTypes []entities.DocumentType
	result := repo.db.WithContext(ctx).Table("employees.document_types").Limit(limit).Offset(offset).Find(&documentTypes)
	return documentTypes, result.Error
}

func (repo *DocumentTypeRepository) FindById(ctx context.Context, id int) (entities.DocumentType, error) {
	var documentType entities.DocumentType
	result := repo.db.WithContext(ctx).First(&documentType, id)
	return documentType, result.Error
}

func (repo *DocumentTypeRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.DocumentType, error) {
	var documentType entities.DocumentType
	result := repo.db.WithContext(ctx).Table("employees.document_types").Where("unique_id=?", uniqueId).First(&documentType)
	return documentType, result.Error
}

func (repo *DocumentTypeRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.DocumentType, error) {
	var documentType entities.DocumentType
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&documentType)
	return documentType, result.Error
}

func (repo *DocumentTypeRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("employees.document_types").Count(&count)
	return count, result.Error
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
