package entities

import "time"

type ContactType struct {
	TypeId    int       `gorm:"autoIncrement;primaryKey"`
	TypeName  string    `gorm:"column:type_name"`
	Code      string    `gorm:"column:code"`
	UniqueId  string    `gorm:"column:unique_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (ContactType) TableName() string {
	return "reference.contact_types"
}
