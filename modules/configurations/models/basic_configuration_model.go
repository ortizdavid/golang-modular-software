package models

import (
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"gorm.io/gorm"
)

type BasicConfigurationModel struct {
}


func (BasicConfigurationModel) FindById(id int) (entities.BasicConfiguration, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var configuration entities.BasicConfiguration
	result := db.First(&configuration, id)
	if result.Error != nil {
		return entities.BasicConfiguration{}, result.Error
	}
	return configuration, nil
}


func (BasicConfigurationModel) Update(configuration entities.BasicConfiguration) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Save(&configuration)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}


func (configurationModel BasicConfigurationModel) FindFirst() (entities.BasicConfiguration, error) {
	configuration, err := configurationModel.FindById(1)
	if err != nil {
		return entities.BasicConfiguration{}, err
	}
	if configuration.MaxRecordsPerPage < 5 {
		configuration.MaxRecordsPerPage = 5
		configurationModel.Update(configuration)
	}
	return configuration, nil
}