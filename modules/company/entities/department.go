package entities

import "time"

type Department struct {
	DepartmentId   int      `gorm:"primaryKey;autoIncrement"`
	CompanyId      int      `gorm:"column:company_id;not null"`
	DepartmentName string    `gorm:"column:department_name;not null"`
	Acronym        string    `gorm:"column:acronym"`
	Description    string    `gorm:"column:description"`
	UniqueId       string    `gorm:"column:unique_id"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (Department) TableName() string {
	return "company.departments"
}