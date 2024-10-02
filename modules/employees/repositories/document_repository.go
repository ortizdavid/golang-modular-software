package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type DocumentRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Document]
}

func NewDocumentRepository(db *database.Database) *DocumentRepository {
	return &DocumentRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Document](db),
	}
}

func (repo *DocumentRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.DocumentData, error) {
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

func (repo *DocumentRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.DocumentData, error) {
	var document entities.DocumentData
	result := repo.db.WithContext(ctx).Table("employees.view_document_data").Where("unique_id=?", uniqueId).First(&document)
	return document, result.Error
}

func (repo *DocumentRepository) GetAllByEmployeeUniqueId(ctx context.Context, uniqueId string) ([]entities.DocumentData, error) {
	var documents []entities.DocumentData
	result := repo.db.WithContext(ctx).Table("employees.view_document_data").Where("employee_unique_id=?", uniqueId).Find(&documents)
	return documents, result.Error
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
