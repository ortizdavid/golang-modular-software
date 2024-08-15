package entities

import "time"

type Department struct {
	DepartmentID   uint      `gorm:"primaryKey;autoIncrement"`
	CompanyID      uint      `gorm:"column:company_id;not null"`
	DepartmentName string    `gorm:"column:department_name;not null"`
	Acronym        string    `gorm:"column:acronym"`
	Description    string    `gorm:"column:description"`
	UniqueID       string    `gorm:"column:unique_id"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (Department) TableName() string {
	return "company.departments"
}