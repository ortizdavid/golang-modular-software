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

type DepartmentService struct {
	repository        *repositories.DepartmentRepository
	companyRepository *repositories.CompanyRepository
}

func NewDepartmentService(db *database.Database) *DepartmentService {
	return &DepartmentService{
		repository:        repositories.NewDepartmentRepository(db),
		companyRepository: repositories.NewCompanyRepository(db),
	}
}

func (s *DepartmentService) Create(ctx context.Context, request entities.CreateDepartmentRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	company, err := s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NotFoundError("company not found")
	}
	exists, err := s.repository.ExistsByName(ctx, company.CompanyId, request.DepartmentName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.BadRequestError("Department already exists for company " + company.CompanyName)
	}
	department := entities.Department{
		CompanyId:      company.CompanyId,
		DepartmentName: request.DepartmentName,
		Acronym:        request.Acronym,
		Description:    request.Description,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, department)
	if err != nil {
		return apperrors.InternalServerError("error while creating department: " + err.Error())
	}
	return nil
}

func (s *DepartmentService) Update(ctx context.Context, uniqueId string, request entities.UpdateDepartmentRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	department, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NotFoundError("department not found")
	}
	_, err = s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NotFoundError("company not found")
	}
	department.CompanyId = request.CompanyId
	department.DepartmentName = request.DepartmentName
	department.CompanyId = request.CompanyId
	department.Acronym = request.Acronym
	department.Description = request.Description
	department.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, department)
	if err != nil {
		return apperrors.InternalServerError("error while updating department: " + err.Error())
	}
	return nil
}

func (s *DepartmentService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.DepartmentData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No departments found")
	}
	departments, err := s.repository.FindAllDataLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, departments, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *DepartmentService) GetAll(ctx context.Context) ([]entities.DepartmentData, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No departments found")
	}
	departments, err := s.repository.FindAllData(ctx)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	return departments, nil
}

func (s *DepartmentService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchDepartmentRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.DepartmentData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NotFoundError("No departments found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	departments, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, departments, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *DepartmentService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.DepartmentData, error) {
	department, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.DepartmentData{}, apperrors.NotFoundError("department not found")
	}
	return department, nil
}
