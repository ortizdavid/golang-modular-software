package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type BranchService struct {
	repository        *repositories.BranchRepository
	companyRepository *repositories.CompanyRepository
}

func NewBranchService(db *database.Database) *BranchService {
	return &BranchService{
		repository:        repositories.NewBranchRepository(db),
		companyRepository: repositories.NewCompanyRepository(db),
	}
}

func (s *BranchService) CreateBranch(ctx context.Context, request entities.CreateBranchRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	company, err := s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NewNotFoundError("company not found")
	}
	exists, err := s.repository.ExistsByName(ctx, company.CompanyId, request.BranchName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("Branch already exists for company " + company.CompanyName)
	}
	branch := entities.Branch{
		BranchId:   0,
		CompanyId:  company.CompanyId,
		BranchName: request.BranchName,
		Code:       request.Code,
		Address:    request.Address,
		Phone:      request.Phone,
		Email:      request.Email,
		BaseEntity: shared.BaseEntity{
			UniqueId:   encryption.GenerateUUID(),
			CreatedAt:  time.Now().UTC(),
			UpdatedAt:  time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, branch)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating branch: " + err.Error())
	}
	return nil
}

func (s *BranchService) UpdateBranch(ctx context.Context, branchId int, request entities.UpdateBranchRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	branch, err := s.repository.FindById(ctx, branchId)
	if err != nil {
		return apperrors.NewNotFoundError("branch not found")
	}
	_, err = s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NewNotFoundError("company not found")
	}
	branch.CompanyId = request.CompanyId
	branch.BranchName = request.BranchName
	branch.Address = request.Address
	branch.Phone = request.Phone
	branch.Email = request.Email
	branch.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, branch)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating branch: " + err.Error())
	}
	return nil
}

func (s *BranchService) GetAllBranchesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.BranchData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No branches found")
	}
	branches, err := s.repository.FindAllDataLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, branches, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *BranchService) GetAllBranches(ctx context.Context) ([]entities.Branch, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No branchs found")
	}
	branchs, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return branchs, nil
}

func (s *BranchService) SearchBranches(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchBranchRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.BranchData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No branches found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	branchs, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, branchs, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *BranchService) GetBranchByUniqueId(ctx context.Context, uniqueId string) (entities.BranchData, error) {
	branch, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.BranchData{}, apperrors.NewNotFoundError("branch not found")
	}
	return branch, nil
}
