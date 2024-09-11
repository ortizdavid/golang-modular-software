package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

type DocumentRepository struct {
	db *database.Database
}

func NewDocumentRepository(db *database.Database) *DocumentRepository {
	return &DocumentRepository{
		db: db,
	}
}

func (repo *DocumentRepository) Create(ctx context.Context, document entities.Document) error {
	result := repo.db.WithContext(ctx).Create(&document)
	return result.Error
}

func (repo *DocumentRepository) Update(ctx context.Context, document entities.Document) error {
	result := repo.db.WithContext(ctx).Save(&document)
	return result.Error
}

func (repo *DocumentRepository) Delete(ctx context.Context, document entities.Document) error {
	result := repo.db.WithContext(ctx).Delete(&document)
	return result.Error
}

func (repo *DocumentRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.DocumentData, error) {
	var documents []entities.DocumentData
	result := repo.db.WithContext(ctx).Table("employees.documents").Limit(limit).Offset(offset).Find(&documents)
	return documents, result.Error
}

func (repo *DocumentRepository) FindAllByEmployeeIdLimit(ctx context.Context, limit int, offset int, employeeId int64) ([]entities.DocumentData, error) {
	var documents []entities.DocumentData
	result := repo.db.WithContext(ctx).Table("employees.view_document_data").
		Where("employee_id=?", employeeId).
		Limit(limit).Offset(offset).
		Find(&documents)
	return documents, result.Error
}

func (repo *DocumentRepository) FindAllByEmployeeId(ctx context.Context, employeeId int64) ([]entities.DocumentData, error) {
	var documents []entities.DocumentData
	result := repo.db.WithContext(ctx).Table("employees.view_document_data").Where("employee_id=?", employeeId).Find(&documents)
	return documents, result.Error
}

func (repo *DocumentRepository) FindById(ctx context.Context, id int64) (entities.Document, error) {
	var document entities.Document
	result := repo.db.WithContext(ctx).First(&document, id)
	return document, result.Error
}

func (repo *DocumentRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.DocumentData, error) {
	var document entities.DocumentData
	result := repo.db.WithContext(ctx).Table("employees.view_document_data").Where("unique_id=?", uniqueId).First(&document)
	return document, result.Error
}

func (repo *DocumentRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Document, error) {
	var document entities.Document
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&document)
	return document, result.Error
}

func (repo *DocumentRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("employees.documents").Count(&count)
	return count, result.Error
}

func (repo *DocumentRepository) CountByEmployee(ctx context.Context, employeeId int64) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("employees.documents").Where("employee_id=?", employeeId).Count(&count)
	return count, result.Error
}

func (repo *DocumentRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.Document, error) {
	var documents []entities.Document
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM employees.documents WHERE document_name LIKE ?", likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&documents)
	return documents, result.Error
}

func (repo *DocumentRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM employees.documents WHERE document_name LIKE ?", likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *DocumentRepository) ExistsByName(ctx context.Context, documentName string, employeeId int64) (bool, error) {
	var document entities.Document
	result := repo.db.WithContext(ctx).Where("document_name=? AND employee_id=?", documentName, employeeId).Find(&document)
	return document.DocumentId != 0, result.Error
}
