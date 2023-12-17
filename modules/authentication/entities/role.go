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
