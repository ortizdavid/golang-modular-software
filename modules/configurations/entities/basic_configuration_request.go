package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

type UpdateBasicConfigurationRequest struct {
	AppName            string `json:"app_name" form:"app_name" validate:"required,min=5,max=100"`
	AppAcronym         string `json:"app_acronym" form:"app_acronym" validate:"required,min=3,max=50"`
	MaxAdmninUsers     int    `json:"max_admin_users" form:"max_admin_users" validate:"required,max=20"`
	MaxSuperAdminUsers int    `json:"max_super_admin_users" form:"max_super_admin_users" validate:"required,max=10"`
	MaxRecordPerPage   int    `json:"max_record_per_page" form:"max_record_per_page" validate:"required,min=10,max=100"`
}

func (req UpdateBasicConfigurationRequest) Validate() error {
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
