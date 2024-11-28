package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type CompanyConfigurationService struct {
	repository *repositories.CompanyConfigurationRepository
}

func NewCompanyConfigurationService(db *database.Database) *CompanyConfigurationService {
	return &CompanyConfigurationService{
		repository: repositories.NewCompanyConfigurationRepository(db),
	}
}

func (s *CompanyConfigurationService) Update(ctx context.Context, request entities.UpdateCompanyConfigurationRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	// Attempt to retrieve the existing configuration
	conf, err := s.repository.FindLast(ctx)
	if err != nil {
		conf = entities.CompanyConfiguration{
			ConfigurationId:  0,
			CompanyName:      request.CompanyName,
			CompanyAcronym:   request.CompanyAcronym,
			CompanyMainColor: request.CompanyMainColor,
			CompanyLogo:      "",
			CompanyPhone:     request.CompanyPhone,
			CompanyEmail:     request.CompanyEmail,
			BaseEntity: shared.BaseEntity{
				UniqueId:  encryption.GenerateUUID(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			},
		}
		err = s.repository.Create(ctx, conf)
		if err != nil {
			return fmt.Errorf("failed to create company configuration: %w", err)
		}
		return nil
	}
	// Update the existing configuration with new values
	conf.CompanyName = request.CompanyName
	conf.CompanyAcronym = request.CompanyAcronym
	conf.CompanyPhone = request.CompanyPhone
	conf.CompanyEmail = request.CompanyEmail
	conf.CompanyMainColor = request.CompanyMainColor
	conf.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, conf)
	if err != nil {
		return fmt.Errorf("failed to update company configuration: %w", err)
	}
	return nil
}

func (s *CompanyConfigurationService) GetCurrent(ctx context.Context) (entities.CompanyConfiguration, error) {
	conf, err := s.repository.FindLast(ctx)
	if err != nil {
		return entities.CompanyConfiguration{}, fmt.Errorf("failed to retrieve company configuration: %s", err.Error())
	}
	return conf, nil
}
