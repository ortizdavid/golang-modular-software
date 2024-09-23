package repositories

import (
	"context"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"

)

type UserRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.User]
}

func NewUserRepository(db *database.Database) *UserRepository {
	return &UserRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.User](db),
	}
}

func (repo *UserRepository) Create(ctx context.Context, user entities.User) error {
	result := repo.db.WithContext(ctx).Create(&user)
	repo.SetLastInsertId(user.UserId)
	return result.Error
}

func (repo *UserRepository) FindAllDataLimit(ctx context.Context, limit int, offset int) ([]entities.UserData, error) {
	var users []entities.UserData
	result := repo.db.WithContext(ctx).
		Table("authentication.view_user_data").
		Limit(limit).
		Offset(offset).Find(&users)
	return users, result.Error
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

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (entities.User, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).First(&user, "email=?", email)
	return user, result.Error
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

func (repo *UserRepository) Exists(ctx context.Context, userName string, password string) (bool, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).Where("user_name=? AND password=?", userName, password).Find(&user)
	return user.UserId !=0 , result.Error
}

func (repo *UserRepository) ExistsByName(ctx context.Context, userName string) (bool, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).Where("user_name=?", userName).Find(&user)
	return user.UserId !=0 , result.Error
}

func (repo *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).Where("email=?", email).Find(&user)
	return user.UserId !=0 , result.Error
}

func (repo *UserRepository) ExistsActive(ctx context.Context, userName string, password string) (bool, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).Where("user_name=? AND password=? AND id_active=true", userName, password).Find(&user)
	return user.UserId !=0 , result.Error
}

func (repo *UserRepository) ExistsActiveUser(ctx context.Context, userName string) (bool, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).Where("user_name=? AND is_active=true", userName).Find(&user)
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

func (repo *UserRepository) GetDataByEmail(ctx context.Context, email string) (entities.UserData, error) {
	var userData entities.UserData
	result := repo.db.WithContext(ctx).Raw("SELECT * FROM authentication.view_user_data WHERE email=?", email).Scan(&userData)
	return userData, result.Error
}

func (repo *UserRepository) FindAllByStatus(ctx context.Context, status string, limit int, offset int) ([]entities.UserData, error) {
	var users []entities.UserData
	result := repo.db.WithContext(ctx).
		Table("authentication.view_user_data").
		Where("is_active=?", status).
		Limit(limit).
		Offset(offset).Find(&users)
	return users, result.Error
}

func (repo *UserRepository) FindAllByActivityStatus(ctx context.Context, status entities.LoginActivityStatus, limit int, offset int) ([]entities.UserData, error) {
	var users []entities.UserData
	result := repo.db.WithContext(ctx).
		Table("authentication.view_user_data").
		Where("status=?", status).
		Limit(limit).
		Offset(offset).Find(&users)
	return users, result.Error
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

func (repo *UserRepository) HasRoles(ctx context.Context, userId int64, roles ...string) (bool, error) {
    var count int64
    result := repo.db.WithContext(ctx).
        Table("authentication.view_user_role_data").
        Where("user_id = ? AND role_code IN (?)  AND role_status = 'Enabled'", userId, roles).
        Count(&count)
    if result.Error != nil {
        return false, result.Error
    }
    return count > 0, nil
}

func (repo *UserRepository) CountByStatus(ctx context.Context, status bool) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.users").Where("is_active = ?", status).Count(&count)
	return count, result.Error
}

func (repo *UserRepository) CountByActivityStatus(ctx context.Context, status entities.LoginActivityStatus) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("authentication.view_user_data").Where("status = ?", status).Count(&count)
	return count, result.Error
}

func (repo *UserRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.UserData, error) {
	var users []entities.UserData
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM authentication.view_user_data WHERE user_name=? OR email=?", param, param).
		Limit(limit).
		Offset(offset).
		Scan(&users)
	return users, result.Error
}

func (repo *UserRepository) CountByParam(ctx context.Context, param string) (int64, error) {
    var count int64
    result := repo.db.WithContext(ctx).
        Raw("SELECT COUNT(*) FROM authentication.view_user_data WHERE user_name=? OR email=?", param, param).
        Scan(&count)
    return count, result.Error
}

func (repo *UserRepository) FindUsersWithoutAssociation(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	result := repo.db.WithContext(ctx).
		Table("authentication.users").
		Where("user_id NOT IN(SELECT user_id FROM authentication.user_associations)").
		Order("created_at DESC").
		Find(&users)
	return users, result.Error
}
