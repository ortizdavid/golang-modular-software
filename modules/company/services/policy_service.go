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
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type PolicyService struct {
	repository *repositories.PolicyRepository
	companyRepository *repositories.CompanyRepository
}

func NewPolicyService(db *database.Database) *PolicyService {
	return &PolicyService{
		repository: repositories.NewPolicyRepository(db),
		companyRepository: repositories.NewCompanyRepository(db),
	}
}

func (s *PolicyService) Create(ctx context.Context, request entities.CreatePolicyRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	effectiveDate, err := datetime.StringToDate(request.EffectiveDate)
	if err != nil {
		return err
	}
	company, err := s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NewNotFoundError("company not found")
	}
	exists, err := s.repository.ExistsByName(ctx, company.CompanyId, request.PolicyName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("Policy already exists for company " + company.CompanyName)
	}
	policy := entities.Policy{
		CompanyId:     company.CompanyId,
		PolicyName:    request.PolicyName,
		Description:   request.Description,
		EffectiveDate: effectiveDate,
		BaseEntity: shared.BaseEntity{
			UniqueId:      encryption.GenerateUUID(),
			CreatedAt:     time.Now().UTC(),
			UpdatedAt:     time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, policy)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating policy: " + err.Error())
	}
	return nil
}

func (s *PolicyService) Update(ctx context.Context, policyId int, request entities.UpdatePolicyRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	effectiveDate, err := datetime.StringToDate(request.EffectiveDate)
	if err != nil {
		return err
	}
	policy, err := s.repository.FindById(ctx, policyId)
	if err != nil {
		return apperrors.NewNotFoundError("policy not found")
	}
	_, err = s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NewNotFoundError("company not found")
	}
	policy.CompanyId = request.CompanyId
	policy.PolicyName = request.PolicyName
	policy.CompanyId = request.CompanyId
	policy.EffectiveDate = effectiveDate
	policy.Description = request.Description
	policy.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, policy)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating policy: " + err.Error())
	}
	return nil
}

func (s *PolicyService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.PolicyData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No policies found")
	}
	policies, err := s.repository.FindAllDataLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, policies, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *PolicyService) GetAll(ctx context.Context) ([]entities.Policy, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No policies found")
	}
	policies, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return policies, nil
}

func (s *PolicyService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchPolicyRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.PolicyData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No policies found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	policies, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, policies, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *PolicyService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.PolicyData, error) {
	policy, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.PolicyData{}, apperrors.NewNotFoundError("policy not found")
	}
	return policy, nil
}

func (s *PolicyService) GetPolicyById(ctx context.Context, policyId int) (entities.Policy, error) {
	policy, err := s.repository.FindById(ctx, policyId)
	if err != nil {
		return entities.Policy{}, apperrors.NewNotFoundError("policy not found")
	}
	return policy, nil
}

func (s *PolicyService) Remove(ctx context.Context, uniqueId string) error {
	err := s.repository.DeleteByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing policy: "+ err.Error())
	}
	return nil
}
