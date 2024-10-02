package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type ProjectAttachmentService struct {
	repository *repositories.ProjectAttachmentRepository
	projectRepository *repositories.ProjectRepository
	companyRepository *repositories.CompanyRepository
}

func NewProjectAttachmentService(db *database.Database) *ProjectAttachmentService {
	return &ProjectAttachmentService{
		repository:        repositories.NewProjectAttachmentRepository(db),
		projectRepository:  repositories.NewProjectRepository(db),
		companyRepository: repositories.NewCompanyRepository(db),
	}
}

func (s *ProjectAttachmentService) Create(ctx context.Context, fiberCtx *fiber.Ctx,   request entities.CreateProjectAttachmentRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	project, err := s.projectRepository.FindById(ctx, request.ProjectId)
	if err != nil {
		return apperrors.NewNotFoundError("Project not found. Invalid id")
	}
	company, err := s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NewNotFoundError("Company not found. Invalid id")
	}
	//------------Upload ------------------------------------------------
	uploadPath := config.UploadDocumentPath() + "/company"
	uploader := helpers.NewUploader(uploadPath, config.MaxUploadDocumentSize(), helpers.ExtDocuments)
	info, err := uploader.UploadSingleFile(fiberCtx, "attachment_file")
	if err != nil {
		return apperrors.NewNotFoundError("error while uploading attachment document: "+err.Error())
	}
	projectAttachment := entities.ProjectAttachment{
		ProjectId:       project.ProjectId,
		CompanyId:      company.CompanyId,
		AttachmentName: request.AttachmentName,
		FileName: 		info.FinalName,
		BaseEntity:     shared.BaseEntity{
			UniqueId:      encryption.GenerateUUID(),
			CreatedAt:     time.Now().UTC(),
			UpdatedAt:     time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, projectAttachment)
	if err != nil {
		return apperrors.NewInternalServerError("")
	}
	return nil
}

func (s *ProjectAttachmentService) Update(ctx context.Context, request entities.UpdateProjectAttachmentRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	return nil
}

func (s *ProjectAttachmentService) GetAllByProjectId(ctx context.Context, projectId int) ([]entities.ProjectAttachment, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No attachments found")
	}
	attachments, err := s.repository.FindAllByProjectId(ctx, projectId)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return attachments, nil
}

func (s *ProjectAttachmentService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.ProjectAttachment, error) {
	projectAttachment, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.ProjectAttachment{}, apperrors.NewNotFoundError("project attachment not found")
	}
	return projectAttachment, nil
}

func (s *ProjectAttachmentService) Remove(ctx context.Context, uniqueId string) error {
	err := s.repository.DeleteByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing project attachment: "+err.Error())
	}
	return nil
}