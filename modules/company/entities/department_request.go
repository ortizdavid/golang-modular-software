package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// ---- Create
type CreateDepartmentRequest struct {
	CompanyId      int `json:"company_id" form:"company_id" validate:"required"`
	DepartmentName string `json:"department_name" form:"department_name" validate:"required,max=150"`
	Acronym        string `json:"acronym" form:"acronym" validate:"max=20"`
	Description    string `json:"description" form:"description" validate:"max=255"`
}

func (req CreateDepartmentRequest) Validate() error  {
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

// ---- Update
type UpdateDepartmentRequest struct {
	CompanyId      int `json:"company_id" form:"company_id" validate:"required"`
	DepartmentName string `json:"department_name" form:"department_name" validate:"required,max=150"`
	Acronym        string `json:"acronym" form:"acronym" validate:"max=20"`
	Description    string `json:"description" form:"description" validate:"max=255"`
}

func (req UpdateDepartmentRequest) Validate() error  {
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


// --- Search
type SearchDepartmentRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
