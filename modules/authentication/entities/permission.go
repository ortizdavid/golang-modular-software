package entities

import "time"

type Permission struct {
	PermissionId    int `gorm:"primaryKey;autoIncrement"`
	PermissionName  string`gorm:"column:permission_name"`
	Code      		string `gorm:"column:code"`
	Description 	string `gorm:"column:description"`
	UniqueId  		string `gorm:"column:unique_id"`
	CreatedAt 		time.Time `gorm:"column:created_at"`
	UpdatedAt 		time.Time `gorm:"column:updated_at"`
}

func (Permission) TableName() string {
	return "authentication.permissions"
}

type PermissionData struct {
	PermissionId    int `json:"permission_id"`
	PermissionName  string `json:"permission_name"`
	Code      		string `json:"code"`
	Description 	string `json:"description"`
	CreatedAt 		string `json:"created_at"`
	UpdatedAt  		string `json:"updated_at"`
	RoleId 			int `json:"role_id"`
	RoleName 		string `json:"role_name"`
	RoleCode 		string `json:"role_code"`
	UserId			int `json:"user_id"`
	UserName 		string `json:"user_name"`
}

