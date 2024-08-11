package entities

import "time"

type PermissionRole struct {
    PermissionRoleId int64 `gorm:"primaryKey;autoIncrement"`
    UserId     int64 `gorm:"column:permission_id"`
    RoleId     int `gorm:"column:role_id"`
    UniqueId   string `gorm:"column:unique_id"`
    CreatedAt  time.Time `gorm:"column:created_at"`
    UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (PermissionRole) TableName() string {
	return "authentication.permission_roles"
}

type PermissionRoleData struct {
	PermissionRoleId int64  `json:"permission_role_id"`
	UniqueId   string `json:"unique_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserId     int64  `json:"permission_id"`
	UserName   string `json:"permission_name"`
}