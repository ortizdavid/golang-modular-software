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
	PermissionId    int
	PermissionName  string
	Code      		string
	Description 	string
	CreatedAt 		string
	UpdatedAt  		string
	RoleId 			int
	RoleName 		string
	RoleCode 		string
	UserId			int 
	UserName 		string
}
