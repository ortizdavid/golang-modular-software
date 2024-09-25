package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

type UpdateBasicConfigurationRequest struct {
	AppName            string `json:"app_name" form:"app_name"`
	AppAcronym         string `json:"app_acronym" form:"app_acronym"`
	MaxAdmninUsers     int    `json:"max_admin_users" form:"max_admin_users"`
	MaxSuperAdminUsers int    `json:"max_super_admin_users" form:"max_super_admin_users"`
	MaxRecordPerPage   int    `json:"max_record_per_page" form:"max_record_per_page"`
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
