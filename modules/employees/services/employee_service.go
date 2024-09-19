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
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type EmployeeService struct {
	repository *repositories.EmployeeRepository
}

func NewEmployeeService(db *database.Database) *EmployeeService {
	return &EmployeeService{
		repository: repositories.NewEmployeeRepository(db),
	}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, request entities.CreateEmployeeRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByIdentificationNumber(ctx, request.IdentificationNumber)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("identification number '"+request.IdentificationNumber+"' already exists")
	}
	employee := entities.Employee{
		IdentificationTypeId: request.IdentificationTypeId,
		CountryId:            request.CountryId,
		MaritalStatusId:      request.MaritalStatusId,
		DepartmentId:         request.DepartmentId,
		JobTitleId:           request.JobTitleId,
		EmploymentStatusId:   request.EmploymentStatusId,
		FirstName:            request.FirstName,
		LastName:             request.LastName,
		IdentificationNumber: request.IdentificationNumber,
		Gender:               request.Gender,
		DateOfBirth:          datetime.StringToDate(request.DateOfBirth),
		BaseEntity:           shared.BaseEntity{
			UniqueId:             encryption.GenerateUUID(),
			CreatedAt:            time.Now().UTC(),
			UpdatedAt:            time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, employee)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating employee: " + err.Error())
	}
	return nil
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, employeeId int64, request entities.UpdateEmployeeRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	employee, err := s.repository.FindById(ctx, employeeId)
	if err != nil {
		return apperrors.NewNotFoundError("employee not found")
	}
	employee.IdentificationTypeId = request.IdentificationTypeId
	employee.CountryId = request.CountryId
	employee.MaritalStatusId = request.MaritalStatusId
	employee.DepartmentId = request.DepartmentId
	employee.JobTitleId = request.JobTitleId
	employee.EmploymentStatusId = request.EmploymentStatusId
	employee.FirstName = request.FirstName
	employee.LastName = request.LastName
	employee.IdentificationNumber = request.IdentificationNumber
	employee.Gender = request.Gender
	employee.DateOfBirth = datetime.StringToDate(request.DateOfBirth)
	employee.UpdatedAt = time.Now().UTC()
	//Update
	err = s.repository.Update(ctx, employee)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating employee: " + err.Error())
	}
	return nil
}

func (s *EmployeeService) GetAllEmployeesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.EmployeeData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No employees found")
	}
	employees, err := s.repository.FindAllDataLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, employees, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *EmployeeService) GetAllEmployees(ctx context.Context) ([]entities.Employee, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No employees found")
	}
	employees, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return employees, nil
}

func (s *EmployeeService) SearchEmployees(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchEmployeeRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.EmployeeData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No employees found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	employees, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, employees, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *EmployeeService) GetEmployeeByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeeData, error) {
	employee, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.EmployeeData{}, apperrors.NewNotFoundError("employee not found")
	}
	return employee, nil
}

func (s *EmployeeService) RemoveEmployee(ctx context.Context, uniqueId string) error {
	employee, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("employee not found")
	}
	err = s.repository.Delete(ctx, employee)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing employee: "+ err.Error())
	}
	return nil
}

func (s *EmployeeService) UpdateEmployeeUserId(ctx context.Context, employeeId int64, userId int64) error {
	employee, err := s.repository.FindById(ctx, employeeId)
	if err != nil {
		return apperrors.NewNotFoundError("employee not found")
	}
	employee.UserId = userId
	employee.UpdatedAt = time.Now().UTC()
	s.repository.Update(ctx, employee)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating employee userId: " + err.Error())
	}
	return nil
}
