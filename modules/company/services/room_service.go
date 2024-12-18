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

type RoomService struct {
	repository        *repositories.RoomRepository
	companyRepository *repositories.CompanyRepository
	branchRepository  *repositories.BranchRepository
}

func NewRoomService(db *database.Database) *RoomService {
	return &RoomService{
		repository:        repositories.NewRoomRepository(db),
		companyRepository: repositories.NewCompanyRepository(db),
		branchRepository:  repositories.NewBranchRepository(db),
	}
}

func (s *RoomService) Create(ctx context.Context, request entities.CreateRoomRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	company, err := s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NotFoundError("company not found")
	}
	exists, err := s.repository.ExistsByName(ctx, company.CompanyId, request.RoomName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.BadRequestError("Room already exists for company " + company.CompanyName)
	}
	branch, err := s.branchRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NotFoundError("branch not found")
	}
	room := entities.Room{
		CompanyId: company.CompanyId,
		BranchId:  branch.BranchId,
		RoomName:  request.RoomName,
		Number:    request.Number,
		Capacity:  request.Capacity,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
	err = s.repository.Create(ctx, room)
	if err != nil {
		return apperrors.InternalServerError("error while creating room: " + err.Error())
	}
	return nil
}

func (s *RoomService) Update(ctx context.Context, uniqueId string, request entities.UpdateRoomRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.BadRequestError(err.Error())
	}
	room, err := s.repository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NotFoundError("room not found")
	}
	_, err = s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NotFoundError("company not found")
	}
	room.CompanyId = request.CompanyId
	room.RoomName = request.RoomName
	room.BranchId = request.BranchId
	room.Number = request.Number
	room.Capacity = request.Capacity
	room.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, room)
	if err != nil {
		return apperrors.InternalServerError("error while updating room: " + err.Error())
	}
	return nil
}

func (s *RoomService) GetAllPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.RoomData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No rooms found")
	}
	rooms, err := s.repository.FindAllDataLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, rooms, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *RoomService) GetAll(ctx context.Context) ([]entities.RoomData, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NotFoundError("No rooms found")
	}
	rooms, err := s.repository.FindAllData(ctx)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	return rooms, nil
}

func (s *RoomService) Search(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchRoomRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.RoomData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NotFoundError("No rooms found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.BadRequestError(err.Error())
	}
	rooms, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.InternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, rooms, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.InternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *RoomService) GetByUniqueId(ctx context.Context, uniqueId string) (entities.RoomData, error) {
	room, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.RoomData{}, apperrors.NotFoundError("room not found")
	}
	return room, nil
}
