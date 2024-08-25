package entities

import (
    "time"
)

type Module struct {
    ModuleId    int       `gorm:"primaryKey;autoIncrement"`
    ModuleName  string    `gorm:"column:module_name"`
    Description string    `gorm:"column:description"`
    UniqueId    string `gorm:"column:unique_id"`
    CreatedAt   time.Time `gorm:"column:created_at"`
    UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Module) TableName() string {
    return "configurations.modules"
}
