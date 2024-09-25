package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// ------ Create --
type CreateCompanyRequest struct {
	CompanyName    string    `json:"company_name" form:"company_name" validate:"required,max=100"`
    CompanyAcronym string    `json:"company_acronym" form:"company_acronym" validate:"max=20"`
    CompanyType    string    `json:"company_type" form:"company_type" validate:"max=50"`
    Industry       string    `json:"industry" form:"industry" validate:"max=50"`
    FoundedDate    string `json:"founded_date" form:"founded_date"`
    Address        string    `json:"address" form:"address" validate:"max=255"`
    Phone          string    `json:"phone" form:"phone" validate:"max=20"`
    Email          string    `json:"email" form:"email" validate:"max=100"`
    WebsiteURL     string    `json:"website_url" form:"website_url" validate:"max=100"`
}

func (req CreateCompanyRequest) Validate() error {
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

// ------ Update --
type UpdateCompanyRequest struct {
	CompanyName    string    `json:"company_name" form:"company_name" validate:"required,max=100"`
    CompanyAcronym string    `json:"company_acronym" form:"company_acronym" validate:"max=20"`
    CompanyType    string    `json:"company_type" form:"company_type" validate:"max=50"`
    Industry       string    `json:"industry" form:"industry" validate:"max=50"`
    FoundedDate    string `json:"founded_date" form:"founded_date"`
    Address        string    `json:"address" form:"address" validate:"max=255"`
    Phone          string    `json:"phone" form:"phone" validate:"max=20"`
    Email          string    `json:"email" form:"email" validate:"max=100"`
    WebsiteURL     string    `json:"website_url" form:"website_url" validate:"max=100"`
}

func (req UpdateCompanyRequest) Validate() error {
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
type SearchCompanyRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
