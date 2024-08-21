package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	"github.com/ortizdavid/golang-modular-software/modules/references/repositories"
)

type CountryService struct {
	repository *repositories.CountryRepository
}

func NewCountryService(db *database.Database) *CountryService {
	return &CountryService{
		repository: repositories.NewCountryRepository(db),
	}
}

func (s *CountryService) CreateCountry(ctx context.Context, request entities.CreateCountryRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByName(ctx, request.CountryName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("Country already exists")
	}
	country := entities.Country{
		CountryName: request.CountryName,
		IsoCode:     request.IsoCode,
		DialingCode: request.DialingCode,
		UniqueId:    encryption.GenerateUUID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	err = s.repository.Create(ctx, country)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating country: " + err.Error())
	}
	return nil
}

func (s *CountryService) UpdateCountry(ctx context.Context, countryId int, request entities.UpdateCountryRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	country, err := s.repository.FindById(ctx, countryId)
	if err != nil {
		return apperrors.NewNotFoundError("country not found")
	}
	country.CountryName = request.CountryName
	country.IsoCode = request.IsoCode
	country.DialingCode = request.DialingCode
	country.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, country)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating country: " + err.Error())
	}
	return nil
}

func (s *CountryService) GetAllCountriesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.CountryData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No countries found")
	}
	countries, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, countries, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *CountryService) GetAllCountries(ctx context.Context) ([]entities.CountryData, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No countries found")
	}
	countries, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return countries, nil
}

func (s *CountryService) SearchCountries(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchCountryRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.CountryData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No countries found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	countries, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, countries, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *CountryService) GetCountryByUniqueId(ctx context.Context, uniqueId string) (entities.CountryData, error) {
	country, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.CountryData{}, apperrors.NewNotFoundError("country not found")
	}
	return country, nil
}

func (s *CountryService) RemoveCountry(ctx context.Context, uniqueId string) error {
	country, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("country not found")
	}
	err = s.repository.Delete(ctx, country)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing country: "+ err.Error())
	}
	return nil
}