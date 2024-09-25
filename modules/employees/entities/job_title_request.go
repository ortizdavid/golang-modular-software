package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// -- Create
type CreateJobTitleRequest struct {
	TitleName   string `json:"title_name" form:"title_name" validate:"required,max=150"`
	Description string `json:"description" form:"description" validate:"required,max=255"`
}

func (req CreateJobTitleRequest) Validate() error {
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
type UpdateJobTitleRequest struct {
	TitleName   string `json:"title_name" form:"title_name" validate:"required,max=150"`
	Description string `json:"description" form:"description" validate:"required,max=255"`
}

func (req UpdateJobTitleRequest) Validate() error {
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

type SearchJobTitleRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}