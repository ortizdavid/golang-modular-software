package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// --------- CREATE-------------------------
type CreateRoleRequest struct {
	RoleName    string `json:"role_name" form:"role_name" validate:"required,min=3,max=100"`
	Code        string `json:"code" form:"code" validate:"required,startswith=role_,min=5,max=50"`
	Description string `json:"description" form:"description" validate:"max=255"`
	Status      string `json:"status" form:"status" validate:"required,oneof=Enabled Disabled"`
}

func (req CreateRoleRequest) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("startswith", helpers.ValidatorStartsWith)
	err := validate.Struct(req)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidatorFormatErrors(errs)
		}
		return err
	}
	return nil
}
// --------- UPDATE---------------------------
type UpdateRoleRequest struct {
	RoleName    string `json:"role_name" form:"role_name" validate:"required,min=3,max=100"`
	Code        string `json:"code" form:"code" validate:"required,startswith=role_,min=5,max=50"`
	Description string `json:"description" form:"description" validate:"max=255"`
	Status      string `json:"status" form:"status" validate:"required,oneof=Enabled Disabled"`
}

func (req UpdateRoleRequest) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("startswith", helpers.ValidatorStartsWith)
	err := validate.Struct(req)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidatorFormatErrors(errs)
		}
		return err
	}
	return nil
}

type SearchRoleRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
