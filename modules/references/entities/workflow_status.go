package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type WorkflowStatus struct {
	StatusId    int    `gorm:"autoIncrement;primaryKey" json:"status_id"`
	StatusName  string `gorm:"column:status_name" json:"status_name"`
	Code        string `gorm:"column:code" json:"code"`
	LblColor    string `gorm:"column:lbl_color" json:"lbl_color"`
	BgColor     string `gorm:"column:bg_color" json:"bg_color"`
	Description string `gorm:"column:description" json:"description"`
	shared.BaseEntity
}

func (WorkflowStatus) TableName() string {
	return "reference.workflow_statuses"
}
