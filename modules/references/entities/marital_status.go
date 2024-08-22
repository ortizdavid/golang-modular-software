package entities

import "time"

type MaritalStatus struct {
	StatusId   int       `gorm:"autoIncrement;primaryKey"`
	StatusName string    `gorm:"column:status_name"`
	Code       string    `gorm:"column:code"`
	UniqueId   string    `gorm:"column:unique_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (MaritalStatus) TableName() string {
	return "reference.marital_statuses"
}
