package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

type UpdateCompanyConfigurationRequest struct {
	CompanyName      string `json:"company_name" form:"company_name" validate:"required,min=5,max=100"`
	CompanyAcronym   string `json:"company_acronym" form:"company_acronym" validate:"required,min=2,max=50"`
	CompanyPhone     string `json:"company_phone" form:"company_phone" validate:"required,max=20"`
	CompanyEmail     string `json:"company_email" form:"company_email" validate:"required,min=8,max=100"`
	CompanyMainColor string `json:"company_main_color" form:"company_main_color" validate:"required,max=10"`
	CompanyLogo string `json:"company_logo" form:"company_logo" validate:"required,max=100"`
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
