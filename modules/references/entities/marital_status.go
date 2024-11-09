package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type MaritalStatus struct {
	StatusId   int    `gorm:"autoIncrement;primaryKey" json:"status_id"`
	StatusName string `gorm:"column:status_name" json:"status_name"`
	Code       string `gorm:"column:code" json:"code"`
	shared.BaseEntity
}

func (MaritalStatus) TableName() string {
	return "reference.marital_statuses"
}
