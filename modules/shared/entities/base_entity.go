package entities

import "time"

type BaseEntity struct {
	UniqueId  string    `gorm:"column:unique_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

