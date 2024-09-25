package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// ---- Create
type CreateBranchRequest struct {
	CompanyId    int `json:"company_id" form:"company_id"`
    BranchName   string `json:"branch_name" form:"branch_name"`
    Code         string `json:"code" form:"code"`
    Address      string `json:"address" form:"address"`
    Phone        string `json:"phone" form:"phone"`
    Email        string `json:"email" form:"email"`
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
	CompanyId    int `json:"company_id" form:"company_id"`
    Code         string `json:"code" form:"code"`
    BranchName   string `json:"branch_name" form:"branch_name"`
    Address      string `json:"address" form:"address"`
    Phone        string `json:"phone" form:"phone"`
    Email        string `json:"email" form:"email"`
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
