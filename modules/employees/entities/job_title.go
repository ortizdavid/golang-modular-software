package entities

import "time"

type JobTitle struct {
	JobTitleId  int    `gorm:"autoIncrement;primaryKey"`
	TitleName   string `gorm:"column:title_name"`
	Description string `gorm:"column:description"`
	UniqueId    string `gorm:"column:unique_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt	time.Time `gorm:"column:updated_at"`
}

func (JobTitle) TableName() string {
	return "employees.job_titles"
}
