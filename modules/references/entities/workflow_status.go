package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type WorkflowStatus struct {
	StatusId    int       `gorm:"autoIncrement;primaryKey"`
	StatusName  string    `gorm:"column:status_name"`
	Code        string    `gorm:"column:code"`
	LblColor    string    `gorm:"column:lbl_color"`
	BgColor     string    `gorm:"column:bg_color"`
	Description string    `gorm:"column:description"`
	shared.BaseEntity
}

func (WorkflowStatus) TableName() string {
	return "reference.workflow_statuses"
}
