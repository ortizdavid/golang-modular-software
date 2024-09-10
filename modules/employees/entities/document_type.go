package entities

import "time"

type DocumentType struct {
	TypeId  int    `gorm:"autoIncrement;primaryKey"`
	TypeName   string `gorm:"column:type_name"`
	Description string `gorm:"column:description"`
	UniqueId    string `gorm:"column:unique_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt	time.Time `gorm:"column:updated_at"`
}

func (DocumentType) TableName() string {
	return "employees.document_types"
}
