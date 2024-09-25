package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// -- Create
type CreateCountryRequest struct {
	CountryName 	string `json:"country_name" form:"country_name" validate:"required,max=150"`
	IsoCode 		string `json:"iso_code" form:"iso_code" validate:"required,max=3"`
	DialingCode 	string `json:"dialing_code" form:"dialing_code" validate:"required,max=7"`
}

func (req CreateCountryRequest) Validate() error {
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

// -- Update
type UpdateCountryRequest struct {
	CountryName 	string `json:"country_name" form:"country_name" validate:"required,max=150"`
	IsoCode 		string `json:"iso_code" form:"iso_code" validate:"required,max=3"`
	DialingCode 	string `json:"dialing_code" form:"dialing_code" validate:"required,max=7"`
}

func (req UpdateCountryRequest) Validate() error {
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


// search
type SearchCountryRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}