package entities

import "time"

type Document struct {
	DocumentId		int64 `gorm:"autoIncrement;primaryKey"`
	EmployeeId		int64 `gorm:"column:employee_id"`
	DocumentTypeId	int	`gorm:"column:document_type_id"`
	DocumentName   	string `gorm:"column:document_name"`
	DocumentNumber 	string `gorm:"column:document_number"`
	ExpirationDate	time.Time `gorm:"column:expiration_date"`
	FileName		string `gorm:"column:file_name"`
	Status			string `gorm:"column:status"`
	UniqueId    	string `gorm:"column:unique_id"`
	CreatedAt   	time.Time `gorm:"column:created_at"`
	UpdatedAt		time.Time `gorm:"column:updated_at"`
}

func (Document) TableName() string {
	return "employees.documents"
}

type DocumentData struct {
	DocumentId			int64 `json:"document_id"`
	DocumentName   		string `json:"document_name"`
	DocumentNumber 		string `json:"document_number"`
	ExpirationDate		string `json:"expiration_date"`
	FileName			string `json:"file_name"`
	Status				string `json:"status"`
	UniqueId    		string `json:"unique_id"`
	CreatedAt   		string `json:"created_at"`
	UpdatedAt			string `json:"updated_at"`
	EmployeeId			int64 `json:"employee_id"`
	FirstName			string `json:"first_name"`
	LastName			string `json:"last_name"`
	DocumentTypeId		int	`json:"document_type_id"`
	DocumentTypeName	string `json:"document_type_name"`
}