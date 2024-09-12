package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Module struct {
	ModuleId    int       `gorm:"primaryKey;autoIncrement"`
	ModuleName  string    `gorm:"column:module_name"`
	Code        string    `gorm:"column:code"`
	Description string    `gorm:"column:description"`
	shared.BaseEntity
}

func (Module) TableName() string {
	return "configurations.modules"
}
