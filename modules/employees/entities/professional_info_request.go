package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// -- Create
type CreateProfessionalInfoRequest struct {
	EmployeeId				int64 `json:"employee_id" form:"employee_id"`
	MaritalStatusId			int `json:"marital_status_id" form:"marital_status_id"`
	DepartmentId			int `json:"department_id" form:"department_id"`
	JobTitleId				int `json:"job_title_id" form:"job_title_id"`
	EmploymentStatusId		int `json:"employment_status_id" form:"employment_status_id"`
}

func (req CreateProfessionalInfoRequest) Validate() error {
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

// -- Update
type UpdateProfessionalInfoRequest struct {
	EmployeeId				int64 `json:"employee_id" form:"employee_id"`
	MaritalStatusId			int `json:"marital_status_id" form:"marital_status_id"`
	DepartmentId			int `json:"department_id" form:"department_id"`
	JobTitleId				int `json:"job_title_id" form:"job_title_id"`
	EmploymentStatusId		int `json:"employment_status_id" form:"employment_status_id"`
}

func (req UpdateProfessionalInfoRequest) Validate() error {
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

type SearchProfessionalInfoRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}