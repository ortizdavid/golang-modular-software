package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type EmployeePhoneService struct {
	repository *repositories.EmployeePhoneRepository
	employeeRepository *repositories.EmployeeRepository
}

func NewEmployeePhoneService(db *database.Database) *EmployeePhoneService {
	return &EmployeePhoneService{
		repository:         repositories.NewEmployeePhoneRepository(db),
		employeeRepository: repositories.NewEmployeeRepository(db),
	}
}

func (s *EmployeePhoneService) Create(ctx context.Context, fiberCtx *fiber.Ctx,  request entities.CreateEmployeePhoneRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	employee, err := s.employeeRepository.FindById(ctx, request.EmployeeId)
	if err != nil {
		return apperrors.NewNotFoundError("employee not found")
	}
	exists, err := s.repository.Exists(ctx, request.PhoneNumber)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("employee phone already exists for employee: "+employee.FirstName)
	}
	employeePhone := entities.EmployeePhone{
		EmployeeId:    request.EmployeeId,
		ContactTypeId: request.ContactTypeId,
		PhoneNumber:   request.DialingCode+" "+request.PhoneNumber,
		BaseEntity:    shared.BaseEntity{
			UniqueId:      encryption.GenerateUUID(),
			CreatedAt:     time.Now().UTC(),
			UpdatedAt:     time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, employeePhone)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating employee phone: " + err.Error())
	}
	return nil
}

func (s *EmployeePhoneService) Update(ctx context.Context, employeePhoneId int64, request entities.UpdateEmployeePhoneRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	employeePhone, err := s.repository.FindById(ctx, employeePhoneId)
	if err != nil {
		return apperrors.NewNotFoundError("employee phone not found")
	}
	employeePhone.EmployeeId = request.EmployeeId
	employeePhone.ContactTypeId = request.ContactTypeId
	employeePhone.PhoneNumber = request.PhoneNumber
	employeePhone.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, employeePhone)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating employee phone: " + err.Error())
	}
	return nil
}

func (s *EmployeePhoneService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam, employeeId int64) (*helpers.Pagination[entities.EmployeePhoneData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.CountByEmployee(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No employee phones found for employee: ")
	}
	employeePhones, err := s.repository.FindAllByEmployeeIdLimit(ctx, params.Limit, params.CurrentPage, employeeId)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, employeePhones, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *EmployeePhoneService) GetAll(ctx context.Context, employeeId int64) ([]entities.EmployeePhoneData, error) {
	_, err := s.repository.CountByEmployee(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No employee phones found")
	}
	employeePhones, err := s.repository.FindAllByEmployeeId(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return employeePhones, nil
}

func (s *EmployeePhoneService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchEmployeePhoneRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.EmployeePhone], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No employee phones found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	employeePhones, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, employeePhones, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *EmployeePhoneService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeePhoneData, error) {
	employeePhone, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.EmployeePhoneData{}, apperrors.NewNotFoundError("employee phone not found")
	}
	return employeePhone, nil
}

func (s *EmployeePhoneService) Remove(ctx context.Context, uniqueId string) error {
	err := s.repository.DeleteByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing employee phone: "+ err.Error())
	}
	return nil
}