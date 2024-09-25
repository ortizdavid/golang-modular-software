package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// -- Create
type CreateDocumentTypeRequest struct {
	TypeName   string `json:"type_name" form:"type_name"`
	Description string `json:"description" form:"description"`
}

func (req CreateDocumentTypeRequest) Validate() error {
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
type UpdateDocumentTypeRequest struct {
	TypeName   string `json:"type_name" form:"type_name"`
	Description string `json:"description" form:"description"`
}

func (req UpdateDocumentTypeRequest) Validate() error {
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

type SearchDocumentTypeRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}