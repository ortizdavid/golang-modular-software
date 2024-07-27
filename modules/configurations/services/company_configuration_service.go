package services

import (
	"context"
	"fmt"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/repositories"
)

type CompanyConfigurationService struct {
	repository *repositories.CompanyConfigurationRepository
}

func NewCompanyConfigurationService(db *database.Database) *CompanyConfigurationService {
	return &CompanyConfigurationService{
		repository: repositories.NewCompanyConfigurationRepository(db),
	}
}

func (s *CompanyConfigurationService) UpdateCompanyConfiguration(ctx context.Context, request entities.UpdateCompanyConfigurationRequest) error {
	conf, err := s.repository.FindFirst(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve company configuration: %s", err.Error())
	}
	conf.CompanyName= request.CompanyName
	conf.CompanyAcronym = request.CompanyAcronym
	conf.CompanyPhone = request.CompanyPhone
	conf.CompanyEmail = request.CompanyEmail
	conf.CompanyMainColor = request.CompanyMainColor
	err = s.repository.Update(ctx, conf)
	if err != nil {
		return fmt.Errorf("failed to update company configuration: %s", err.Error())
	}
	return nil
}

func (s *CompanyConfigurationService) GetCompanyConfiguration(ctx context.Context) (entities.CompanyConfiguration, error) {
	conf, err := s.repository.FindFirst(ctx)
	if err != nil {
		return entities.CompanyConfiguration{}, fmt.Errorf("failed to retrieve company configuration: %s", err.Error())
	}
	return conf, nil
}
