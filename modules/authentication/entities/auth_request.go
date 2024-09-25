package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// -- Login
type LoginRequest struct {
	UserName 	string `json:"user_name" form:"user_name"` // can be UserName, Email or Other
	Password	string `json:"password" form:"password"`
}	

func (req LoginRequest) Validate() error {
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

// --RecoverPassword
type RecoverPasswordRequest struct {
	Password		string `json:"password" form:"password"`
	PasswordConf	string `json:"password_conf" form:"password_conf"`
}

func (req RecoverPasswordRequest) Validate() error {
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

// GetRecoverLink
type GetRecoverLinkRequest struct {
	Email		string `json:"email" form:"email"`
}

func (req GetRecoverLinkRequest) Validate() error {
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

// RefreshToken
type RefreshTokenRequest struct {
	RefreshToken	string `json:"refresh_token"`
}

func (req RefreshTokenRequest) Validate() error {
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
