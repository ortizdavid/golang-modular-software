package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

type UpdateCompanyConfigurationRequest struct {
	CompanyName      string `json:"company_name" form:"company_name"`
	CompanyAcronym   string `json:"company_acronym" form:"company_acronym"`
	CompanyPhone     string `json:"company_phone" form:"company_phone"`
	CompanyEmail     string `json:"company_email" form:"company_email"`
	CompanyMainColor string `json:"company_main_color" form:"company_main_color"`
}

func (req UpdateCompanyConfigurationRequest) Validate() error {
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
