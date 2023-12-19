package models

import (
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/go-nopain/mailer"
)


func GetBasicConfiguration() (entities.BasicConfiguration, error) {
	return BasicConfigurationModel{}.FindFirst()
}


func GetEmailConfiguration() (entities.EmailConfiguration, error) {
	return EmailConfigurationModel{}.FindFirst()
}


func GetCompanyConfiguration() (entities.CompanyConfiguration, error) {
	return CompanyConfigurationModel{}.FindFirst()
}


func DefaultEmailService() *mailer.EmailService {
	configuracao, _ := GetEmailConfiguration()
	return mailer.NewEmailService(
		configuracao.SenderEmail,
		configuracao.SenderPassword,
		configuracao.SMTPServer,
		configuracao.SMTPPort,
	)
}