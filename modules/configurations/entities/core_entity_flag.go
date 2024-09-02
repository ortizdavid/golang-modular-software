package entities

import (
    "time"
)

type CoreEntityFlag struct {
    FlagId    int       `gorm:"primaryKey;autoIncrement"`
    EntityId  int       `gorm:"column:entity_id"`
    ModuleId  int       `gorm:"column:module_id"`
    Status    string    `gorm:"column:status"` 
    UniqueId    string `gorm:"column:unique_id"`
    CreatedAt time.Time `gorm:"column:created_at"`
    UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (CoreEntityFlag) TableName() string {
    return "configurations.core_entity_flag"
}

type CoreEntityFlagData struct {
    FlagId      int `json:"flag_id"`
    UniqueId    string `json:"unique_id"`
    Status      string `json:"status"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
    EntityId    int `json:"entity_id"`
    EntityName  string `json:"entity_name"`
    Code        string `json:"code"`
    ModuleId    int `json:"module_id"`
    ModuleName  string `json:"module_name"`
    ModuleCode  string `json:"module_code"`
}
