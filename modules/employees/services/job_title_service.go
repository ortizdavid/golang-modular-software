package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	"github.com/ortizdavid/golang-modular-software/modules/employees/repositories"
)

type JobTitleService struct {
	repository *repositories.JobTitleRepository
}

func NewJobTitleService(db *database.Database) *JobTitleService {
	return &JobTitleService{
		repository: repositories.NewJobTitleRepository(db),
	}
}


func (s *JobTitleService) CreateJobTitle(ctx context.Context, request entities.CreateJobTitleRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	exists, err := s.repository.ExistsByName(ctx, request.TitleName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("job title already exists")
	}
	jobTitle := entities.JobTitle{
		TitleName:   request.TitleName,
		Description: request.Description,
		UniqueId:    encryption.GenerateUUID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	err = s.repository.Create(ctx, jobTitle)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating job title: " + err.Error())
	}
	return nil
}

func (s *JobTitleService) UpdateJobTitle(ctx context.Context, jobTitleId int, request entities.UpdateJobTitleRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	jobTitle, err := s.repository.FindById(ctx, jobTitleId)
	if err != nil {
		return apperrors.NewNotFoundError("job title not found")
	}
	jobTitle.TitleName = request.TitleName
	jobTitle.Description = request.Description
	jobTitle.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, jobTitle)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating job title: " + err.Error())
	}
	return nil
}

func (s *JobTitleService) GetAllJobTitlesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.JobTitle], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No job titles found")
	}
	jobTitles, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, jobTitles, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *JobTitleService) GetAllJobTitles(ctx context.Context) ([]entities.JobTitle, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No job titles found")
	}
	jobTitles, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return jobTitles, nil
}

func (s *JobTitleService) SearchJobTitles(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchJobTitleRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.JobTitle], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No job titles found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	jobTitles, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, jobTitles, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *JobTitleService) GetJobTitleByUniqueId(ctx context.Context, uniqueId string) (entities.JobTitle, error) {
	jobTitle, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.JobTitle{}, apperrors.NewNotFoundError("job title not found")
	}
	return jobTitle, nil
}

func (s *JobTitleService) RemoveJobTitle(ctx context.Context, uniqueId string) error {
	jobTitle, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("job title not found")
	}
	err = s.repository.Delete(ctx, jobTitle)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing job title: "+ err.Error())
	}
	return nil
}