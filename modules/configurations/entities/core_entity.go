package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type CoreEntity struct {
	EntityId    int       `gorm:"primaryKey;autoIncrement"`
	ModuleId    int       `gorm:"column:module_id"`
	EntityName  string    `gorm:"column:entity_name"`
	Code        string    `gorm:"column:code"`
	Description string    `gorm:"column:description"`
	shared.BaseEntity
}

func (CoreEntity) TableName() string {
	return "configurations.core_entities"
}

type CoreEntityData struct {
	EntityId    int       `json:"entity_id"`
	UniqueId    string    `json:"unique_id"`
	EntityName  string    `json:"entity_name"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	ModuleId    int       `json:"module_id"`
	ModuleName  string      `json:"module_name"`
}
