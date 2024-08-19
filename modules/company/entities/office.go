package entities

import (
    "time"
)

type Office struct {
    OfficeId    int `gorm:"primaryKey;autoIncrement"`
    CompanyId    int `gorm:"column:company_id"`
    OfficeName  string `gorm:"column:office_name"`
    Code        string `gorm:"column:code"`
    Address     string `gorm:"column:address"`
    Phone       string `gorm:"column:phone"`
    Email       string `gorm:"column:email"`
    UniqueId    string `gorm:"column:unique_id"`
    CreatedAt   time.Time `gorm:"column:created_at"`
    UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Office) TableName() string {
    return "company.offices"
}

type OfficeData struct {
    OfficeId     int `json:"office_id"`
    UniqueId     string `json:"unique_id"`
    OfficeName   string `json:"office_name"`
    Code         string `json:"code"`
    Address      string `json:"address"`
    Phone        string `json:"phone"`
    Email        string `json:"email"`
    CreatedAt    string `json:"created_at"`
    UpdatedAt    string `json:"updated_at"` 
    CompanyId    int `json:"company_id"`
    CompanyName  string `json:"company_name"`
}