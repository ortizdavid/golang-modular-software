package entities


type EmailConfiguration struct {
	ConfigurationId		int `gorm:"column:primaryKey"`
	SMTPServer   		string `gorm:"column:smtp_server"`
	SMTPPort   			int `gorm:"column:smtp_port"`
	SenderEmail   		string `gorm:"column:ender_email"`
	SenderPassword   	string `gorm:"column:sender_password"`
}


func (EmailConfiguration) TableName() string {
	return "configurations.email_configuration"
}
