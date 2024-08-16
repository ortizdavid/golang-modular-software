package entities

import "time"

type UserRole struct {
    UserRoleId int64 `gorm:"primaryKey;autoIncrement"`
    UserId     int64 `gorm:"column:user_id"`
    RoleId     int `gorm:"column:role_id"`
    UniqueId   string `gorm:"column:unique_id"`
    CreatedAt  time.Time `gorm:"column:created_at"`
    UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (UserRole) TableName() string {
	return "authentication.user_roles"
}

type UserRoleData struct {
	UserRoleId 		int64 `json:"user_role_id"`
	UniqueId   		string `json:"unique_id"`
	CreatedAt  		string `json:"created_at"`
	UpdatedAt  		string `json:"updated_at"`
	RoleId     		int64 `json:"role_id"`
	RoleName   		string `json:"role_name"`
	RoleCode		string `json:"role_code"`
	UserId     		int64 `json:"user_id"`
	UserUniqueId	string `json:"user_unique_id"`
	UserName   		string `json:"user_name"`
}