package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Country struct {
	CountryId int `gorm:"autoIncrement;primaryKey" json:"country_id"`
	CountryName  string    `gorm:"column:country_name" json:"country_name"`
	IsoCode  string    `gorm:"column:iso_code" json:"iso_code"`
	DialingCode  string    `gorm:"column:dialing_code" json:"dialing_code"`
	shared.BaseEntity
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