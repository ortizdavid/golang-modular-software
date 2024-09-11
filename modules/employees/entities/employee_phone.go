package entities

import "time"

type EmployeePhone struct {
	PhoneId			int64 `gorm:"autoIncrement;primaryKey"`
	EmployeeId		int64 `gorm:"column:employee_id"`
	ContactTypeId	int	`gorm:"column:contact_type_id"`
	PhoneNumber   	string `gorm:"column:phone_number"`
	UniqueId  		string    `gorm:"column:unique_id"`
	CreatedAt 		time.Time `gorm:"column:created_at"`
	UpdatedAt 		time.Time `gorm:"column:updated_at"`
}

func (EmployeePhone) TableName() string {
	return "employees.employee_phones"
}

type EmployeePhoneData struct {
	PhoneId				int64 `json:"phone_id"`
	PhoneName   		string `json:"phone_name"`
	PhoneNumber 		string `json:"phone_number"`
	UniqueId    		string `json:"unique_id"`
	CreatedAt   		string `json:"created_at"`
	UpdatedAt			string `json:"updated_at"`
	EmployeeId			int64 `json:"employee_id"`
	EmployeeUniqueId	string `json:"employee_unique_id"`
	FirstName			string `json:"first_name"`
	LastName			string `json:"last_name"`
	ContactTypeId		int	`json:"contact_type_id"`
	ContactTypeName		string `json:"contact_type_name"`
}