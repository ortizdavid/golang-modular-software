package entities

type Role struct {
	RoleId		int `gorm:"primaryKey;autoIncrement"`
	RoleName	string `gorm:"column:role_name;type:varchar(100)"`
	Code		string `gorm:"column:code;type:varchar(20)"`
}

func (Role) TableName() string {
	return "roles"
}