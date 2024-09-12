package repositories

import (
	"context"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

type UserApiKeyRepository struct {
	db *database.Database
}

func NewUserApiKeyRepository(db *database.Database) *UserApiKeyRepository {
	return &UserApiKeyRepository {
		db: db,
	}
}

func (repo *UserApiKeyRepository) Create(ctx context.Context, userApiKey entities.UserApiKey) error {
	result := repo.db.WithContext(ctx).Create(&userApiKey)
	return result.Error
}

func (repo *UserApiKeyRepository) Update(ctx context.Context, userApiKey entities.UserApiKey) error {
	result := repo.db.WithContext(ctx).Save(&userApiKey)
	return result.Error
}

func (repo *UserApiKeyRepository) Delete(ctx context.Context, userApiKey entities.UserApiKey) error {
	result := repo.db.WithContext(ctx).Delete(&userApiKey)
	return result.Error
}

func (repo *UserApiKeyRepository) FindAll(ctx context.Context) ([]entities.UserApiKey, error) {
	var userApiKeys []entities.UserApiKey
	result := repo.db.WithContext(ctx).Find(&userApiKeys)
	return userApiKeys, result.Error
}

func (repo *UserApiKeyRepository) FindAllDataByUserId(ctx context.Context, userId int64) ([]entities.UserApiKeyData, error) {
	var userApiKeys []entities.UserApiKeyData
	result := repo.db.WithContext(ctx).Table("authentication.view_user_api_key_data").Where("user_id = ?", userId).Find(&userApiKeys)
	return userApiKeys, result.Error
}

func (repo *UserApiKeyRepository) FindById(ctx context.Context, id int) (entities.UserApiKeyData, error) {
	var userApiKey entities.UserApiKeyData
	result := repo.db.WithContext(ctx).First(&userApiKey, id)
	return userApiKey, result.Error
}

func (repo *UserApiKeyRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.UserApiKey, error) {
	var userApiKey entities.UserApiKey
	result := repo.db.WithContext(ctx).Where("unique_id = ?", uniqueId).First(&userApiKey)
	return userApiKey, result.Error
}

func (repo *UserApiKeyRepository) FindByUserId(ctx context.Context, userId int64) (entities.UserApiKey, error) {
	var userApiKey entities.UserApiKey
	result := repo.db.WithContext(ctx).Where("user_id=?", userId).First(&userApiKey)
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

func (repo *UserApiKeyRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.user_api_key").Count(&count)
	return count, result.Error
}
