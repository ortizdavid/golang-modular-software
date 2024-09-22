package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type ProfessionalInfo struct {
	ProfissionalId			int64 `gorm:"autoIncrement;primaryKey;column:professional_id"`
	EmployeeId				int64 `gorm:"column:employee_id"`
	DepartmentId			int `gorm:"column:department_id"`
	JobTitleId				int `gorm:"column:job_title_id"`
	EmploymentStatusId		int `gorm:"column:employment_status_id"`
	shared.BaseEntity
}


func (ProfessionalInfo) TableName() string {
	return "employees.professional_info"
}


type ProfessionalInfoData struct {
	ProfessionalId			int64 `json:"professional_id"`
	UniqueId  				string    `json:"unique_id"`
	CreatedAt 				string `json:"created_at"`
	UpdatedAt 				string `json:"updated_at"`
	EmployeeId				int64 `json:"employee_id"`
	FirstName				string `json:"first_name"`
	LastName				string `json:"last_name"`
	DepartmentId			int `json:"department_id"`
	DepartmentName			string `json:"department_name"`
	EmploymentStatusId		int `json:"employment_status_id"`
	EmployeeUniqueId  		string    `json:"employee_unique_id"`
	EmploymentStatusName	string `json:"employment_status_name"`
	JobTitleId				int `json:"job_title_id"`
	JobTitleName			string `json:"job_title_name"`
}
