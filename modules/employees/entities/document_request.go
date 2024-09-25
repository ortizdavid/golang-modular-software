package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// -- Create
type CreateDocumentRequest struct {
	EmployeeId		int64 `json:"employee_id" form:"employee_id" validate:"required"`
	DocumentTypeId	int	`json:"document_type_id" form:"document_type_id" validate:"required"`
	DocumentName   	string `json:"document_name" form:"document_name" validate:"required,max=150"`
	DocumentNumber 	string `json:"document_number" form:"document_number" validate:"required,max=40"`
	ExpirationDate	string `json:"expiration_date" form:"expiration_date"`
	Status			string `json:"status" form:"status" validate:"required,oneof=Expired Active"`
}

func (req CreateDocumentRequest) Validate() error {
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

// -- Updaye 
type UpdateDocumentRequest struct {
	EmployeeId		int64 `json:"employee_id" form:"employee_id" validate:"required"`
	DocumentTypeId	int	`json:"document_type_id" form:"document_type_id" validate:"required"`
	DocumentName   	string `json:"document_name" form:"document_name" validate:"required,max=150"`
	DocumentNumber 	string `json:"document_number" form:"document_number" validate:"required,max=40"`
	ExpirationDate	string `json:"expiration_date" form:"expiration_date"`
	Status			string `json:"status" form:"status" validate:"required,oneof=Expired Active"`
}

func (req UpdateDocumentRequest) Validate() error {
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

type SearchDocumentRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}