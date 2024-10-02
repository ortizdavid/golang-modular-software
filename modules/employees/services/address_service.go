package services

import (
	"context"
	"time"

	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type AddressService struct {
	repository         *repositories.AddressRepository
	employeeRepository *repositories.EmployeeRepository
}

func NewAddressService(db *database.Database) *AddressService {
	return &AddressService{
		repository:         repositories.NewAddressRepository(db),
		employeeRepository: repositories.NewEmployeeRepository(db),
	}
}

func (s *AddressService) Create(ctx context.Context, request entities.CreateAddressRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	employee, err := s.employeeRepository.FindById(ctx, request.EmployeeId)
	if err != nil {
		return apperrors.NewNotFoundError("employee not found")
	}
	exists, err := s.repository.Exists(ctx, request)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("address already exists for employee: " + employee.FirstName)
	}
	// Update Current Address as false
	err = s.repository.UpdateCurrent(ctx, employee.EmployeeId)
	if err != nil {
		return err
	}
	// Create Address
	address := entities.Address{
		EmployeeId:        request.EmployeeId,
		State:             request.State,
		City:              request.City,
		Neighborhood:      request.Neighborhood,
		Street:            request.Street,
		HouseNumber:       request.HouseNumber,
		PostalCode:        request.PostalCode,
		CountryCode:       request.CountryCode,
		AdditionalDetails: request.AdditionalDetails,
		IsCurrent:         true,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, address)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating address: " + err.Error())
	}
	return nil
}

func (s *AddressService) Update(ctx context.Context, addressId int64, request entities.UpdateAddressRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	address, err := s.repository.FindById(ctx, addressId)
	if err != nil {
		return apperrors.NewNotFoundError("address not found")
	}
	if !address.IsCurrent {
		return apperrors.NewBadRequestError("Cannot update address. Only the current address can be updated.")
	}
	address.EmployeeId = request.EmployeeId
	address.State = request.State
	address.City = request.City
	address.Neighborhood = request.Neighborhood
	address.Street = request.Street
	address.HouseNumber = request.HouseNumber
	address.PostalCode = request.PostalCode
	address.CountryCode = request.CountryCode
	address.AdditionalDetails = request.AdditionalDetails
	address.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, address)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating address: " + err.Error())
	}
	return nil
}

func (s *AddressService) GetAll(ctx context.Context, employeeId int64) ([]entities.AddressData, error) {
	_, err := s.repository.CountByEmployee(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No addresss found")
	}
	addresss, err := s.repository.FindAllByEmployeeId(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return addresss, nil
}

func (s *AddressService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.AddressData, error) {
	address, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.AddressData{}, apperrors.NewNotFoundError("address not found")
	}
	return address, nil
}

func (s *AddressService) GetAllByEmployeeUniqueId(ctx context.Context, uniqueId string) ([]entities.AddressData, error) {
	addresses, err := s.repository.GetAllByEmployeeUniqueId(ctx, uniqueId)
	if err != nil {
		return nil, apperrors.NewNotFoundError("address not found")
	}
	return addresses, nil
}

func (s *AddressService) Remove(ctx context.Context, uniqueId string) error {
	err := s.repository.DeleteByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing address: " + err.Error())
	}
	return nil
}
