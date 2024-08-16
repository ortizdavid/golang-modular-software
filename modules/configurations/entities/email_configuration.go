package entities

import "time"

type EmailConfiguration struct {
    ConfigurationId int `gorm:"autoIncrement;primaryKey" json:"configuration_id"`
    SMTPServer      string `gorm:"column:smtp_server" json:"smtp_server"`
    SMTPPort        string `gorm:"column:smtp_port" json:"smtp_port"`
    SenderEmail     string `gorm:"column:sender_email" json:"sender_email"`
    SenderPassword  string `gorm:"column:sender_password" json:"sender_password"`
    UniqueId        string `gorm:"column:unique_id" json:"unique_id"`
    CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (EmailConfiguration) TableName() string {
	return "configurations.email_configuration"
}
