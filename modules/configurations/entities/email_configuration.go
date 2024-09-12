package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type EmailConfiguration struct {
    ConfigurationId int `gorm:"autoIncrement;primaryKey" json:"configuration_id"`
    SMTPServer      string `gorm:"column:smtp_server" json:"smtp_server"`
    SMTPPort        string `gorm:"column:smtp_port" json:"smtp_port"`
    SenderEmail     string `gorm:"column:sender_email" json:"sender_email"`
    SenderPassword  string `gorm:"column:sender_password" json:"sender_password"`
    shared.BaseEntity
}

func (EmailConfiguration) TableName() string {
	return "configurations.email_configuration"
}
