package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
)

type LoginActivityRepository struct {
	db *database.Database
}

func NewLoginActivityRepository(db *database.Database) *LoginActivityRepository {
	return &LoginActivityRepository{
		db: db,
	}
}

func (repo *LoginActivityRepository) Create(ctx context.Context, loginAct entities.LoginActivity) error {
	result := repo.db.WithContext(ctx).Create(&loginAct)
	return result.Error
}

func (repo *LoginActivityRepository) Update(ctx context.Context, loginAct entities.LoginActivity) error {
	result := repo.db.WithContext(ctx).Save(&loginAct)
	return result.Error
}

func (repo *LoginActivityRepository) FindById(ctx context.Context, userId int64) (entities.LoginActivity, error) {
	var loginAct entities.LoginActivity
	result := repo.db.WithContext(ctx).First(&loginAct)
	return loginAct, result.Error
}

func (repo *LoginActivityRepository) FindByUserId(ctx context.Context, userId int64) (entities.LoginActivity, error) {
	var loginAct entities.LoginActivity
	result := repo.db.WithContext(ctx).Where("user_id=?", userId).First(&loginAct)
	return loginAct, result.Error
}