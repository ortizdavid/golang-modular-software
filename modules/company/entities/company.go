package entities

import (
    "time"
)

type Company struct {
    CompanyId    uint           `gorm:"primaryKey;autoIncrement"`
    CompanyName  string         `gorm:"column:company_name"`
    CompanyAcronym string       `gorm:"column:company_acronym"`
    CompanyType  string         `gorm:"column:company_type"`
    Industry     string         `gorm:"column:industry"`
    FoundedDate  time.Time     `gorm:"column:founded_date"`
    Address      string         `gorm:"column:address"`
    Phone        string         `gorm:"column:phone"`
    Email        string         `gorm:"column:email"`
    WebsiteURL   string         `gorm:"column:website_url"`
    UniqueID     string         `gorm:"column:unique_id"`
    CreatedAt    time.Time      `gorm:"column:created_at"`
    UpdatedAt    time.Time      `gorm:"column:updated_at"`
}

func (Company) TableName() string {
    return "company.companies"
}
