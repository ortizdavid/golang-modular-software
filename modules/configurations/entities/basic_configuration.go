package entities

type BasicConfiguration struct {
	ConfigurationId   		int `gorm:"column:configuration_id"`
	MaxRecordsPerPage   	int `gorm:"column:max_records_per_page"`
	MaxAdmninUsers		   	int `gorm:"column:max_admin_users"`
	MaxSuperAdmninUsers		int `gorm:"column:max_super_admin_users"`
}

func (BasicConfiguration) TableName() string {
	return "configurations.basic_configuration"
}
