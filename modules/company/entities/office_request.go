package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// ---- Create
type CreateOfficeRequest struct {
	CompanyId    int `json:"company_id" form:"company_id"`
    OfficeName   string `json:"office_name" form:"office_name"`
    Code         string `json:"code" form:"code"`
    Address      string `json:"address" form:"address"`
    Phone        string `json:"phone" form:"phone"`
    Email        string `json:"email" form:"email"`
}

func (req CreateOfficeRequest) Validate() error  {
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
type UpdateOfficeRequest struct {
	CompanyId    int `json:"company_id" form:"company_id"`
    Code         string `json:"code" form:"code"`
    OfficeName   string `json:"office_name" form:"office_name"`
    Address      string `json:"address" form:"address"`
    Phone        string `json:"phone" form:"phone"`
    Email        string `json:"email" form:"email"`
}

func (req UpdateOfficeRequest) Validate() error  {
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
type SearchOfficeRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
