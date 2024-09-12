package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Permission struct {
	PermissionId    int `gorm:"primaryKey;autoIncrement"`
	PermissionName  string`gorm:"column:permission_name"`
	Code      		string `gorm:"column:code"`
	Description 	string `gorm:"column:description"`
	shared.BaseEntity
}

func (Permission) TableName() string {
	return "authentication.permissions"
}

type PermissionData struct {
	PermissionId    int `json:"permission_id"`
	UniqueId		string `json:"unique_id"`
	PermissionName  string `json:"permission_name"`
	Code      		string `json:"code"`
	Description 	string `json:"description"`
	CreatedAt 		string `json:"created_at"`
	UpdatedAt  		string `json:"updated_at"`
}
