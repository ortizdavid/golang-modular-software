package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/datetime"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type DocumentService struct {
	repository *repositories.DocumentRepository
	employeeRepository *repositories.EmployeeRepository
}

func NewDocumentService(db *database.Database) *DocumentService {
	return &DocumentService{
		repository:         repositories.NewDocumentRepository(db),
		employeeRepository: repositories.NewEmployeeRepository(db),
	}
}

func (s *DocumentService) Create(ctx context.Context, fiberCtx *fiber.Ctx,  request entities.CreateDocumentRequest) error {
	//----- Validation
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	expirationDate, err := datetime.StringToDate(request.ExpirationDate)
	if err != nil {
		return err
	}
	//-----------------------
	employee, err := s.employeeRepository.FindById(ctx, request.EmployeeId)
	if err != nil {
		return apperrors.NewNotFoundError("employee not found")
	}
	exists, err := s.repository.ExistsByName(ctx, request.DocumentName, request.EmployeeId)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("document already exists for employee: "+employee.FirstName)
	}
	//------------Upload ------------------------------------------------
	uploadPath := config.UploadDocumentPath() + "/employees"
	uploader := helpers.NewUploader(uploadPath, config.MaxUploadDocumentSize(), helpers.ExtDocuments)
	info, err := uploader.UploadSingleFile(fiberCtx, "document_file")
	if err != nil {
		return apperrors.NewNotFoundError("error while uploading employee document: "+err.Error())
	}
	//----------------------------------------
	document := entities.Document{
		EmployeeId:     request.EmployeeId,
		DocumentTypeId: request.DocumentTypeId,
		DocumentName:   request.DocumentName,
		DocumentNumber: request.DocumentNumber,
		ExpirationDate: expirationDate,
		FileName:       info.FinalName,
		Status:         request.Status,
		BaseEntity:     shared.BaseEntity{
			UniqueId:       encryption.GenerateUUID(),
			CreatedAt:      time.Now().UTC(),
			UpdatedAt:      time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, document)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating document: " + err.Error())
	}
	return nil
}

func (s *DocumentService) Update(ctx context.Context, documentId int64, request entities.UpdateDocumentRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	expirationDate, err := datetime.StringToDate(request.ExpirationDate)
	if err != nil {
		return err
	}
	document, err := s.repository.FindById(ctx, documentId)
	if err != nil {
		return apperrors.NewNotFoundError("document not found")
	}
	document.DocumentTypeId = request.DocumentTypeId
	document.DocumentName = request.DocumentName
	document.DocumentNumber = request.DocumentNumber
	document.ExpirationDate = expirationDate
	document.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, document)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating document: " + err.Error())
	}
	return nil
}

func (s *DocumentService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam, employeeId int64) (*helpers.Pagination[entities.DocumentData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.CountByEmployee(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No documents found for employee: ")
	}
	documents, err := s.repository.FindAllByEmployeeIdLimit(ctx, params.Limit, params.CurrentPage, employeeId)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, documents, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *DocumentService) GetAllEmployeeDocuments(ctx context.Context, employeeId int64) ([]entities.DocumentData, error) {
	_, err := s.repository.CountByEmployee(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No documents found")
	}
	documents, err := s.repository.FindAllByEmployeeId(ctx, employeeId)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return documents, nil
}

func (s *DocumentService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchDocumentRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.Document], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No documents found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	documents, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, documents, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *DocumentService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.DocumentData, error) {
	document, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.DocumentData{}, apperrors.NewNotFoundError("document not found")
	}
	if document.EmployeeId == 0 {
		return entities.DocumentData{}, apperrors.NewNotFoundError("document not found")
	}
	return document, nil
}

func (s *DocumentService) GetAllByEmployeeUniqueId(ctx context.Context, uniqueId string) ([]entities.DocumentData, error) {
	documents, err := s.repository.GetAllByEmployeeUniqueId(ctx, uniqueId)
	if err != nil {
		return nil, apperrors.NewNotFoundError("document not found")
	}
	return documents, nil
}

func (s *DocumentService) Remove(ctx context.Context, uniqueId string) error {
	err := s.repository.DeleteByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing document: "+ err.Error())
	}
	return nil
}