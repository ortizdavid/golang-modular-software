package entities

import (
	"time"

	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Employee struct {
	EmployeeId           int64     `gorm:"autoIncrement;primaryKey"`
	UserId               int64     `gorm:"column:user_id"`
	IdentificationTypeId int       `gorm:"column:identification_type_id"`
	CountryId            int       `gorm:"column:country_id"`
	MaritalStatusId      int       `gorm:"column:marital_status_id"`
	FirstName            string    `gorm:"column:first_name"`
	LastName             string    `gorm:"column:last_name"`
	IdentificationNumber string    `gorm:"column:identification_number"`
	Gender               string    `gorm:"column:gender"`
	DateOfBirth          time.Time `gorm:"column:date_of_birth"`
	shared.BaseEntity
}

func (Employee) TableName() string {
	return "employees.employees"
}

type EmployeeData struct {
	EmployeeId             int64  `json:"employee_id"`
	UniqueId               string `json:"unique_id"`
	FirstName              string `json:"first_name"`
	LastName               string `json:"last_name"`
	IdentificationNumber   string `json:"identification_number"`
	Gender                 string `json:"gender"`
	DateOfBirth            string `json:"date_of_birth"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
	IdentificationTypeId   int    `json:"identification_type_id"`
	IdentificationTypeName string `json:"identification_type_name"`
	CountryId              int    `json:"country_id"`
	CountryName            string `json:"country_name"`
	MaritalStatusId        int    `json:"marital_status_id"`
	MaritalStatusName      string `json:"marital_status_name"`
}
