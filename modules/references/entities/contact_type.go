package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type ContactType struct {
	TypeId    int       `gorm:"autoIncrement;primaryKey"`
	TypeName  string    `gorm:"column:type_name"`
	Code      string    `gorm:"column:code"`
	shared.BaseEntity
}

func (ContactType) TableName() string {
	return "reference.contact_types"
}
