package entities

import "time"

type Policy struct {
    PolicyID     uint           `gorm:"primaryKey;autoIncrement"`
    PolicyName   string         `gorm:"column:policy_name;not null"`
    Description  string         `gorm:"column:description"`
    EffectiveDate *time.Time    `gorm:"column:effective_date"`
    CompanyID    uint           `gorm:"column:company_id"`
    UniqueID     string         `gorm:"column:unique_id"`
    CreatedAt    time.Time      `gorm:"column:created_at"`
    UpdatedAt    time.Time      `gorm:"column:updated_at"`
}

func (Policy) TableName() string {
    return "company.policies"
}
