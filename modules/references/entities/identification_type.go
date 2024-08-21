package entities

import "time"

type IdentificationType struct {
	TypeId    int       `gorm:"autoIncrement;primaryKey;column:type_id"`
	TypeName  string    `gorm:"column:type_name;unique"`
	Code      string    `gorm:"column:code;unique"`
	UniqueId  string    `gorm:"column:unique_id;unique;default:uuid_generate_v4()::text"`
	CreatedAt time.Time `gorm:"column:created_at;default:now()"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:now()"`
}

func (IdentificationType) TableName() string {
	return "reference.identification_types"
}
