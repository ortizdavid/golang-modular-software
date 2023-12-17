package entities

import "time"


type User struct {
	UserId    	int `gorm:"autoIncrement;primarykey"`
	RoleId  	int `gorm:"column:role_id"`
	UserName  	string `gorm:"column:user_name"`
	Password  	string `gorm:"column:password"`
	Active  	string `gorm:"column:active"`
	Image  		string `gorm:"column:user_image"`
	Token  		string `gorm:"column:token;"`
	UniqueId  	string `gorm:"column:unique_id"`
	CreatedAt 	time.Time `gorm:"column:created_at"`
	UpdatedAt  	time.Time `gorm:"column:updated_at"`
}


func TableName() string {
	return "authentication.users"
}


type UserData struct {
	UserId			int 
	UniqueId 		string
	Token 			string
	UserName 		string
	Password 		string
	Active   		string
	UserImage   	string
	CreatedAtStr 	string
	UpdatedAtStr  	string
	RoleId 			int
	RoleName 		string
	RoleCode 		string
}
