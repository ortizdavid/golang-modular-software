package entities

import (
    "time"
)

type FeatureFlag struct {
    FlagId      int       `gorm:"primaryKey;autoIncrement"`
    FeatureId   int       `gorm:"column:feature_id"`
    Status      string    `gorm:"column:status"`
    UniqueId    string `gorm:"column:unique_id"`
    CreatedAt   time.Time `gorm:"column:created_at"`
    UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (FeatureFlag) TableName() string {
    return "configurations.feature_flag"
}

type FeatureFlagData struct {
    FlagId      int       `json:"flag_id"`
    UniqueId    string `json:"unique_id"`
    Status      string    `json:"status"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
    FeatureId   int       `json:"feature_id"`
    FeatureName   string       `json:"feature_name"`
}
