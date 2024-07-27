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
	UserRoleId int64     `json:"user_role_id"`
	UniqueId   string    `json:"unique_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserId     int64     `json:"user_id"`
	UserName   string    `json:"user_name"`
}