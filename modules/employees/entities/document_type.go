package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type DocumentType struct {
	TypeId  int `gorm:"autoIncrement;primaryKey"`
	TypeName   string `gorm:"column:type_name"`
	Description string `gorm:"column:description"`
	shared.BaseEntity
}

func (DocumentType) TableName() string {
	return "employees.document_types"
}
