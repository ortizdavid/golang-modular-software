package entities

import (
    "time"
)

type ModuleFlag struct {
    FlagId    int       `gorm:"primaryKey;autoIncrement"`
    ModuleId  int       `gorm:"column:module_id"`
    Status    string    `gorm:"column:status"` 
    UniqueId    string `gorm:"column:unique_id"`
    CreatedAt time.Time `gorm:"column:created_at"`
    UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (ModuleFlag) TableName() string {
    return "configurations.module_flag"
}

type ModuleFlagData struct {
    FlagId      int `json:"flag_id"`
    UniqueId    string `json:"unique_id"`
    Status      string `json:"status"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
    ModuleId    int `json:"module_id"`
    ModuleName   string `json:"module_name"`
    Code         string `json:"code"`
}
