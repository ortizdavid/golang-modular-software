package entities

import (
	"time"

	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Company struct {
    CompanyId    int `gorm:"primaryKey;autoIncrement"`
    CompanyName  string `gorm:"column:company_name"`
    CompanyAcronym string `gorm:"column:company_acronym"`
    CompanyType  string `gorm:"column:company_type"`
    Industry     string `gorm:"column:industry"`
    FoundedDate  time.Time `gorm:"column:founded_date"`
    Address      string `gorm:"column:address"`
    Phone        string `gorm:"column:phone"`
    Email        string `gorm:"column:email"`
    WebsiteURL   string `gorm:"column:website_url"`
    shared.BaseEntity
}

func (Company) TableName() string {
    return "company.companies"
}

type CompanyData struct {
    CompanyId    int `gorm:"primaryKey;autoIncrement"`
    UniqueId     string `json:"unique_id"`
    CompanyName  string `json:"company_name"`
    CompanyAcronym string `json:"company_acronym"`
    CompanyType  string `json:"company_type"`
    Industry     string `json:"industry"`
    FoundedDate  string `json:"founded_date"`
    Address      string `json:"address"`
    Phone        string `json:"phone"`
    Email        string `json:"email"`
    WebsiteURL   string `json:"website_url"`
    CreatedAt    string `json:"created_at"`
    UpdatedAt    string `json:"updated_at"`
}