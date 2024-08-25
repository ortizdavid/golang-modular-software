package entities

import (
    "time"
)

type Feature struct {
    FeatureId   int       `gorm:"primaryKey;autoIncrement"`
    ModuleId    int       `gorm:"column:module_id"`
    FeatureName string    `gorm:"column:feature_name"`
    Code        string    `gorm:"column:code"`
    Description string    `gorm:"column:description"`
    UniqueId    string `gorm:"column:unique_id"`
    CreatedAt   time.Time `gorm:"column:created_at"`
    UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Feature) TableName() string {
    return "configurations.features"
}

type FeatureData struct {
    FeatureId   int       `json:"feature_id"`
    UniqueId    string `json:"unique_id"`
    FeatureName string    `json:"feature_name"`
    Code        string    `json:"code"`
    Description string    `json:"description"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
    ModuleId    int       `json:"module_id"`
    ModuleName string    `json:"module_name"`
}

