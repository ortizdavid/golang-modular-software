package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type User struct {
	UserId    	int64 `gorm:"autoIncrement;primarykey"`
	UserName  	string `gorm:"column:user_name"`
	Email  		string `gorm:"column:email"`
	Password  	string `gorm:"column:password"`
	IsActive  	bool `gorm:"column:is_active"`
	UserImage  	string `gorm:"column:user_image"`
	InitialRole	string `gorm:"column:initial_role"`
	Token  		string `gorm:"column:token"`
	Status      LoginActivityStatus `gorm:"column:status"`
	shared.BaseEntity
}

func (User) TableName() string {
	return "authentication.users"
}

type UserData struct {
	UserId       	int64 `json:"user_id"`
	UniqueId     	string `json:"unique_id"`
	Token        	string `json:"token"`
	UserName     	string `json:"user_name"`
	Email        	string `json:"email"`
	Password     	string `json:"password"`  
	IsActive     	string `json:"is_active"`
	UserImage    	string `json:"user_image"`
	CreatedAt    	string `json:"created_at"`
	UpdatedAt    	string `json:"updated_at"`
	Status     		string `json:"status"`
}

