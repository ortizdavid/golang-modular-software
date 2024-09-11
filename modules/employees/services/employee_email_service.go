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
)

type EmployeeEmailService struct {
	repository *repositories.EmployeeEmailRepository
	employeeRepository *repositories.EmployeeRepository
}

func NewEmployeeEmailService(db *database.Database) *EmployeeEmailService {
	return &EmployeeEmailService{
		repository:         repositories.NewEmployeeEmailRepository(db),
		employeeRepository: repositories.NewEmployeeRepository(db),
	}
}

func (s *EmployeeEmailService) CreateEmployeeEmail(ctx context.Context, fiberCtx *fiber.Ctx,  request entities.CreateEmployeeEmailRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	employee, err := s.employeeRepository.FindById(ctx, request.EmployeeId)
	if err != nil {
		return apperrors.NewNotFoundError("employee not found")
	}
	exists, err := s.repository.Exists(ctx, request.EmailAddress)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("employee email already exists for employee: "+employee.FirstName)
	}
	employeeEmail := entities.EmployeeEmail{
		EmployeeId:    request.EmployeeId,
		ContactTypeId: request.ContactTypeId,
		EmailAddress:  request.EmailAddress,
		UniqueId:      encryption.GenerateUUID(),
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}
	err = s.repository.Create(ctx, employeeEmail)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating employee email: " + err.Error())
	}
	return nil
}

func (s *EmployeeEmailService) UpdateEmployeeEmail(ctx context.Context, employeeEmailId int64, request entities.UpdateEmployeeEmailRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	employeeEmail, err := s.repository.FindById(ctx, employeeEmailId)
	if err != nil {
		return apperrors.NewNotFoundError("employee email not found")
	}
	employeeEmail.EmployeeId = request.EmployeeId
	employeeEmail.ContactTypeId = request.ContactTypeId
	employeeEmail.EmailAddress = request.EmailAddress
	employeeEmail.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, employeeEmail)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating employee email: " + err.Error())
	}
	return nil
}

func (s *EmployeeEmailService) GetAllEmployeeEmployeeEmailsPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam, employeeId int64) (*helpers.Pagination[entities.EmployeeEmailData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.CountByEmployee(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No employee emails found for employee: ")
	}
	employeeEmails, err := s.repository.FindAllByEmployeeIdLimit(ctx, params.Limit, params.CurrentPage, employeeId)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, employeeEmails, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *EmployeeEmailService) GetAllEmployeeEmployeeEmails(ctx context.Context, employeeId int64) ([]entities.EmployeeEmailData, error) {
	_, err := s.repository.CountByEmployee(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No employee emails found")
	}
	employeeEmails, err := s.repository.FindAllByEmployeeId(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return employeeEmails, nil
}

func (s *EmployeeEmailService) SearchEmployeeEmails(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchEmployeeEmailRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.EmployeeEmail], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No employee emails found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	employeeEmails, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, employeeEmails, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *EmployeeEmailService) GetEmployeeEmailByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeeEmailData, error) {
	employeeEmail, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.EmployeeEmailData{}, apperrors.NewNotFoundError("employee email not found")
	}
	return employeeEmail, nil
}

func (s *EmployeeEmailService) RemoveEmployeeEmail(ctx context.Context, uniqueId string) error {
	employeeEmail, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("employee email not found")
	}
	err = s.repository.Delete(ctx, employeeEmail)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing employee email: "+ err.Error())
	}
	return nil
}