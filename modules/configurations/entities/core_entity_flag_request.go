package entities

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

type ManageCoreEntityFlagRequest struct {
	FlagId int    `json:"flag_id" form:"flag_id" validate:"required"`
	Status string `json:"statuses" form:"status" validate:"required,oneof=Enabled Disabled"`
}

func (req ManageCoreEntityFlagRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidatorFormatErrors(errs)
		}
		return err
	}
	if req.Status != "Enabled" && req.Status != "Disabled" {
		return fmt.Errorf("invalid status: %s", req.Status)
	}
	return nil
}
