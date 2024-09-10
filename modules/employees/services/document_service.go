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


func (s *DocumentService) CreateDocument(ctx context.Context, fiberCtx *fiber.Ctx,  request entities.CreateDocumentRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
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
	info, err := uploader.UploadSingleFile(fiberCtx, "file_name")
	if err != nil {
		return apperrors.NewNotFoundError("error while uploading employee document: "+err.Error())
	}
	//----------------------------------------
	document := entities.Document{
		EmployeeId:     request.EmployeeId,
		DocumentTypeId: request.DocumentTypeId,
		DocumentName:   request.DocumentName,
		DocumentNumber: request.DocumentNumber,
		ExpirationDate: datetime.StringToDate(request.ExpirationDate),
		FileName:       info.FinalName,
		Status:         request.Status,
		UniqueId:       encryption.GenerateUUID(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
	err = s.repository.Create(ctx, document)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating document: " + err.Error())
	}
	return nil
}

func (s *DocumentService) UpdateDocument(ctx context.Context, documentId int, request entities.UpdateDocumentRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	document, err := s.repository.FindById(ctx, documentId)
	if err != nil {
		return apperrors.NewNotFoundError("document not found")
	}
	document.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, document)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating document: " + err.Error())
	}
	return nil
}

func (s *DocumentService) GetAllEmployeeDocumentsPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam, employeeId int64) (*helpers.Pagination[entities.Document], error) {
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

func (s *DocumentService) GetAllEmployeeDocuments(ctx context.Context, employeeId int64) ([]entities.Document, error) {
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

func (s *DocumentService) SearchDocuments(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchDocumentRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.Document], error) {
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

func (s *DocumentService) GetDocumentByUniqueId(ctx context.Context, uniqueId string) (entities.Document, error) {
	document, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.Document{}, apperrors.NewNotFoundError("document not found")
	}
	return document, nil
}

func (s *DocumentService) RemoveDocument(ctx context.Context, uniqueId string) error {
	document, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("document not found")
	}
	err = s.repository.Delete(ctx, document)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing document: "+ err.Error())
	}
	return nil
}