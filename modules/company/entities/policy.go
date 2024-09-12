package entities

import (
	"time"

	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Policy struct {
    PolicyId        int `gorm:"primaryKey;autoIncrement"`
    CompanyId       int `gorm:"column:company_id"`
    PolicyName      string `gorm:"column:policy_name"`
    Description     string `gorm:"column:description"`
    EffectiveDate   time.Time `gorm:"column:effective_date"`
    shared.BaseEntity
}

func (Policy) TableName() string {
    return "company.policies"
}

type PolicyData struct {
    PolicyId        int `json:"policy_id"`
    UniqueId        string `json:"unique_id"`
    PolicyName      string `json:"policy_name"`
    Description     string `json:"description"`
    EffectiveDate   string `json:"effective_date"`
    CreatedAt       string `json:"created_at"`
    UpdatedAt       string `json:"updated_at"`
    CompanyId       int `json:"company_id"`
    CompanyName     string `json:"company_name"`
}
