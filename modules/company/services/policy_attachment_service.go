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

type PolicyAttachmentService struct {
	repository *repositories.PolicyAttachmentRepository
	policyRepository *repositories.PolicyRepository
	companyRepository *repositories.CompanyRepository
}

func NewPolicyAttachmentService(db *database.Database) *PolicyAttachmentService {
	return &PolicyAttachmentService{
		repository:        repositories.NewPolicyAttachmentRepository(db),
		policyRepository:  repositories.NewPolicyRepository(db),
		companyRepository: repositories.NewCompanyRepository(db),
	}
}

func (s *PolicyAttachmentService) Create(ctx context.Context, fiberCtx *fiber.Ctx,   request entities.CreatePolicyAttachmentRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	policy, err := s.policyRepository.FindById(ctx, request.PolicyId)
	if err != nil {
		return apperrors.NewNotFoundError("Policy not found. Invalid id")
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
	policyAttachment := entities.PolicyAttachment{
		PolicyId:       policy.PolicyId,
		CompanyId:      company.CompanyId,
		AttachmentName: request.AttachmentName,
		FileName: 		info.FinalName,
		BaseEntity:     shared.BaseEntity{
			UniqueId:      encryption.GenerateUUID(),
			CreatedAt:     time.Now().UTC(),
			UpdatedAt:     time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, policyAttachment)
	if err != nil {
		return apperrors.NewInternalServerError("")
	}
	return nil
}

func (s *PolicyAttachmentService) Update(ctx context.Context, request entities.UpdatePolicyAttachmentRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	return nil
}

func (s *PolicyAttachmentService) GetAllByPolicyId(ctx context.Context, policyId int) ([]entities.PolicyAttachment, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No attachments found")
	}
	attachments, err := s.repository.FindAllByPolicyId(ctx, policyId)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return attachments, nil
}

func (s *PolicyAttachmentService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.PolicyAttachment, error) {
	policyAttachment, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.PolicyAttachment{}, apperrors.NewNotFoundError("policy attachment not found")
	}
	return policyAttachment, nil
}

func (s *PolicyAttachmentService) Remove(ctx context.Context, uniqueId string) error {
	err := s.repository.DeleteByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing policy attachment: "+err.Error())
	}
	return nil
}