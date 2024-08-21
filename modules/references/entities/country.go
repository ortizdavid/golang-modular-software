package entities

import "time"

type Country struct {
	CountryId int `gorm:"autoIncrement;primaryKey"`
	CountryName  string    `gorm:"column:country_name"`
	IsoCode  string    `gorm:"column:iso_code"`
	DialingCode  string    `gorm:"column:dialing_code"`
	UniqueId  string    `gorm:"column:unique_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Country) TableName() string {
	return "reference.countries"
}

type CountryData struct {
	CountryId 		int `json:"country_id"`
	UniqueId  		string    `json:"unique_id"`
	CountryName  	string    `json:"country_name"`
	IsoCode  		string    `json:"iso_code"`
	IsoCodeLower  	string    `json:"iso_code_lower"`
	DialingCode  	string    `json:"dialing_code"`
	CreatedAt 		string `json:"created_at"`
	UpdatedAt 		string `json:"updated_at"`
}