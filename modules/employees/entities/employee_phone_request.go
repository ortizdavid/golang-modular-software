package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// -- Create
type CreateEmployeeEmailRequest struct {
	EmployeeId		int64 `json:"employee_id" form:"employee_id"`
	ContactTypeId	int	`json:"contact_type_id" form:"contact_type_id"`
	EmailAddress   	string `json:"email_address" form:"email_address"`
}

func (req CreateEmployeeEmailRequest) Validate() error {
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

// -- Updaye 
type UpdateEmployeeEmailRequest struct {
	EmployeeId		int64 `json:"employee_id" form:"employee_id"`
	ContactTypeId	int	`json:"contact_type_id" form:"contact_type_id"`
	EmailAddress   	string `json:"email_address" form:"email_address"`
}

func (req UpdateEmployeeEmailRequest) Validate() error {
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

type SearchEmployeeEmailRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}