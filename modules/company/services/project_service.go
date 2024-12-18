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

type ProjectService struct {
	repository        *repositories.ProjectRepository
	companyRepository *repositories.CompanyRepository
}

func NewProjectService(db *database.Database) *ProjectService {
	return &ProjectService{
		repository:        repositories.NewProjectRepository(db),
		companyRepository: repositories.NewCompanyRepository(db),
	}
}

func (s *ProjectService) Create(ctx context.Context, request entities.CreateProjectRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	startDate, err := datetime.StringToDate(request.StartDate)
	if err != nil {
		return err
	}
	endDate, err := datetime.StringToDate(request.EndDate)
	if err != nil {
		return err
	}
	company, err := s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NotFoundError("company not found")
	}
	exists, err := s.repository.ExistsByName(ctx, company.CompanyId, request.ProjectName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.BadRequestError("Project already exists for company " + company.CompanyName)
	}
	project := entities.Project{
		ProjectName: request.ProjectName,
		Description: request.Description,
		StartDate:   startDate,
		EndDate:     endDate,
		Status:      request.Status,
		CompanyId:   company.CompanyId,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, project)
	if err != nil {
		return apperrors.InternalServerError("error while creating project: " + err.Error())
	}
	return nil
}

func (s *ProjectService) Update(ctx context.Context, uniqueId string, request entities.UpdateProjectRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	startDate, err := datetime.StringToDate(request.StartDate)
	if err != nil {
		return err
	}
	endDate, err := datetime.StringToDate(request.EndDate)
	if err != nil {
		return err
	}
	project, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NotFoundError("project not found")
	}
	_, err = s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NotFoundError("company not found")
	}
	project.CompanyId = request.CompanyId
	project.ProjectName = request.ProjectName
	project.StartDate = startDate
	project.EndDate = endDate
	project.Status = request.Status
	project.Description = request.Description
	project.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, project)
	if err != nil {
		return apperrors.InternalServerError("error while updating project: " + err.Error())
	}
	return nil
}

func (s *ProjectService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.ProjectData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No projects found")
	}
	projects, err := s.repository.FindAllDataLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, projects, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *ProjectService) GetAll(ctx context.Context) ([]entities.Project, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No projects found")
	}
	projects, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	return projects, nil
}

func (s *ProjectService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchProjectRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.ProjectData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NotFoundError("No projects found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	projects, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, projects, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *ProjectService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.ProjectData, error) {
	project, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.ProjectData{}, apperrors.NotFoundError("project not found")
	}
	return project, nil
}

func (s *ProjectService) GetById(ctx context.Context, projectId int) (entities.Project, error) {
	project, err := s.repository.FindById(ctx, projectId)
	if err != nil {
		return entities.Project{}, apperrors.NotFoundError("project not found")
	}
	return project, nil
}

func (s *ProjectService) Remove(ctx context.Context, uniqueId string) error {
	err := s.repository.DeleteByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.InternalServerError("error while removing project: " + err.Error())
	}
	return nil
}
