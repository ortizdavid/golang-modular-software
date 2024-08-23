package entities

import "time"

type EvaluationStatus struct {
	StatusId    int       `gorm:"autoIncrement;primaryKey"`
	StatusName  string    `gorm:"column:status_name"`
	Code        string    `gorm:"column:code"`
	Weight     	int    `gorm:"column:weight"`
	LblColor    string    `gorm:"column:lbl_color"`
	BgColor     string    `gorm:"column:bg_color"`
	Description string    `gorm:"column:description"`
	UniqueId    string    `gorm:"column:unique_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (EvaluationStatus) TableName() string {
	return "reference.evaluation_statuses"
}
