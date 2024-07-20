package models

import (
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
	"gorm.io/gorm"
)

type CompanyConfigurationModel struct {
}


func (CompanyConfigurationModel) FindById(id int) (entities.CompanyConfiguration, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var configuration entities.CompanyConfiguration
	result := db.First(&configuration, id)
	if result.Error != nil {
		return entities.CompanyConfiguration{}, result.Error
	}
	return configuration, nil
}


func (CompanyConfigurationModel) Update(configuration entities.CompanyConfiguration) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Save(&configuration)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}


func (configurationModel CompanyConfigurationModel) FindFirst() (entities.CompanyConfiguration, error) {
	configuration, err := configurationModel.FindById(1)
	if err != nil {
		return entities.CompanyConfiguration{}, err
	}
	return configuration, nil
}