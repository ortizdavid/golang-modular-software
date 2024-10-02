package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// - Create
type CreateAddressRequest struct {
	EmployeeId        int64  `json:"employee_id" form:"employee_id" validate:"required"`
	State             string `json:"state" form:"state" validate:"max=100"`
	City              string `json:"city" form:"city" validate:"max=100"`
	Neighborhood      string `json:"neighborhood" form:"neighborhood" validate:"max=200"`
	Street            string `json:"street" form:"street" validate:"max=200"`
	HouseNumber       string `json:"house_number" form:"house_number" validate:"max=20"`
	PostalCode        string `json:"portal_code" form:"postal_code" validate:"max=20"`
	CountryCode       string `json:"country_code" form:"country_code" validate:"max=3"`
	AdditionalDetails string `json:"additional_details" form:"additional_details" validate:"max=255"`
}

func (req CreateAddressRequest) Validate() error {
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

// - Update
type UpdateAddressRequest struct {
	EmployeeId        int64  `json:"employee_id" form:"employee_id" validate:"required"`
	State             string `json:"state" form:"state" validate:"max=100"`
	City              string `json:"city" form:"city" validate:"max=100"`
	Neighborhood      string `json:"neighborhood" form:"neighborhood" validate:"max=200"`
	Street            string `json:"street" form:"street" validate:"max=200"`
	HouseNumber       string `json:"house_number" form:"house_number" validate:"max=20"`
	PostalCode        string `json:"portal_code" form:"postal_code" validate:"max=20"`
	CountryCode       string `json:"country_code" form:"country_code" validate:"max=3"`
	AdditionalDetails string `json:"additional_details" form:"additional_details" validate:"max=255"`
}

func (req UpdateAddressRequest) Validate() error {
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
