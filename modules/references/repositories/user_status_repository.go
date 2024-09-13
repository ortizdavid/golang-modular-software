package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/references/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type UserStatusRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.UserStatus]
}

func NewUserStatusRepository(db *database.Database) *UserStatusRepository {
	return &UserStatusRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.UserStatus](db),
	}
}

func (repo *UserStatusRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.UserStatus, error) {
	var userStatus entities.UserStatus
	result := repo.db.WithContext(ctx).Table("reference.user_statuses").Where("unique_id=?", uniqueId).First(&userStatus)
	return userStatus, result.Error
}

func (repo *UserStatusRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.UserStatus, error) {
	var userStatuses []entities.UserStatus
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM reference.user_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&userStatuses)
	return userStatuses, result.Error
}

func (repo *UserStatusRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM reference.user_statuses WHERE status_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *UserStatusRepository) ExistsByName(ctx context.Context, userStatusName string) (bool, error) {
	var userStatus entities.UserStatus
	result := repo.db.WithContext(ctx).Where("status_name=?", userStatusName).Find(&userStatus)
	return userStatus.StatusId != 0, result.Error
}
