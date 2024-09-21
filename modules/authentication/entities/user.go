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
	Token  		string `gorm:"column:token"`
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
	Host       		string `json:"host"`
	Browser    		string `json:"browser"`
	IPAddress  		string `json:"ip_address"`
	Device     		string `json:"device"`
	Location   		string `json:"location"`
	LastLogin  		string `json:"last_login"`
	LastLogout 		string `json:"last_logout"`
	TotalLogin		int64 `json:"total_login"`
	TotalLogout		int64 `json:"total_logout"`
}

