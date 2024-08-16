package entities

import "time"

type Policy struct {
    PolicyId     int `gorm:"primaryKey;autoIncrement"`
    PolicyName   string `gorm:"column:policy_name"`
    Description  string `gorm:"column:description"`
    EffectiveDate *time.Time `gorm:"column:effective_date"`
    CompanyId    int `gorm:"column:company_id"`
    UniqueId     string `gorm:"column:unique_id"`
    CreatedAt    time.Time `gorm:"column:created_at"`
    UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (Policy) TableName() string {
    return "company.policies"
}
