package entities


type BasicConfiguration struct {
	ConfigurationId   	int `gorm:"column:configuration_id"`
	CompanyName   		string `gorm:"column:company_name"`
	CompanyAcronym   	string `gorm:"column:company_acronym"`
	NumOfRecordsPerPage  int `gorm:"column:num_of_record_per_page"`
	CompanyMainColor   	string `gorm:"column:company_namin_color"`
}


func (BasicConfiguration) TableName() string {
	return "configurations.basic_configuration"
}