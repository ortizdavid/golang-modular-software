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

type EmployeeEmailService struct {
	repository         *repositories.EmployeeEmailRepository
	employeeRepository *repositories.EmployeeRepository
}

func NewEmployeeEmailService(db *database.Database) *EmployeeEmailService {
	return &EmployeeEmailService{
		repository:         repositories.NewEmployeeEmailRepository(db),
		employeeRepository: repositories.NewEmployeeRepository(db),
	}
}

func (s *EmployeeEmailService) Create(ctx context.Context, fiberCtx *fiber.Ctx, request entities.CreateEmployeeEmailRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	employee, err := s.employeeRepository.FindById(ctx, request.EmployeeId)
	if err != nil {
		return apperrors.NotFoundError("employee not found")
	}
	exists, err := s.repository.Exists(ctx, request.EmailAddress)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.BadRequestError("employee email already exists for employee: " + employee.FirstName)
	}
	employeeEmail := entities.EmployeeEmail{
		EmployeeId:    request.EmployeeId,
		ContactTypeId: request.ContactTypeId,
		EmailAddress:  request.EmailAddress,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, employeeEmail)
	if err != nil {
		return apperrors.InternalServerError("error while creating employee email: " + err.Error())
	}
	return nil
}

func (s *EmployeeEmailService) Update(ctx context.Context, employeeEmailId int64, request entities.UpdateEmployeeEmailRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	employeeEmail, err := s.repository.FindById(ctx, employeeEmailId)
	if err != nil {
		return apperrors.NotFoundError("employee email not found")
	}
	employeeEmail.EmployeeId = request.EmployeeId
	employeeEmail.ContactTypeId = request.ContactTypeId
	employeeEmail.EmailAddress = request.EmailAddress
	employeeEmail.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, employeeEmail)
	if err != nil {
		return apperrors.InternalServerError("error while updating employee email: " + err.Error())
	}
	return nil
}

func (s *EmployeeEmailService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam, employeeId int64) (*helpers.Pagination[entities.EmployeeEmailData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	count, err := s.repository.CountByEmployee(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NotFoundError("No employee emails found for employee: ")
	}
	employeeEmails, err := s.repository.FindAllByEmployeeIdLimit(ctx, params.Limit, params.CurrentPage, employeeId)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, employeeEmails, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *EmployeeEmailService) GetAll(ctx context.Context, employeeId int64) ([]entities.EmployeeEmailData, error) {
	_, err := s.repository.CountByEmployee(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NotFoundError("No employee emails found")
	}
	employeeEmails, err := s.repository.FindAllByEmployeeId(ctx, employeeId)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	return employeeEmails, nil
}

func (s *EmployeeEmailService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchEmployeeEmailRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.EmployeeEmail], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NotFoundError("No employee emails found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	employeeEmails, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, employeeEmails, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *EmployeeEmailService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeeEmailData, error) {
	employeeEmail, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.EmployeeEmailData{}, apperrors.NotFoundError("employee email not found")
	}
	if employeeEmail.EmployeeId == 0 {
		return entities.EmployeeEmailData{}, apperrors.NotFoundError("employee email not found")
	}
	return employeeEmail, nil
}

func (s *EmployeeEmailService) GetAllByEmployeeUniqueId(ctx context.Context, uniqueId string) ([]entities.EmployeeEmailData, error) {
	employeeEmails, err := s.repository.GetAllByEmployeeUniqueId(ctx, uniqueId)
	if err != nil {
		return nil, apperrors.NotFoundError("employee email not found")
	}
	return employeeEmails, nil
}

func (s *EmployeeEmailService) Remove(ctx context.Context, uniqueId string) error {
	err := s.repository.DeleteByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.InternalServerError("error while removing employee email: " + err.Error())
	}
	return nil
}
