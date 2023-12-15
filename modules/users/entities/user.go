package entities

import "time"

type User struct {
	UserId    	int `gorm:"autoIncrement;primarykey"`
	RoleId  	int `gorm:"column:role_id;type:int"`
	UserName  	string `gorm:"column:user_name;type:varchar(100)"`
	Password  	string `gorm:"column:password;type:varchar(150)"`
	Active  	string `gorm:"column:active;type:enum('Yes', 'No')"`
	Image  		string `gorm:"column:image;type:varchar(100)"`
	UniqueId  	string `gorm:"column:unique_id;type:varchar(50)"`
	Token  		string `gorm:"column:token;type:varchar(200)"`
	CreatedAt 	time.Time `gorm:"column:created_at"`
	UpdatedAt  	time.Time `gorm:"column:updated_at"`
}

func TableName() string {
	return "users"
}

type UserData struct {
	UserId		int 
	UniqueId 	string
	Token 		string
	UserName 	string
	Password 	string
	Active   	string
	Image   	string
	CreatedAtStr 	string
	UpdatedAtStr  	string
	RoleId 		int
	RoleName 	string
	RoleCode 	string
}
