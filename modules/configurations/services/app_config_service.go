package services

import (
	"context"
	"log"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

type AppConfigurationService struct {
	basicService *BasicConfigurationService
	companyService *CompanyConfigurationService
	emailService *EmailConfigurationService
}

func NewAppConfigurationService(db *database.Database) *AppConfigurationService {
	return &AppConfigurationService{
		basicService:   NewBasicConfigurationService(db),
		companyService: NewCompanyConfigurationService(db),
		emailService:   NewEmailConfigurationService(db),
	}
}

func(s *AppConfigurationService) LoadAppConfigurations(ctx context.Context) *entities.AppConfiguration {
	basicConfig, err := s.basicService.GetBasicConfiguration(ctx)
	if err != nil {
		log.Printf("Failed to load basic configuration: %v", err)
		basicConfig = entities.BasicConfiguration{} // Provide a default configuration or handle accordingly
	}
	companyConfig, err := s.companyService.GetCompanyConfiguration(ctx)
	if err != nil {
		log.Printf("Failed to load company configuration: %v", err)
		companyConfig = entities.CompanyConfiguration{} // Provide a default configuration or handle accordingly
	}
	emailConfig, err := s.emailService.GetEmailConfiguration(ctx)
	if err != nil {
		log.Printf("Failed to load email configuration: %v", err)
		emailConfig = entities.EmailConfiguration{} // Provide a default configuration or handle accordingly
	}
	return &entities.AppConfiguration{
		BasicConfig:   basicConfig,
		CompanyConfig: companyConfig,
		EmailConfig:   emailConfig,
	}
}