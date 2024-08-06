package entities

import "time"

type User struct {
	UserId    	int64 `gorm:"autoIncrement;primarykey"`
	UserName  	string `gorm:"column:user_name"`
	Email  		string `gorm:"column:email"`
	Password  	string `gorm:"column:password"`
	IsActive  	bool `gorm:"column:is_active"`
	Image  		string `gorm:"column:user_image"`
	Token  		string `gorm:"column:token;"`
	UniqueId  	string `gorm:"column:unique_id"`
	CreatedAt 	time.Time `gorm:"column:created_at"`
	UpdatedAt  	time.Time `gorm:"column:updated_at"`
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
}

