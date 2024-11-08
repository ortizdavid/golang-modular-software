package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type ContactType struct {
	TypeId    int       `gorm:"autoIncrement;primaryKey" json:"type_id"`
	TypeName  string    `gorm:"column:type_name" json:"type_name"`
	Code      string    `gorm:"column:code" json:"code"`
	shared.BaseEntity
}

func (ContactType) TableName() string {
	return "reference.contact_types"
}
