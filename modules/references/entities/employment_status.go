package entities

import "time"

type EmploymentStatus struct {
	StatusId    int       `gorm:"autoIncrement;primaryKey"`
	StatusName  string    `gorm:"column:status_name"`
	Code        string    `gorm:"column:code"`
	Description string    `gorm:"column:description"`
	UniqueId    string    `gorm:"column:unique_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (EmploymentStatus) TableName() string {
	return "reference.employment_statuses"
}
