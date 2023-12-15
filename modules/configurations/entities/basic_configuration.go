package entities


type BasicConfiguration struct {
	ConfigurationId   	int `gorm:"column:ConfigurationId;primaryKey"`
	CompanyName   		string `gorm:"column:CompanyName;type:varchar(50)"`
	CompanyAcronym   	string `gorm:"column:CompanyAcronym;type:varchar(20)"`
	NumOfRecordsPerPage  int `gorm:"column:NumOfRecordsPerPage;type:int"`
	CompanyMainColor   	string `gorm:"column:CompanyMainColor;type:varchar(15)"`
}


func (BasicConfiguration) TableName() string {
	return "BasicConfiguration"
}