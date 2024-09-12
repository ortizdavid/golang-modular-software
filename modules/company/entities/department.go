package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Department struct {
	DepartmentId   int `gorm:"primaryKey;autoIncrement"`
	CompanyId      int `gorm:"column:company_id"`
	DepartmentName string `gorm:"column:department_name"`
	Acronym        string `gorm:"column:acronym"`
	Description    string `gorm:"column:description"`
	shared.BaseEntity
}

func (Department) TableName() string {
	return "company.departments"
}

type DepartmentData struct {
	DepartmentId   int `json:"department_id"`
	DepartmentName string `json:"department_name"`
	Acronym        string `json:"acronym"`
	Description    string `json:"description"`
	UniqueId       string `json:"unique_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	CompanyId      int `json:"company_id"`
	CompanyName 	string `json:"company_name"`
}
