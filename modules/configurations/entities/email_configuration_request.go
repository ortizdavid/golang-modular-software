package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

type UpdateEmailConfigurationRequest struct {
	SMTPServer     string `json:"smtp_server" form:"smtp_server" validate:"required,min=8,max=50"`
	SMTPPort       string `json:"smtp_port" form:"smtp_port" validate:"required,max=5"`
	SenderEmail    string `json:"sender_email" form:"sender_email" validate:"required,min=8,max=100"`
	SenderPassword string `json:"sender_password" form:"sender_password" validate:"required,min=8,max=150"`
}

func (req UpdateEmailConfigurationRequest) Validate() error {
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
