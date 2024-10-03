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
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"

)

type CurrencyService struct {
	repository *repositories.CurrencyRepository
}

func NewCurrencyService(db *database.Database) *CurrencyService {
	return &CurrencyService{
		repository: repositories.NewCurrencyRepository(db),
	}
}

func (s *CurrencyService) Create(ctx context.Context, request entities.CreateCurrencyRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByName(ctx, request.CurrencyName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("Currency already exists")
	}
	currency := entities.Currency{
		CurrencyName: request.CurrencyName,
		Code:         request.Code,
		BaseEntity: shared.BaseEntity{
			UniqueId:         encryption.GenerateUUID(),
			CreatedAt:        time.Now().UTC(),
			UpdatedAt:        time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, currency)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating currency: " + err.Error())
	}
	return nil
}

func (s *CurrencyService) Update(ctx context.Context, uniqueId string, request entities.UpdateCurrencyRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	currency, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("currency not found")
	}
	currency.CurrencyName = request.CurrencyName
	currency.Code = request.Code
	currency.Symbol = request.Symbol
	currency.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, currency)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating currency: " + err.Error())
	}
	return nil
}

func (s *CurrencyService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.Currency], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No currencies found")
	}
	currencies, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, currencies, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *CurrencyService) GetAll(ctx context.Context) ([]entities.Currency, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No currencies found")
	}
	currencies, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return currencies, nil
}

func (s *CurrencyService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchCurrencyRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.Currency], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No currencies found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	currencies, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, currencies, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *CurrencyService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.Currency, error) {
	currency, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.Currency{}, apperrors.NewNotFoundError("currency not found")
	}
	return currency, nil
}

func (s *CurrencyService) Remove(ctx context.Context, uniqueId string) error {
	currency, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("currency not found")
	}
	err = s.repository.Delete(ctx, currency)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing currency: "+ err.Error())
	}
	return nil
}