package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Role struct {
	RoleId    	int `gorm:"primaryKey;autoIncrement"`
	RoleName  	string`gorm:"column:role_name"`
	Code      	string `gorm:"column:code"`
	Description string `gorm:"column:description"`
	Status      string `gorm:"column:status"`
	shared.BaseEntity
}

func (Role) TableName() string {
	return "authentication.roles"
}

type RoleData struct {
	RoleId      int `json:"role_id"`
	UniqueId    string `json:"unique_id"`
	RoleName    string `json:"role_name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Status		string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
