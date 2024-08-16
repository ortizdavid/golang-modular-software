package entities

import (
    "time"
)

type Office struct {
    OfficeId    int `gorm:"primaryKey;autoIncrement"`
    BranchId    int `gorm:"column:branch_id"`
    OfficeName  string `gorm:"column:office_name"`
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
