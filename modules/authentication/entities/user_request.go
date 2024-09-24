package entities

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// -- Change Password
type CreateUserRequest struct {
	UserName		string `json:"user_name" form:"user_name" validate:"required,min=5,max=100"`
	RoleId			int `json:"role_id" form:"role_id" validate:"required"`
	Email			string `json:"email" form:"email" validate:"required,email"`
	Password		string `json:"password" form:"password" validate:"required,min=6,max=255"`
}

func (req CreateUserRequest) Validate() error {
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

// -- Change Password
type UpdatePasswordRequest struct {
	NewPassword		string `json:"new_password" form:"new_password" validate:"required,min=6,max=255"`
	PasswordConf	string `json:"password_conf" form:"password_conf" validate:"required,min=6,max=255"`
}

func (req UpdatePasswordRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidatorFormatErrors(errs)
		}
		return err
	}
	if req.NewPassword != req.PasswordConf {
		return errors.New("password and confirmation do not match")
	}
	return nil
}


// -- Change Profile Image
type ChangeImageRequest struct {
	UserImage		string `json:"user_image" form:"user_image"`
}

func (req ChangeImageRequest) Validate() error {
	return nil
}

// -- Add User Role
type AssignUserRoleRequest struct {
	RoleId	int `json:"role_id" form:"role_id" validate:"required"`
}

func (req AssignUserRoleRequest) Validate() error {
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

type AssociateUserRequest struct {
	UserId		int64 `json:"user_id" form:"user_id" validate:"required"`
	EntityId	int64 `json:"entity_id" form:"entity_id" validate:"required"`
	EntityName	string `json:"entity_name" form:"entity_name" validate:"required"`
}

func (req AssociateUserRequest) Validate() error {
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

// 
type UpdateUserRequest struct {
	UserName	string `json:"user_name" form:"user_name" validate:"required,min=5,max=100"`
}

func (req UpdateUserRequest) Validate() error {
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

type SearchUserRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}

