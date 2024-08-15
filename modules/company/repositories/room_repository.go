package repositories

import "github.com/ortizdavid/golang-modular-software/database"

type RoomRepository struct {
	db *database.Database
}

func NewRoomRepository(db *database.Database) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}