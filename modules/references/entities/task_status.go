package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type TaskStatus struct {
	StatusId    int       `gorm:"autoIncrement;primaryKey"`
	StatusName  string    `gorm:"column:status_name"`
	Code        string    `gorm:"column:code"`
	LblColor    string    `gorm:"column:lbl_color"`
	BgColor     string    `gorm:"column:bg_color"`
	Description string    `gorm:"column:description"`
	shared.BaseEntity
}

func (TaskStatus) TableName() string {
	return "reference.task_statuses"
}
