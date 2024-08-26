package entities

import (
    "time"
)

type Module struct {
    ModuleId    int       `gorm:"primaryKey;autoIncrement"`
    ModuleName  string    `gorm:"column:module_name"`
    Code        string    `gorm:"column:code"`
    Description string    `gorm:"column:description"`
    UniqueId    string `gorm:"column:unique_id"`
    CreatedAt   time.Time `gorm:"column:created_at"`
    UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Module) TableName() string {
    return "configurations.modules"
}

const (
    ModuleAuthentication = 1
    ModuleCompany = 2
    ModuleEmployees = 3
    ModuleReferences = 4
    ModuleReports = 5
    ModuleConfigurations = 6
)
