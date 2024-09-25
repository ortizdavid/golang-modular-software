package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/go-nopain/mailer"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type EmailConfigurationService struct {
	repository *repositories.EmailConfigurationRepository
}

func NewEmailConfigurationService(db *database.Database) *EmailConfigurationService {
	return &EmailConfigurationService{
		repository: repositories.NewEmailConfigurationRepository(db),
	}
}

func (s *EmailConfigurationService) UpdateEmailConfiguration(ctx context.Context, request entities.UpdateEmailConfigurationRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	// Attempt to retrieve the existing configuration
    conf, err := s.repository.FindLast(ctx)
    if err != nil {
		// Create a new configuration if none exists
		conf = entities.EmailConfiguration{
			ConfigurationId: 0,
			SMTPServer:      request.SMTPServer,
			SMTPPort:        request.SMTPPort,
			SenderEmail:     request.SenderEmail,
			SenderPassword:  request.SenderPassword,
			BaseEntity:      shared.BaseEntity{
				UniqueId:         encryption.GenerateUUID(),
				CreatedAt:        time.Now().UTC(),
				UpdatedAt:        time.Now().UTC(),
			},
		}
		err = s.repository.Create(ctx, conf)
		if err != nil {
			return fmt.Errorf("failed to create email configuration: %w", err)
		}
		return nil
    }
    // Update the existing configuration with new values
    conf.SMTPServer = request.SMTPServer
    conf.SMTPPort = request.SMTPPort
    conf.SenderEmail = request.SenderEmail
    conf.SenderPassword = request.SenderPassword
    err = s.repository.Update(ctx, conf)
    if err != nil {
        return fmt.Errorf("failed to update email configuration: %w", err)
    }
    return nil
}


func (s *EmailConfigurationService) GetEmailConfiguration(ctx context.Context) (entities.EmailConfiguration, error) {
	conf, err := s.repository.FindLast(ctx)
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