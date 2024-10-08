package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type EmployeeEmail struct {
	EmailId			int64 `gorm:"autoIncrement;primaryKey"`
	EmployeeId		int64 `gorm:"column:employee_id"`
	ContactTypeId	int	`gorm:"column:contact_type_id"`
	EmailAddress   	string `gorm:"column:email_address"`
	shared.BaseEntity
}

func (EmployeeEmail) TableName() string {
	return "employees.employee_emails"
}

type EmployeeEmailData struct {
	EmailId				int64 `json:"email_id"`
	EmailAddress 		string `json:"email_address"`
	UniqueId    		string `json:"unique_id"`
	CreatedAt   		string `json:"created_at"`
	UpdatedAt			string `json:"updated_at"`
	EmployeeId			int64 `json:"employee_id"`
	EmployeeUniqueId	string `json:"employee_unique_id"`
	FirstName			string `json:"first_name"`
	LastName			string `json:"last_name"`
	ContactTypeId		int	`json:"contact_type_id"`
	ContactTypeName		string `json:"contact_type_name"`
}

