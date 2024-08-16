package entities

import "time"

type Project struct {
	ProjectID   int `gorm:"primaryKey;autoIncrement"`
	ProjectName string  `gorm:"column:project_name"`
	Description string  `gorm:"column:description"`
	StartDate   time.Time `gorm:"column:start_date"`
	EndDate     time.Time `gorm:"column:end_date"`
	Status      string  `gorm:"column:status"`
	CompanyID   int `gorm:"column:company_id"`
	UniqueID    string  `gorm:"column:unique_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Project) TableName() string {
	return "company.projects"
}