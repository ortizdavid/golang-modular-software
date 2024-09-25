package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// --- CREATE
type CreateTypeRequest struct {
	TypeName  string `json:"type_name" form:"type_name"`
	Code        string `json:"code" form:"code"`
}

func (req CreateTypeRequest) Validate() error {
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

// --- UPDATE 
type UpdateTypeRequest struct {
	TypeName  string `json:"type_name" form:"type_name"`
	Code        string `json:"code" form:"code"`
}

func (req UpdateTypeRequest) Validate() error {
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


// search
type SearchTypeRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}