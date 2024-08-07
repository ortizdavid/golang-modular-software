package entities

import "time"

type Role struct {
	RoleId    	int `gorm:"primaryKey;autoIncrement"`
	RoleName  	string`gorm:"column:role_name"`
	Code      	string `gorm:"column:code"`
	Description string `gorm:"column:role_name"`
	UniqueId  	string `gorm:"column:unique_id"`
	CreatedAt 	time.Time `gorm:"column:created_at"`
	UpdatedAt 	time.Time `gorm:"column:updated_at"`
}

func (Role) TableName() string {
	return "authentication.roles"
}

type RoleData struct {
	RoleId      int    `json:"role_id"`
	UniqueId    string `json:"unique_id"`
	RoleName    string `json:"role_name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
