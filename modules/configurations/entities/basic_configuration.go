package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type BasicConfiguration struct {
    ConfigurationId     int `gorm:"autoIncrement;primaryKey" json:"configuration_id"`
    AppName             string `gorm:"column:app_name" json:"app_name"`
    AppAcronym          string `gorm:"column:app_acronym" json:"app_acronym"`
    MaxRecordsPerPage   int `gorm:"column:max_records_per_page" json:"max_records_per_page"`
    MaxAdmninUsers      int `gorm:"column:max_admin_users" json:"max_admin_users"`
    MaxSuperAdmninUsers int `gorm:"column:max_super_admin_users" json:"max_super_admin_users"`
    shared.BaseEntity
}

func (BasicConfiguration) TableName() string {
	return "configurations.basic_configuration"
}
