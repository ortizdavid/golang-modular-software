package entities

import (
	"time"

	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Project struct {
	ProjectId   int `gorm:"primaryKey;autoIncrement"`
	ProjectName string  `gorm:"column:project_name"`
	Description string  `gorm:"column:description"`
	StartDate   time.Time `gorm:"column:start_date"`
	EndDate     time.Time `gorm:"column:end_date"`
	Status      string  `gorm:"column:status"`
	CompanyId   int `gorm:"column:company_id"`
	shared.BaseEntity
}

func (Project) TableName() string {
	return "company.projects"
}

type ProjectData struct {
	ProjectId   int `json:"project_id"`
	UniqueId    string  `json:"unique_id"`
	ProjectName string  `json:"project_name"`
	Description string  `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Status      string  `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	CompanyId   int `json:"company_id"`
	CompanyName string  `json:"company_name"`
}