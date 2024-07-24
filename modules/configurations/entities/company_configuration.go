package entities

type CompanyConfiguration struct {
	ConfigurationId   	int `gorm:"column:configuration_id"`
	CompanyName   		string `gorm:"column:company_name"`
	CompanyAcronym   	string `gorm:"column:company_acronym"`
	CompanyMainColor   	string `gorm:"column:company_main_color"`
	CompanyLogo   		string `gorm:"column:company_logo"`
	CompanyPhone   		string `gorm:"column:company_phone"`
	CompanyEmail   		string `gorm:"column:company_email"`
}

func (CompanyConfiguration) TableName() string {
	return "configurations.company_configuration"
}
