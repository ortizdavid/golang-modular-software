package models

import (
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"gorm.io/gorm"
)


type EmailConfigurationModel struct {
}


func (EmailConfigurationModel) FindById(id int) (entities.EmailConfiguration, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var confuracao entities.EmailConfiguration
	result := db.First(&confuracao, id)
	if result.Error != nil {
		return entities.EmailConfiguration{}, result.Error
	}
	return confuracao, nil
}


func (EmailConfigurationModel) Update(configuration entities.EmailConfiguration) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Save(&configuration)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}


func (configurationModel EmailConfigurationModel) FindFirst() (entities.EmailConfiguration, error) {
	configuration, err := configurationModel.FindById(1)
	if err != nil {
		return entities.EmailConfiguration{}, nil
	}
	return configuration, nil
}