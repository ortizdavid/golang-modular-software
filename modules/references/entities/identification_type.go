package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type IdentificationType struct {
	TypeId    int       `gorm:"autoIncrement;primaryKey"`
	TypeName  string    `gorm:"column:type_name"`
	Code      string    `gorm:"column:code"`
	shared.BaseEntity
}

func (IdentificationType) TableName() string {
	return "reference.identification_types"
}
