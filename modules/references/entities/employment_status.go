package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type EmploymentStatus struct {
	StatusId    int       `gorm:"autoIncrement;primaryKey"`
	StatusName  string    `gorm:"column:status_name"`
	Code        string    `gorm:"column:code"`
	Description string    `gorm:"column:description"`
	shared.BaseEntity
}

func (EmploymentStatus) TableName() string {
	return "reference.employment_statuses"
}
