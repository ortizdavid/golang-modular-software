package entities

import "time"

type MaritalStatus struct {
	StatusId   int       `gorm:"autoIncrement;primaryKey;column:status_id"`
	StatusName string    `gorm:"column:status_name;unique"`
	Code       string    `gorm:"column:code;unique"`
	UniqueId   string    `gorm:"column:unique_id;unique;default:uuid_generate_v4()::text"`
	CreatedAt  time.Time `gorm:"column:created_at;default:now()"`
	UpdatedAt  time.Time `gorm:"column:updated_at;default:now()"`
}

func (MaritalStatus) TableName() string {
	return "reference.marital_statuses"
}
