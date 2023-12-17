package models

import (
	"github.com/ortizdavid/golang-modular-software/config"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"gorm.io/gorm"
)

type UserModel struct {
	LastInsertId int
}

func (userModel *UserModel) Create(user entities.User) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	userModel.LastInsertId = user.UserId
	return result, nil
}

func (UserModel) FindAll() ([]entities.User, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	users := []entities.User{}
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (UserModel) Update(user entities.User) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (UserModel) FindById(id int) (entities.User, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.First(&user, id)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

func (UserModel) FindByUniqueId(uniqueId string) (entities.User, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.First(&user, "unique_id=?", uniqueId)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

func (UserModel) FindByUserName(userName string) (entities.User, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.First(&user, "user_name=?", userName)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

func (UserModel) FindByToken(token string) (entities.User, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.First(&user, "token=?", token)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}


func (UserModel) Search(param interface{}) ([]entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var users []entities.UserData
	result := db.Raw("SELECT * FROM view_user_data WHERE user_name=? OR role_name=?", param, param).Scan(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (UserModel) Count() (int64, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	result := db.Table("users").Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (UserModel) FindAllOrdered() ([]entities.User, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	users := []entities.User{}
	result := db.Order("user_name ASC").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (UserModel) GetDataById(id int) (entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var userData entities.UserData
	result := db.Raw("SELECT * FROM view_user_data WHERE user_id=?", id).Scan(&userData)
	if result.Error != nil {
		return entities.UserData{}, result.Error
	}
	return userData, nil
}

func (UserModel) GetDataByUniqueId(uniqueId string) (entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var userData entities.UserData
	result := db.Raw("SELECT * FROM view_user_data WHERE unique_id=?", uniqueId).Scan(&userData)
	if result.Error != nil {
		return entities.UserData{}, result.Error
	}
	return userData, nil
}

func (UserModel) FindAllData() ([]entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var users []entities.UserData
	result := db.Raw("SELECT * FROM view_user_data").Scan(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (UserModel) FindAllDataLimit(start int, end int) ([]entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var users []entities.UserData
	result := db.Raw("SELECT * FROM view_user_data LIMIT ?, ?", start, end).Scan(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (UserModel) Exists(userName string, password string) (bool, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.Where("user_name=? AND password=?", userName, password).Find(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return user.UserId !=0 , nil
}

func (UserModel) ExistsActive(userName string, password string) (bool, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.Where("user_name=? AND password=? AND active='Yes'", userName, password).Find(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return user.UserId !=0 , nil
}

func (UserModel) ExistsActiveUser(userName string) (bool, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.Where("user_name=? AND active='Yes'", userName).Find(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return user.UserId !=0 , nil
}

func (UserModel) GetByUserNameAndPassword(userName string, password string) (entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var userData entities.UserData
	result := db.Raw("SELECT * FROM view_user_data WHERE user_name=? AND password=?", userName, password).Scan(&userData)
	if result.Error != nil {
		return entities.UserData{}, result.Error
	}
	return userData, nil
}

func (UserModel) GetDataByUserName(userName string) (entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var userData entities.UserData
	result := db.Raw("SELECT * FROM view_user_data WHERE user_name=?", userName).Scan(&userData)
	if result.Error != nil {
		return entities.UserData{}, result.Error
	}
	return userData, nil
}

func (UserModel) FindAllByStatus(statusName string) ([]entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	users := []entities.UserData{}
	db.Raw("SELECT * FROM view_user_data WHERE active=?",  statusName).Find(&users)
	return users, nil
}

func (UserModel) FindAllByRole(roleName string) ([]entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	users := []entities.UserData{}
	db.Raw("SELECT * FROM view_user_data WHERE role_name=?",  roleName).Find(&users)
	return users, nil
}

func (UserModel) FindInactiveByRole(roleName string) ([]entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var users []entities.UserData
	result := db.Raw("SELECT * FROM view_user_data WHERE role_name=? and active='No'",  roleName).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}