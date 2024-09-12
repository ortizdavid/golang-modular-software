package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type UserRole struct {
    UserRoleId int64 `gorm:"primaryKey;autoIncrement"`
    UserId     int64 `gorm:"column:user_id"`
    RoleId     int `gorm:"column:role_id"`
    shared.BaseEntity
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
	RoleStatus		string `json:"role_status"`
	UserId     		int64 `json:"user_id"`
	UserUniqueId	string `json:"user_unique_id"`
	UserName   		string `json:"user_name"`
}