package services

import (
	"context"
	"fmt"
	"github.com/ortizdavid/go-nopain/mailer"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/repositories"
	"gorm.io/gorm"
)

type EmailConfigurationService struct {
	repository *repositories.EmailConfigurationRepository
}

func NewEmailConfigurationService(db *gorm.DB) *EmailConfigurationService {
	return &EmailConfigurationService{
		repository: repositories.NewEmailConfigurationRepository(db),
	}
}

func (s *EmailConfigurationService) UpdateEmailConfiguration(ctx context.Context, request entities.UpdateEmailConfigurationRequest) error {
	conf, err := s.repository.FindFirst(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve email configuration: %s", err.Error())
	}
	conf.SMTPServer = request.SMTPServer
	conf.SMTPPort = request.SMTPPort
	conf.SenderEmail = request.SenderEmail
	conf.SenderPassword = request.SenderPassword
	err = s.repository.Update(ctx, conf)
	if err != nil {
		return fmt.Errorf("failed to update email configuration: %s", err.Error())
	}
	return nil
}

func (s *EmailConfigurationService) GetEmailConfiguration(ctx context.Context) (entities.EmailConfiguration, error) {
	conf, err := s.repository.FindFirst(ctx)
	if err != nil {
		return entities.EmailConfiguration{}, fmt.Errorf("failed to retrieve email configuration: %s", err.Error())
	}
	return conf, nil
}

func (s *EmailConfigurationService) GetDefaultEmailService(ctx context.Context) (*mailer.EmailService, error) {
	conf, err := s.GetEmailConfiguration(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed get default mail service: %s", err.Error())
	}
	sMailer := mailer.NewEmailService(
		conf.SenderEmail,
		conf.SenderPassword,
		conf.SMTPServer,
		conf.SMTPPort,
	)
	return &sMailer, nil
}