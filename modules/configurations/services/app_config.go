package services

import (
	"context"
	"log"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

type AppConfiguration struct {
	BasicConfig entities.BasicConfiguration
	CompanyConfig entities.CompanyConfiguration
	EmailConfig entities.EmailConfiguration
}

func LoadAppConfigurations(db *database.Database) *AppConfiguration {
	ctx := context.Background()
	basicConfig, err := NewBasicConfigurationService(db).GetBasicConfiguration(ctx)
	if err != nil {
		log.Printf("Failed to load basic configuration: %v", err)
		basicConfig = entities.BasicConfiguration{} // Provide a default configuration or handle accordingly
	}
	companyConfig, err := NewCompanyConfigurationService(db).GetCompanyConfiguration(ctx)
	if err != nil {
		log.Printf("Failed to load company configuration: %v", err)
		companyConfig = entities.CompanyConfiguration{} // Provide a default configuration or handle accordingly
	}
	emailConfig, err := NewEmailConfigurationService(db).GetEmailConfiguration(ctx)
	if err != nil {
		log.Printf("Failed to load email configuration: %v", err)
		emailConfig = entities.EmailConfiguration{} // Provide a default configuration or handle accordingly
	}
	return &AppConfiguration{
		BasicConfig:   basicConfig,
		CompanyConfig: companyConfig,
		EmailConfig:   emailConfig,
	}
}