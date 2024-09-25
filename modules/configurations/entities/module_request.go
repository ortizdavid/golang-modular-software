package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

//--CREATE
type CreateModuleRequest struct {
    ModuleName  string    `json:":module_name" form:"module_name"`
    Code        string    `json:":code" form:"code"`
    Description string    `json:":description" form:"description"`
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
    ModuleName  string    `json:":module_name" form:"module_name"`
    Code        string    `json:":code" form:"code"`
    Description string    `json:":description" form:"description"`
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
