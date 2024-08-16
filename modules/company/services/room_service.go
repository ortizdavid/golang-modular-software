package services

import (
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
)

type RoomService struct {
	repository *repositories.RoomRepository
}

func NewRoomService(db *database.Database) *RoomService {
	return &RoomService{
		repository: repositories.NewRoomRepository(db),
	}
}