package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/datetime"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
)

type CompanyService struct {
	repository *repositories.CompanyRepository
}

func NewCompanyService(db *database.Database) *CompanyService {
	return &CompanyService{
		repository: repositories.NewCompanyRepository(db),
	}
}

func (s *CompanyService) CreateCompany(ctx context.Context, request entities.CreateCompanyRequest) error {
    if err := request.Validate(); err != nil {
        return apperrors.NewBadRequestError(err.Error())
    }
    company := entities.Company{
        CompanyName:    request.CompanyName,
        CompanyAcronym: request.CompanyAcronym,
        CompanyType:    request.CompanyType,
        Industry:       request.Industry,
        FoundedDate:    datetime.StringToDate(request.FoundedDate),
        Address:        request.Address,
        Phone:          request.Phone,
        Email:          request.Email,
        WebsiteURL:     request.WebsiteURL,
        UniqueId:       encryption.GenerateUUID(),
        CreatedAt:      time.Now().UTC(),
        UpdatedAt:      time.Now().UTC(),
    }
    err := s.repository.Create(ctx, company)
    if err != nil {
        return apperrors.NewInternalServerError("error while creating company: " + err.Error())
    }
    return nil
}

func (s *CompanyService) UpdateCompany(ctx context.Context, companyId int, request entities.UpdateCompanyRequest) error {
    if err := request.Validate(); err != nil {
        return apperrors.NewBadRequestError(err.Error())
    }
    company, err := s.repository.FindById(ctx, companyId)
    if err != nil {
        return apperrors.NewNotFoundError("company not found")
    }
    company.CompanyName = request.CompanyName
    company.CompanyAcronym = request.CompanyAcronym
    company.CompanyType = request.CompanyType
    company.Industry = request.Industry
    company.FoundedDate = datetime.StringToDate(request.FoundedDate)
    company.Address = request.Address
    company.Phone = request.Phone
    company.Email = request.Email
    company.WebsiteURL = request.WebsiteURL
    company.UpdatedAt = time.Now().UTC()
    err = s.repository.Update(ctx, company)
    if err != nil {
        return apperrors.NewInternalServerError("error while updating company: " + err.Error())
    }
    return nil
}

func (s *CompanyService) GetAllCompaniesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.CompanyData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No companies found")
	}
	companies, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, companies, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *CompanyService) GetAllCompanys(ctx context.Context) ([]entities.Company, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No companys found")
	}
	companys, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	return companys, nil
}

func (s *CompanyService) SearchCompanies(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchCompanyRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.CompanyData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No companies found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	companies, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: "+err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, companies, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: "+err.Error())
	}
	return pagination, nil
}

func (s *CompanyService) GetCompanyByUniqueId(ctx context.Context, uniqueId string) (entities.CompanyData, error) {
	company, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.CompanyData{}, apperrors.NewNotFoundError("company not found")
	}
	return company, nil
}
