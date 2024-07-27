package services

import (
	"context"
	"fmt"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/repositories"
)

type BasicConfigurationService struct {
	repository *repositories.BasicConfigurationRepository
}

func NewBasicConfigurationService(db *database.Database) *BasicConfigurationService {
	return &BasicConfigurationService{
		repository: repositories.NewBasicConfigurationRepository(db),
	}
}

func (s *BasicConfigurationService) UpdateBasicConfiguration(ctx context.Context, request entities.UpdateBasicConfigurationRequest) error {
	conf, err := s.repository.FindFirst(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve basic configuration: %s", err.Error())
	}
	conf.MaxAdmninUsers = request.MaxAdmninUsers
	conf.MaxSuperAdmninUsers = request.MaxRecordPerPage
	conf.MaxRecordsPerPage = request.MaxRecordPerPage
	err = s.repository.Update(ctx, conf)
	if err != nil {
		return fmt.Errorf("failed to update basic configuration: %s", err.Error())
	}
	return nil
}

func (s *BasicConfigurationService) GetBasicConfiguration(ctx context.Context) (entities.BasicConfiguration, error) {
	conf, err := s.repository.FindFirst(ctx)
	maxRecords := 20
	if conf.MaxRecordsPerPage < maxRecords {
		conf.MaxRecordsPerPage = maxRecords
		s.repository.Update(ctx, conf)
	}
	return conf, err
}
