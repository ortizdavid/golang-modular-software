package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Address struct {
	AddressId         int64  `gorm:"autoIncrement;primaryKey"`
	EmployeeId        int64  `gorm:"column:employee_id"`
	State             string `gorm:"column:state"`
	City              string `gorm:"column:city"`
	Neighborhood      string `gorm:"column:neighborhood"`
	Street            string `gorm:"column:street"`
	HouseNumber       string `gorm:"house_number"`
	PostalCode        string `gorm:"postal_code"`
	CountryCode       string `gorm:"column:country_code"`
	AdditionalDetails string `gorm:"column:additional_details"`
	IsCurrent         bool   `gorm:"column:is_current"`
	shared.BaseEntity
}

func (Address) TableName() string {
	return "employees.address"
}

type AddressData struct {
	AddressId         int64  `json:"address_id"`
	UniqueId          string `json:"unique_id"`
	State             string `json:"state"`
	City              string `json:"city"`
	Neighborhood      string `json:"neighborhood"`
	Street            string `json:"street"`
	HouseNumber       string `json:"house_number"`
	PostalCode        string `json:"postal_code"`
	CountryCode       string `json:"country_code"`
	AdditionalDetails string `json:"additional_details"`
	IsCurrent         bool   `json:"is_current"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	EmployeeId        int64  `json:"employee_id"`
	EmployeeUniqueId  string `json:"employee_unique_id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
}
