package entities

import (
    "time"
)

type Branch struct {
    BranchId     int `gorm:"primaryKey;autoIncrement"`
    CompanyId    int `gorm:"column:company_id"`
    BranchName   string `gorm:"column:branch_name"`
    Address      string `gorm:"column:address"`
    Phone        string `gorm:"column:phone"`
    Email        string `gorm:"column:email"`
    UniqueId     string `gorm:"column:unique_id"`
    CreatedAt    time.Time `gorm:"column:created_at"`
    UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (Branch) TableName() string {
    return "company.branches"
}

