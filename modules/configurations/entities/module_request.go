package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

//--CREATE
type CreateModuleRequest struct {
    ModuleName  string `json:":module_name" form:"module_name" validate:"required,min=5,max=100"`
    Code        string `json:":code" form:"code" validate:"required,min=4,max=30"`
    Description string `json:":description" form:"description" validate:"required,max=255"`
}

func (req CreateModuleRequest) Validate() error {
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


//--CREATE
type UpdateModuleRequest struct {
    ModuleName  string `json:":module_name" form:"module_name" validate:"required,min=5,max=100"`
    Code        string `json:":code" form:"code" validate:"required,min=4,max=30"`
    Description string `json:":description" form:"description" validate:"required,max=255"`
}

func (req UpdateModuleRequest) Validate() error {
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
type SearchModuleRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
