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

func (s *RoomService) CreateRoom(ctx context.Context, request entities.CreateRoomRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	company, err := s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NewNotFoundError("company not found")
	}
	exists, err := s.repository.ExistsByName(ctx, company.CompanyId, request.RoomName)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewBadRequestError("Room already exists for company " + company.CompanyName)
	}
	branch, err := s.branchRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NewNotFoundError("branch not found")
	}
	room := entities.Room{
		CompanyId: company.CompanyId,
		BranchId:  branch.BranchId,
		RoomName:  request.RoomName,
		Number:    request.Number,
		Capacity:  request.Capacity,
		UniqueId:  encryption.GenerateUUID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	err = s.repository.Create(ctx, room)
	if err != nil {
		return apperrors.NewInternalServerError("error while creating room: " + err.Error())
	}
	return nil
}

func (s *RoomService) UpdateRoom(ctx context.Context, roomId int, request entities.UpdateRoomRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	room, err := s.repository.FindById(ctx, roomId)
	if err != nil {
		return apperrors.NewNotFoundError("room not found")
	}
	_, err = s.companyRepository.FindById(ctx, request.CompanyId)
	if err != nil {
		return apperrors.NewNotFoundError("company not found")
	}
	room.CompanyId = request.CompanyId
	room.RoomName = request.RoomName
	room.BranchId = request.BranchId
	room.Number = request.Number
	room.Capacity = request.Capacity
	room.UpdatedAt = time.Now().UTC()
	err = s.repository.Update(ctx, room)
	if err != nil {
		return apperrors.NewInternalServerError("error while updating room: " + err.Error())
	}
	return nil
}

func (s *RoomService) GetAllCompaniesPaginated(ctx context.Context, fiberCtx *fiber.Ctx, params helpers.PaginationParam) (*helpers.Pagination[entities.RoomData], error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No rooms found")
	}
	rooms, err := s.repository.FindAllLimit(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, rooms, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *RoomService) GetAllRooms(ctx context.Context) ([]entities.Room, error) {
	_, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No rooms found")
	}
	rooms, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	return rooms, nil
}

func (s *RoomService) SearchRooms(ctx context.Context, fiberCtx *fiber.Ctx, request entities.SearchRoomRequest, paginationParams helpers.PaginationParam) (*helpers.Pagination[entities.RoomData], error) {
	count, err := s.repository.CountByParam(ctx, request.SearchParam)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No rooms found")
	}
	if err := paginationParams.Validate(); err != nil {
		return nil, apperrors.NewBadRequestError(err.Error())
	}
	rooms, err := s.repository.Search(ctx, request.SearchParam, paginationParams.Limit, paginationParams.CurrentPage)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error fetching rows: " + err.Error())
	}
	pagination, err := helpers.NewPagination(fiberCtx, rooms, count, paginationParams.CurrentPage, paginationParams.Limit)
	if err != nil {
		return nil, apperrors.NewInternalServerError("Error creating pagination: " + err.Error())
	}
	return pagination, nil
}

func (s *RoomService) GetRoomByUniqueId(ctx context.Context, uniqueId string) (entities.RoomData, error) {
	room, err := s.repository.GetDataByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.RoomData{}, apperrors.NewNotFoundError("room not found")
	}
	return room, nil
}
