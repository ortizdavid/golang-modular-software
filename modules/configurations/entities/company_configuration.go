package entities

import "time"

type CompanyConfiguration struct {
    ConfigurationId  int `gorm:"autoIncrement;primaryKey" json:"configuration_id"`
    CompanyName      string `gorm:"column:company_name" json:"company_name"`
    CompanyAcronym   string `gorm:"column:company_acronym" json:"company_acronym"`
    CompanyMainColor string `gorm:"column:company_main_color" json:"company_main_color"`
    CompanyLogo      string `gorm:"column:company_logo" json:"company_logo"`
    CompanyPhone     string `gorm:"column:company_phone" json:"company_phone"`
    CompanyEmail     string `gorm:"column:company_email" json:"company_email"`
    UniqueId         string `gorm:"column:unique_id" json:"unique_id"`
    CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (CompanyConfiguration) TableName() string {
	return "configurations.company_configuration"
}
