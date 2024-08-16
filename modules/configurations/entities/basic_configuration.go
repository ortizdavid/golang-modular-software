package entities

import "time"

type BasicConfiguration struct {
    ConfigurationId     int `gorm:"autoIncrement;primaryKey" json:"configuration_id"`
    AppName             string `gorm:"column:app_name" json:"app_name"`
    AppAcronym          string `gorm:"column:app_acronym" json:"app_acronym"`
    MaxRecordsPerPage   int `gorm:"column:max_records_per_page" json:"max_records_per_page"`
    MaxAdmninUsers      int `gorm:"column:max_admin_users" json:"max_admin_users"`
    MaxSuperAdmninUsers int `gorm:"column:max_super_admin_users" json:"max_super_admin_users"`
    UniqueId            string `gorm:"column:unique_id" json:"unique_id"`
    CreatedAt           time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedAt           time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (BasicConfiguration) TableName() string {
	return "configurations.basic_configuration"
}
