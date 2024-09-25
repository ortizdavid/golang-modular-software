package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// --- CREATE
type CreateStatusRequest struct {
	StatusName  string `json:"status_name" form:"status_name" validate:"required,max=100"`
	Code        string `json:"code" form:"code" validate:"required,max=30"`
	Weight      int `json:"weight" form:"weight"`
	LblColor    string `json:"lbl_color" form:"lbl_color" validate:"max=20"`
	BgColor     string `json:"bg_color" form:"bg_color" validate:"max=20"`
	Description string `json:"description" form:"description" validate:"max=255"`
}

func (req CreateStatusRequest) Validate() error {
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

// --- UPDATE 
type UpdateStatusRequest struct {
	StatusName  string `json:"status_name" form:"status_name" validate:"required,max=100"`
	Code        string `json:"code" form:"code" validate:"required,max=30"`
	Weight      int `json:"weight" form:"weight"`
	LblColor    string `json:"lbl_color" form:"lbl_color" validate:"max=20"`
	BgColor     string `json:"bg_color" form:"bg_color" validate:"max=20"`
	Description string `json:"description" form:"description" validate:"max=255"`
}

func (req UpdateStatusRequest) Validate() error {
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
type SearchStatusRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}