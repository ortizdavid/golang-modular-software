package entities


type EmailConfiguration struct {
	ConfigurationId		int `gorm:"column:ConfigurationId;primaryKey"`
	SMTPServer   		string `gorm:"column:SMTPServer;type:varchar(50)"`
	SMTPPort   			int `gorm:"column:SMTPPort;type:int"`
	SenderEmail   		string `gorm:"column:SenderEmail;type:varchar(100);int"`
	SenderPassword   	string `gorm:"column:SenderPassword;type:varchar(100);;int"`
}


func (EmailConfiguration) TableName() string {
	return "EmailConfiguration"
}
