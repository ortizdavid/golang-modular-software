package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// --CREATE
type CreateCoreEntityRequest struct {
	ModuleId    int    `json:"module_id" form:"module_id" validate:"required"`
	Code        string `json:"code" form:"code" validate:"required,min=4,max=50"`
	EntityName  string `json:"entity_name" form:"entity_name" validate:"required,min=5,max=100"`
	Description string `json:"description" form:"description" validate:"required,max=255"`
}

func (req CreateCoreEntityRequest) Validate() error {
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

// --CREATE
type UpdateCoreEntityRequest struct {
	ModuleId    int    `json:"module_id" form:"module_id" validate:"required"`
	Code        string `json:"code" form:"code" validate:"required,min=4,max=50"`
	EntityName  string `json:"entity_name" form:"entity_name" validate:"required,min=5,max=100"`
	Description string `json:"description" form:"description" validate:"required,max=255"`
}

func (req UpdateCoreEntityRequest) Validate() error {
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
type SearchCoreEntityRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}
