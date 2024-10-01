package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// ---- Create
type CreateProjectAttachmentRequest struct {
	ProjectId    int `json:"project_id" form:"project_id" validate:"required"`
	CompanyId    int `json:"company_id" form:"company_id" validate:"required"`
    AttachmentName   string `json:"attachment_name" form:"attachment_name" validate:"required,max=100"`
}

func (req CreateProjectAttachmentRequest) Validate() error  {
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

// ---- Update
type UpdateProjectAttachmentRequest struct {
	ProjectId    int `json:"project_id" form:"project_id" validate:"required"`
	CompanyId    int `json:"company_id" form:"company_id" validate:"required"`
    AttachmentName   string `json:"attachment_name" form:"attachment_name" validate:"required,max=100"`
}

func (req UpdateProjectAttachmentRequest) Validate() error  {
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