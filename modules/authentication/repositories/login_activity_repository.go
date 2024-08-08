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

func (repo *LoginActivityRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.login_activity").Count(&count)
	return count, result.Error
}

func (repo *LoginActivityRepository) CountByStatus(ctx context.Context, status string) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.login_activity").Where("status = ?", status).Count(&count)
	return count, result.Error
}

func (repo *LoginActivityRepository) SumLoginAndLogout(ctx context.Context) (int64, int64, error) {
	var result struct {
		SumLogin  int64 `json:"sum_login"`
		SumLogout int64 `json:"sum_logout"`
	}
	query := `
		SELECT 
			COALESCE(SUM(total_login), 0) AS sum_login, 
			COALESCE(SUM(total_logout), 0) AS sum_logout 
		FROM authentication.login_activity
	`
	err := repo.db.WithContext(ctx).Raw(query).Scan(&result).Error
	return result.SumLogin, result.SumLogout, err
}
