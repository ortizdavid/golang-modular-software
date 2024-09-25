package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// ---- Create
type CreatePolicyRequest struct {
	CompanyId      	int `json:"company_id" form:"company_id"`
    PolicyName   	string `json:"policy_name" form:"policy_name"`
    Description  	string `json:"description" form:"description"`
    EffectiveDate 	string `json:"effective_date" form:"effective_date"`
}

func (req CreatePolicyRequest) Validate() error  {
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
type UpdatePolicyRequest struct {
    CompanyId      	int `json:"company_id" form:"company_id"`
    PolicyName   	string `json:"policy_name" form:"policy_name"`
    Description  	string `json:"description" form:"description"`
    EffectiveDate 	string `json:"effective_date" form:"effective_date"`
}

func (req UpdatePolicyRequest) Validate() error  {
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
type SearchPolicyRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
