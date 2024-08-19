package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
)

type RoomRepository struct {
	db *database.Database
}

func NewRoomRepository(db *database.Database) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (repo *RoomRepository) Create(ctx context.Context, company entities.Room) error {
	result := repo.db.WithContext(ctx).Create(&company)
	return result.Error
}

func (repo *RoomRepository) Update(ctx context.Context, company entities.Room) error {
	result := repo.db.WithContext(ctx).Save(&company)
	return result.Error
}

func (repo *RoomRepository) Delete(ctx context.Context, company entities.Room) error {
	result := repo.db.WithContext(ctx).Delete(&company)
	return result.Error
}

func (repo *RoomRepository) FindAll(ctx context.Context) ([]entities.Room, error) {
	var rooms []entities.Room
	result := repo.db.WithContext(ctx).Find(&rooms)
	return rooms, result.Error
}

func (repo *RoomRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.RoomData, error) {
	var rooms []entities.RoomData
	result := repo.db.WithContext(ctx).Table("company.view_room_data").Limit(limit).Offset(offset).Find(&rooms)
	return rooms, result.Error
}

func (repo *RoomRepository) FindById(ctx context.Context, id int) (entities.Room, error) {
	var company entities.Room
	result := repo.db.WithContext(ctx).First(&company, id)
	return company, result.Error
}

func (repo *RoomRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Room, error) {
	var company entities.Room
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *RoomRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.RoomData, error) {
	var company entities.RoomData
	result := repo.db.WithContext(ctx).Table("company.view_room_data").Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *RoomRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("company.rooms").Count(&count)
	return count, result.Error
}

func (repo *RoomRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.RoomData, error) {
	var rooms []entities.RoomData
	likeParam := "%"+param+"%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM company.view_room_data WHERE room_name LIKE ? OR acronym LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&rooms)
	return rooms, result.Error
}

func (repo *RoomRepository) CountByParam(ctx context.Context, param string) (int64, error) {
    var count int64
	likeParam := "%"+param+"%"
    result := repo.db.WithContext(ctx).
        Raw("SELECT COUNT(*) FROM company.view_room_data WHERE room_name LIKE ? OR acronym LIKE ?", likeParam, likeParam).
        Scan(&count)
    return count, result.Error
}

func (repo *RoomRepository) ExistsByName(ctx context.Context, companyId int, roomName string) (bool, error) {
	var room entities.Room
	result := repo.db.WithContext(ctx).Where("company_id=? AND room_name=?", companyId, roomName).Find(&room)
	return room.RoomId !=0 , result.Error
}