package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"

)

type UserApiKeyRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.UserApiKey]
}

func NewUserApiKeyRepository(db *database.Database) *UserApiKeyRepository {
	return &UserApiKeyRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.UserApiKey](db),
	}
}

func (repo *UserApiKeyRepository) FindAllDataByUserId(ctx context.Context, userId int64) ([]entities.UserApiKeyData, error) {
	var userApiKeys []entities.UserApiKeyData
	result := repo.db.WithContext(ctx).Table("authentication.view_user_api_key_data").Where("user_id = ?", userId).Find(&userApiKeys)
	return userApiKeys, result.Error
}

func (repo *UserApiKeyRepository) FindByUserId(ctx context.Context, userId int64) (entities.UserApiKey, error) {
	var userApiKey entities.UserApiKey
	result := repo.db.WithContext(ctx).Where("user_id=?", userId).First(&userApiKey)
	return userApiKey, result.Error
}

func (repo *UserApiKeyRepository) FindByXUserId(ctx context.Context, xUserId string) (entities.UserApiKey, error) {
	var userApiKey entities.UserApiKey
	result := repo.db.WithContext(ctx).Where("x_user_id=?", xUserId).First(&userApiKey)
	return userApiKey, result.Error
}

func (repo *UserApiKeyRepository) GetDataById(ctx context.Context, userApiKeyId int) (entities.UserApiKeyData, error) {
	var userApiKey entities.UserApiKeyData
	result := repo.db.WithContext(ctx).Table("authentication.view_user_api_key_data").Where("user_role_id=?", userApiKeyId).First(&userApiKey)
	return userApiKey, result.Error
}

func (repo *UserApiKeyRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.UserApiKeyData, error) {
	var userApiKey entities.UserApiKeyData
	result := repo.db.WithContext(ctx).Table("authentication.view_user_api_key_data").Where("unique_id=?", uniqueId).First(&userApiKey)
	return userApiKey, result.Error
}
