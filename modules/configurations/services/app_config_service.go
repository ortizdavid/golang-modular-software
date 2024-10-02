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
	basicConfig, err := s.basicService.GetCurrent(ctx)
	if err != nil {
		log.Printf("Failed to load basic configuration: %v", err)
		basicConfig = entities.BasicConfiguration{} 
	}
	companyConfig, err := s.companyService.GetCurrent(ctx)
	if err != nil {
		log.Printf("Failed to load company configuration: %v", err)
		companyConfig = entities.CompanyConfiguration{} 
	}
	emailConfig, err := s.emailService.GetCurrent(ctx)
	if err != nil {
		log.Printf("Failed to load email configuration: %v", err)
		emailConfig = entities.EmailConfiguration{} 
	}
	return &entities.AppConfiguration{
		BasicConfig:   basicConfig,
		CompanyConfig: companyConfig,
		EmailConfig:   emailConfig,
	}
}