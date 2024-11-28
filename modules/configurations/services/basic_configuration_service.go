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

type BasicConfigurationService struct {
	repository *repositories.BasicConfigurationRepository
}

func NewBasicConfigurationService(db *database.Database) *BasicConfigurationService {
	return &BasicConfigurationService{
		repository: repositories.NewBasicConfigurationRepository(db),
	}
}

func (s *BasicConfigurationService) Update(ctx context.Context, request entities.UpdateBasicConfigurationRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	conf, err := s.repository.FindLast(ctx)
	if err != nil {
		// Create a new configuration if none exists
		conf = entities.BasicConfiguration{
			ConfigurationId:     0,
			AppName:             request.AppName,
			AppAcronym:          request.AppAcronym,
			MaxRecordsPerPage:   request.MaxRecordPerPage,
			MaxAdmninUsers:      request.MaxAdmninUsers,
			MaxSuperAdmninUsers: request.MaxSuperAdminUsers,
			BaseEntity: shared.BaseEntity{
				UniqueId:  encryption.GenerateUUID(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			},
		}
		err = s.repository.Create(ctx, conf)
		if err != nil {
			return fmt.Errorf("failed to create basic configuration: %w", err)
		}
		return nil
	}
	// Update the existing configuration
	conf.AppName = request.AppName
	conf.AppAcronym = request.AppAcronym
	conf.MaxAdmninUsers = request.MaxAdmninUsers
	conf.MaxSuperAdmninUsers = request.MaxSuperAdminUsers // Fixed the assignment here
	conf.MaxRecordsPerPage = request.MaxRecordPerPage
	conf.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, conf)
	if err != nil {
		return fmt.Errorf("failed to update basic configuration: %w", err)
	}
	return nil
}

func (s *BasicConfigurationService) GetCurrent(ctx context.Context) (entities.BasicConfiguration, error) {
	conf, err := s.repository.FindLast(ctx)
	maxRecords := 20
	if conf.MaxRecordsPerPage < maxRecords {
		conf.MaxRecordsPerPage = maxRecords
		s.repository.Update(ctx, conf)
	}
	return conf, err
}
