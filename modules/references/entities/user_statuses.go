package entities

import "time"

type UserStatus struct {
	StatusId          int       `gorm:"autoIncrement;primaryKey"`
	StatusName  string    `gorm:"column:status_name;unique"`
	Code        string    `gorm:"column:code;unique"`
	LblColor    string    `gorm:"column:lbl_color"`
	BgColor     string    `gorm:"column:bg_color"`
	Description string    `gorm:"column:description"`
	UniqueId    string    `gorm:"column:unique_id;unique;default:uuid_generate_v4()::text"`
	CreatedAt   time.Time `gorm:"column:created_at;default:now()"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:now()"`
}

func (UserStatus) TableName() string {
	return "reference.user_statuses"
}
