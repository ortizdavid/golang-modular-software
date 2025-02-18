package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type OfficeService struct {
	repository        *repositories.OfficeRepository
	companyRepository *repositories.CompanyRepository
}

func NewOfficeService(db *database.Database) *OfficeService {
	return &OfficeService{
		repository:        repositories.NewOfficeRepository(db),
		companyRepository: repositories.NewCompanyRepository(db),
	}
}

func (s *OfficeService) Create(ctx context.Context, request entities.CreateOfficeRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	company, err := s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NotFoundError("company not found")
	}
	exists, err := s.repository.ExistsByName(ctx, company.CompanyId, request.OfficeName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.BadRequestError("Office already exists for company " + company.CompanyName)
	}
	office := entities.Office{
		CompanyId:  company.CompanyId,
		OfficeName: request.OfficeName,
		Code:       request.Code,
		Address:    request.Address,
		Phone:      request.Phone,
		Email:      request.Email,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, office)
	if err != nil {
		return apperrors.InternalServerError("error while creating office: " + err.Error())
	}
	return nil
}

func (s *OfficeService) Update(ctx context.Context, uniqueId string, request entities.UpdateOfficeRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	office, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NotFoundError("office not found")
	}
	_, err = s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NotFoundError("company not found")
	}
	office.CompanyId = request.CompanyId
	office.OfficeName = request.OfficeName
	office.Address = request.Address
	office.Phone = request.Phone
	office.Email = request.Email
	office.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, office)
	if err != nil {
		return apperrors.InternalServerError("error while updating office: " + err.Error())
	}
	return nil
}

func (s *OfficeService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.OfficeData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No offices found")
	}
	offices, err := s.repository.FindAllDataLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, offices, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *OfficeService) GetAll(ctx context.Context) ([]entities.OfficeData, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No offices found")
	}
	offices, err := s.repository.FindAllData(ctx)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	return offices, nil
}

func (s *OfficeService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.OfficeData, error) {
	office, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.OfficeData{}, apperrors.NotFoundError("office not found")
	}
	return office, nil
}

func (s *OfficeService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchOfficeRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.OfficeData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NotFoundError("No offices found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	offices, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, offices, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}
