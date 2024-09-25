package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// -- Create
type CreateCurrencyRequest struct {
	CurrencyName string    `json:"currency_name" form:"currency_name" validate:"required,max=100"`
	Code         string    `json:"code" form:"code" validate:"required,max=3"`
	Symbol       string  `json:"symbol" form:"symbol" validate:"max=10"`
}

func (req CreateCurrencyRequest) Validate() error {
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
type UpdateCurrencyRequest struct {
	CurrencyName string    `json:"currency_name" form:"currency_name" validate:"required,max=100"`
	Code         string    `json:"code" form:"code" validate:"required,max=3"`
	Symbol       string  `json:"symbol" form:"symbol" validate:"max=10"`
}

func (req UpdateCurrencyRequest) Validate() error {
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
type SearchCurrencyRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}