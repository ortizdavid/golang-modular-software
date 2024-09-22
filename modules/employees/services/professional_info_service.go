package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type ProfessionalInfoService struct {
	repository *repositories.ProfessionalInfoRepository
	employeeRepository *repositories.EmployeeRepository
}

func NewProfessionalInfoService(db *database.Database) *ProfessionalInfoService {
	return &ProfessionalInfoService{
		repository:         repositories.NewProfessionalInfoRepository(db),
		employeeRepository: repositories.NewEmployeeRepository(db),
	}
}

func (s *ProfessionalInfoService) CreateProfessionalInfo(ctx context.Context, fiberCtx *fiber.Ctx,  request entities.CreateProfessionalInfoRequest) error {
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
		return apperrors.NewBadRequestError("professional info already exists for employee: "+employee.FirstName)
	}
	professionalInfo := entities.ProfessionalInfo{
		EmployeeId:         request.EmployeeId,
		DepartmentId:       request.DepartmentId,
		JobTitleId:         request.JobTitleId,
		EmploymentStatusId: request.EmploymentStatusId,
		BaseEntity:         shared.BaseEntity{
			UniqueId: encryption.GenerateUUID(), 
			CreatedAt: time.Now().UTC(), 
			UpdatedAt: time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, professionalInfo)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating professional info: " + err.Error())
	}
	return nil
}

func (s *ProfessionalInfoService) UpdateProfessionalInfo(ctx context.Context, professionalInfoId int64, request entities.UpdateProfessionalInfoRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	professionalInfo, err := s.repository.FindById(ctx, professionalInfoId)
	if err != nil {
		return apperrors.NewNotFoundError("professional info not found")
	}
	professionalInfo.EmployeeId = request.EmployeeId
	professionalInfo.DepartmentId = request.DepartmentId
	professionalInfo.JobTitleId = request.JobTitleId
	professionalInfo.EmploymentStatusId = request.EmploymentStatusId
	professionalInfo.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, professionalInfo)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating professional info: " + err.Error())
	}
	return nil
}

func (s *ProfessionalInfoService) GetProfessionalInfoByUniqueId(ctx context.Context, uniqueId string) (entities.ProfessionalInfoData, error) {
	professionalInfo, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.ProfessionalInfoData{}, apperrors.NewNotFoundError("professional info not found")
	}
	return professionalInfo, nil
}

func (s *ProfessionalInfoService) GetProfessionalInfoByEmployeeId(ctx context.Context, employeeId int64) (entities.ProfessionalInfoData, error) {
	professionalInfo, err := s.repository.GetDataByEmployeeId(ctx, employeeId)
	if err != nil {
		return entities.ProfessionalInfoData{}, apperrors.NewNotFoundError("professional info not found")
	}
	return professionalInfo, nil
}

func (s *ProfessionalInfoService) RemoveProfessionalInfo(ctx context.Context, uniqueId string) error {
	professionalInfo, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("professional info not found")
	}
	err = s.repository.Delete(ctx, professionalInfo)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing professional info: "+ err.Error())
	}
	return nil
}