package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// ------ Create --
// CreateCompanyRequest represents the request structure for creating a new company.
type CreateCompanyRequest struct {
	CompanyName    string    `json:"company_name" form:"company_name"`
    CompanyAcronym string    `json:"company_acronym" form:"company_acronym"`
    CompanyType    string    `json:"company_type" form:"company_type"`
    Industry       string    `json:"industry" form:"industry"`
    FoundedDate    string `json:"founded_date" form:"founded_date" `
    Address        string    `json:"address" form:"address"`
    Phone          string    `json:"phone" form:"phone"`
    Email          string    `json:"email" form:"email"`
    WebsiteURL     string    `json:"website_url" form:"website_url"`
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
// UpdateCompanyRequest represents the request structure for updating an existing company.
type UpdateCompanyRequest struct {
	CompanyName    string    `json:"company_name" form:"company_name"`
    CompanyAcronym string    `json:"company_acronym" form:"company_acronym"`
    CompanyType    string    `json:"company_type" form:"company_type"`
    Industry       string    `json:"industry" form:"industry"`
    FoundedDate    string `json:"founded_date" form:"founded_date"`
    Address        string    `json:"address" form:"address"`
    Phone          string    `json:"phone" form:"phone"`
    Email          string    `json:"email" form:"email"`
    WebsiteURL     string    `json:"website_url" form:"website_url"`
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
