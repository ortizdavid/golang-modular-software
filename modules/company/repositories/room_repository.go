package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type RoomRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.Room]
}

func NewRoomRepository(db *database.Database) *RoomRepository {
	return &RoomRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.Room](db),
	}
}

func (repo *RoomRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.RoomData, error) {
	var rooms []entities.RoomData
	result := repo.db.WithContext(ctx).Table("company.view_room_data").Limit(limit).Offset(offset).Find(&rooms)
	return rooms, result.Error
}

func (repo *RoomRepository) FindAllData(ctx context.Context) ([]entities.RoomData, error) {
	var rooms []entities.RoomData
	result := repo.db.WithContext(ctx).Table("company.view_room_data").Find(&rooms)
	return rooms, result.Error
}

func (repo *RoomRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.RoomData, error) {
	var room entities.RoomData
	result := repo.db.WithContext(ctx).Table("company.view_room_data").Where("unique_id=?", uniqueId).First(&room)
	return room, result.Error
}

func (repo *RoomRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.RoomData, error) {
	var rooms []entities.RoomData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM company.view_room_data WHERE room_name LIKE ? OR acronym LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&rooms)
	return rooms, result.Error
}

func (repo *RoomRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM company.view_room_data WHERE room_name LIKE ? OR acronym LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *RoomRepository) ExistsByName(ctx context.Context, companyId int, roomName string) (bool, error) {
	var room entities.Room
	result := repo.db.WithContext(ctx).Where("company_id=? AND room_name=?", companyId, roomName).Find(&room)
	return room.RoomId != 0, result.Error
}
