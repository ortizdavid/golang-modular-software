package entities

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/go-nopain/datetime"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// -- Create
type CreateEmployeeRequest struct {
	IdentificationTypeId	int `json:"identification_type_id" form:"identification_type_id" validate:"required"`
	CountryId				int `json:"country_id" form:"country_id"`
	MaritalStatusId			int `json:"marital_status_id" form:"marital_status_id"`
	DepartmentId			int `json:"department_id" form:"department_id"`
	JobTitleId				int `json:"job_title_id" form:"job_title_id"`
	EmploymentStatusId		int `json:"employment_status_id" form:"employment_status_id"`
	FirstName				string `json:"first_name" form:"first_name" validate:"required,max=150"`
	LastName				string `json:"last_name" form:"last_name" validate:"required,max=150"`
	IdentificationNumber	string `json:"identification_number" form:"identification_number" validate:"required,max=30"`
	Gender					string `json:"gender" form:"gender" validate:"oneof=Male Female"`
	DateOfBirth				string `json:"date_of_birth" form:"date_of_birth"`
}

func (req CreateEmployeeRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidatorFormatErrors(errs)
		}
		return err
	}
	if !datetime.IsValidDate(req.DateOfBirth) {
		return fmt.Errorf("invalid format for 'Date of birth'. must be YYYY-MM-DD")
	}
	return nil
}

// -- Update
type UpdateEmployeeRequest struct {
	IdentificationTypeId	int `json:"identification_type_id" form:"identification_type_id" validate:"required"`
	CountryId				int `json:"country_id" form:"country_id"`
	MaritalStatusId			int `json:"marital_status_id" form:"marital_status_id"`
	DepartmentId			int `json:"department_id" form:"department_id"`
	JobTitleId				int `json:"job_title_id" form:"job_title_id"`
	EmploymentStatusId		int `json:"employment_status_id" form:"employment_status_id"`
	FirstName				string `json:"first_name" form:"first_name" validate:"required,max=150"`
	LastName				string `json:"last_name" form:"last_name" validate:"required,max=150"`
	IdentificationNumber	string `json:"identification_number" form:"identification_number" validate:"required,max=30"`
	Gender					string `json:"gender" form:"gender" validate:"oneof=Male Female"`
	DateOfBirth				string `json:"date_of_birth" form:"date_of_birth"`
}

func (req UpdateEmployeeRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidatorFormatErrors(errs)
		}
		return err
	}
	return nil
}


type SearchEmployeeRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}