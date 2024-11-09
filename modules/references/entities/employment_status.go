package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type EmploymentStatus struct {
	StatusId    int    `gorm:"autoIncrement;primaryKey" json:"status_id"`
	StatusName  string `gorm:"column:status_name" json:"status_name"`
	Code        string `gorm:"column:code" json:"code"`
	Description string `gorm:"column:description" json:"description"`
	shared.BaseEntity
}

func (EmploymentStatus) TableName() string {
	return "reference.employment_statuses"
}
