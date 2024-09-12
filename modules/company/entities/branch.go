package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Branch struct {
    BranchId     int `gorm:"primaryKey;autoIncrement"`
    CompanyId    int `gorm:"column:company_id"`
    BranchName   string `gorm:"column:branch_name"`
    Code         string `gorm:"column:code"`
    Address      string `gorm:"column:address"`
    Phone        string `gorm:"column:phone"`
    Email        string `gorm:"column:email"`
    shared.BaseEntity
}

func (Branch) TableName() string {
    return "company.branches"
}

type BranchData struct {
    BranchId     int `json:"branch_id"`
    UniqueId     string `json:"unique_id"`
    BranchName   string `json:"branch_name"`
    Code         string `json:"code"`
    Address      string `json:"address"`
    Phone        string `json:"phone"`
    Email        string `json:"email"`
    CreatedAt    string `json:"created_at"`
    UpdatedAt    string `json:"updated_at"` 
    CompanyId    int `json:"company_id"`
    CompanyName  string `json:"company_name"`
}