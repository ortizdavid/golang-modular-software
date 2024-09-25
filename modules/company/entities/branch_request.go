package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// ---- Create
type CreateBranchRequest struct {
	CompanyId    int `json:"company_id" form:"company_id" validate:"required"`
    BranchName   string `json:"branch_name" form:"branch_name" validate:"required,max=100"`
    Code         string `json:"code" form:"code" validate:"max=20"`
    Address      string `json:"address" form:"address" validate:"max=255"`
    Phone        string `json:"phone" form:"phone" validate:"max=20"`
    Email        string `json:"email" form:"email" validate:"max=100"`
}

func (req CreateBranchRequest) Validate() error  {
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
type UpdateBranchRequest struct {
	CompanyId    int `json:"company_id" form:"company_id" validate:"required"`
    BranchName   string `json:"branch_name" form:"branch_name" validate:"required,max=100"`
    Code         string `json:"code" form:"code" validate:"max=20"`
    Address      string `json:"address" form:"address" validate:"max=255"`
    Phone        string `json:"phone" form:"phone" validate:"max=20"`
    Email        string `json:"email" form:"email" validate:"max=100"`
}

func (req UpdateBranchRequest) Validate() error  {
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
type SearchBranchRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
