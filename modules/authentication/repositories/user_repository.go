package repositories

import (
	"context"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
	LastInsertId int64
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Create(ctx context.Context, user entities.User) error {
	result := repo.db.WithContext(ctx).Create(&user)
	repo.LastInsertId = user.UserId
	return result.Error
}

func (repo *UserRepository) Update(ctx context.Context, user entities.User) error {
	result := repo.db.WithContext(ctx).Save(&user)
	return result.Error
}

func (repo *UserRepository) FindAll(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	result := repo.db.WithContext(ctx).Find(&users)
	return users, result.Error
}

func (repo *UserRepository) FindById(ctx context.Context, id int) (entities.User, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).First(&user, id)
	return user, result.Error
}

func (repo *UserRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.User, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).First(&user, "unique_id=?", uniqueId)
	return user, result.Error
}

func (repo *UserRepository) FindByUserName(ctx context.Context, userName string) (entities.User, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).First(&user, "user_name=?", userName)
	return user, result.Error
}

func (repo *UserRepository) FindByToken(ctx context.Context, token string) (entities.User, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).First(&user, "token=?", token)
	return user, result.Error
}

func (repo *UserRepository) Search(ctx context.Context, param interface{}) ([]entities.UserData, error) {
	var users []entities.UserData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM authentication.view_user_data WHERE user_name=? OR role_name=?", param, param).Scan(&users)
	return users, result.Error
}

func (repo *UserRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.users").Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (repo *UserRepository) FindAllOrdered(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	result := repo.db.WithContext(ctx).Order("user_name ASC").Find(&users)
	return users, result.Error
}

func (repo *UserRepository) GetDataById(ctx context.Context, id int64) (entities.UserData, error) {
	var userData entities.UserData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM authentication.view_user_data WHERE user_id=?", id).Scan(&userData)
	return userData, result.Error
}

func (repo *UserRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.UserData, error) {
	var userData entities.UserData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM authentication.view_user_data WHERE unique_id=?", uniqueId).Scan(&userData)
	return userData, result.Error
}

func (repo *UserRepository) FindAllData(ctx context.Context) ([]entities.UserData, error) {
	var users []entities.UserData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM authentication.view_user_data").Scan(&users)
	return users, result.Error
}

func (repo *UserRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.UserData, error) {
	var users []entities.UserData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM authentication.view_user_data LIMIT ? OFFSET ?", limit, offset).Scan(&users)
	return users, result.Error
}

func (repo *UserRepository) Exists(ctx context.Context, userName string, password string) (bool, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).Where("user_name=? AND password=?", userName, password).Find(&user)
	return user.UserId !=0 , result.Error
}

func (repo *UserRepository) ExistsActive(ctx context.Context, userName string, password string) (bool, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).Where("user_name=? AND password=? AND active='Yes'", userName, password).Find(&user)
	return user.UserId !=0 , result.Error
}

func (repo *UserRepository) ExistsActiveUser(ctx context.Context, userName string) (bool, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).Where("user_name=? AND active='Yes'", userName).Find(&user)
	return user.UserId !=0 , result.Error
}

func (repo *UserRepository) GetByUserNameAndPassword(ctx context.Context, userName string, password string) (entities.UserData, error) {
	var userData entities.UserData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM authentication.view_user_data WHERE user_name=? AND password=?", userName, password).Scan(&userData)
	return userData, result.Error
}

func (repo *UserRepository) GetDataByUserName(ctx context.Context, userName string) (entities.UserData, error) {
	var userData entities.UserData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM authentication.view_user_data WHERE user_name=?", userName).Scan(&userData)
	return userData, result.Error
}

func (repo *UserRepository) FindAllByStatus(ctx context.Context, statusName string) ([]entities.UserData, error) {
	var users []entities.UserData
	repo.db.WithContext(ctx).Raw("SELECT * FROM authentication.view_user_data WHERE active=?",  statusName).Find(&users)
	return users, nil
}

func (repo *UserRepository) FindAllByRole(ctx context.Context, roleName string) ([]entities.UserData, error) {
	var users []entities.UserData
	repo.db.WithContext(ctx).Raw("SELECT * FROM authentication.view_user_data WHERE role_name=?",  roleName).Find(&users)
	return users, nil
}

func (repo *UserRepository) FindInactiveByRole(ctx context.Context, roleName string) ([]entities.UserData, error) {
	var users []entities.UserData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM authentication.view_user_data WHERE role_name=? and active='No'",  roleName).Find(&users)
	return users, result.Error
}