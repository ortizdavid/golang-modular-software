package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type JobTitle struct {
	JobTitleId  int    `gorm:"autoIncrement;primaryKey"`
	TitleName   string `gorm:"column:title_name"`
	Description string `gorm:"column:description"`
	shared.BaseEntity
}

func (JobTitle) TableName() string {
	return "employees.job_titles"
}
