package models

import (
	"gorm.io/gorm"
	"github.com/ortizdavid/golang-modular-software/config"
	"github.com/ortizdavid/golang-modular-software/modules/users/entities"
)

type RoleModel struct {
}


func (RoleModel) Create(role entities.Role) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Create(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (RoleModel) FindAll() ([]entities.Role, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	roles := []entities.Role{}
	result := db.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

func (RoleModel) Update(role entities.Role) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Save(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (RoleModel) FindById(id int) (entities.Role, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var role entities.Role
	result := db.First(&role, id)
	if result.Error != nil {
		return entities.Role{}, result.Error
	}
	return role, nil
}

func (RoleModel) FindByName(name string) (entities.Role, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var role entities.Role
	result := db.Where("role_name=?", name).First(&role)
	if result.Error != nil {
		return entities.Role{}, result.Error
	}
	return role, nil
}

func (RoleModel) Count() (int64, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	result := db.Table("roles").Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
