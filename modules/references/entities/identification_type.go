package entities

import "time"

type IdentificationType struct {
	TypeId    int       `gorm:"autoIncrement;primaryKey"`
	TypeName  string    `gorm:"column:type_name"`
	Code      string    `gorm:"column:code"`
	UniqueId  string    `gorm:"column:unique_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (IdentificationType) TableName() string {
	return "reference.identification_types"
}
