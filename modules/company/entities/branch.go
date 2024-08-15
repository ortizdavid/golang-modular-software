package entities

import (
    "time"
)

type Branch struct {
    BranchID     uint           `gorm:"primaryKey;autoIncrement"`
    CompanyID    uint           `gorm:"column:company_id;not null"`
    BranchName   string         `gorm:"column:branch_name;not null"`
    Address      string         `gorm:"column:address"`
    Phone        string         `gorm:"column:phone"`
    Email        string         `gorm:"column:email"`
    UniqueID     string         `gorm:"column:unique_id"`
    CreatedAt    time.Time      `gorm:"column:created_at"`
    UpdatedAt    time.Time      `gorm:"column:updated_at"`
}

func (Branch) TableName() string {
    return "company.branches"
}

