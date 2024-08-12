package entities

import "time"

type PermissionRole struct {
    PermissionRoleId 	int64 `gorm:"primaryKey;autoIncrement"`
    PermissionId     	int64 `gorm:"column:permission_id"`
    RoleId     			int `gorm:"column:role_id"`
    UniqueId   			string `gorm:"column:unique_id"`
    CreatedAt  			time.Time `gorm:"column:created_at"`
    UpdatedAt  			time.Time `gorm:"column:updated_at"`
}

func (PermissionRole) TableName() string {
	return "authentication.permission_roles"
}

type PermissionRoleData struct {
	PermissionRoleId 	int  `json:"permission_role_id"`
	UniqueId   			string `json:"unique_id"`
	Code  				string `json:"code_at"`
	CreatedAt  			string `json:"created_at"`
	UpdateAt  			string `json:"updated_at"`
	PermissionId     	int  `json:"permission_id"`
	PermissionName   	string `json:"permission_name"`
	PermissionCode     	string  `json:"permission_code"`
	RoleId				int `json:"role_id"`
	RoleUniqueId 		string `json:"role_unique_id"`
	RoleName			string `json:"role_name"`
}
