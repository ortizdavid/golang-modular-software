package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type LoginActivityRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.LoginActivity]
}

func NewLoginActivityRepository(db *database.Database) *LoginActivityRepository {
	return &LoginActivityRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.LoginActivity](db),
	}
}

func (repo *LoginActivityRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.LoginActivityData, error) {
	var loginActivities []entities.LoginActivityData
	result := repo.db.WithContext(ctx).Table("authentication.view_login_activity_data").Limit(limit).Offset(offset).Find(&loginActivities)
	return loginActivities, result.Error
}

func (repo *LoginActivityRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.LoginActivityData, error) {
	var loginActivityData entities.LoginActivityData
	result := repo.db.WithContext(ctx).Table("authentication.view_login_activity_data").Where("unique_id=?", uniqueId).First(&loginActivityData)
	return loginActivityData, result.Error
}

func (repo *LoginActivityRepository) GetDataByUserId(ctx context.Context, userId int64) (entities.LoginActivityData, error) {
	var loginActivityData entities.LoginActivityData
	result := repo.db.WithContext(ctx).Table("authentication.view_login_activity_data").Where("user_id=?", userId).First(&loginActivityData)
	return loginActivityData, result.Error
}

func (repo *LoginActivityRepository) GetDataByUserName(ctx context.Context, userName string) (entities.LoginActivityData, error) {
	var loginActivityData entities.LoginActivityData
	result := repo.db.WithContext(ctx).Table("authentication.view_login_activity_data").Where("user_name=?", userName).First(&loginActivityData)
	return loginActivityData, result.Error
}

func (repo *LoginActivityRepository) GetDataByEmail(ctx context.Context, email string) (entities.LoginActivityData, error) {
	var loginActivityData entities.LoginActivityData
	result := repo.db.WithContext(ctx).Table("authentication.view_login_activity_data").Where("email=?", email).First(&loginActivityData)
	return loginActivityData, result.Error
}

func (repo *LoginActivityRepository) FindByUserId(ctx context.Context, userId int64) (entities.LoginActivity, error) {
	var loginAct entities.LoginActivity
	result := repo.db.WithContext(ctx).Where("user_id=?", userId).First(&loginAct)
	return loginAct, result.Error
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
		FROM authentication.login_activity`
	err := repo.db.WithContext(ctx).Raw(query).Scan(&result).Error
	return result.SumLogin, result.SumLogout, err
}

func (repo *LoginActivityRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.LoginActivityData, error) {
	var loginActivities []entities.LoginActivityData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM authentication.view_login_activity_data WHERE user_name LIKE ? OR email LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&loginActivities)
	return loginActivities, result.Error
}

func (repo *LoginActivityRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM authentication.view_login_activity_data WHERE user_name LIKE ? OR email LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}
