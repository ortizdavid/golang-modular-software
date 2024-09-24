package entities

import "errors"

// -- Change Password
type CreateUserRequest struct {
	UserName		string `json:"user_name" form:"user_name"`
	RoleId			int `json:"role_id" form:"role_id"`
	Email			string `json:"email" form:"email"`
	Password		string `json:"password" form:"password"`
}

func (req CreateUserRequest) Validate() error {
	return nil
}

// -- Change Password
type UpdatePasswordRequest struct {
	NewPassword		string `json:"new_password" form:"new_password"`
	PasswordConf	string `json:"password_conf" form:"password_conf"`
}

func (req UpdatePasswordRequest) Validate() error {
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
	RoleId	int `json:"role_id" form:"role_id"`
}

func (req AssignUserRoleRequest) Validate() error {
	return nil
}

type AssociateUserRequest struct {
	UserId		int64 `json:"user_id" form:"user_id"`
	EntityId	int64 `json:"entity_id" form:"entity_id"`
	EntityName	string `json:"entity_name" form:"entity_name"`
}

func (req AssociateUserRequest) Validate() error {
	return nil
}

// 
type UpdateUserRequest struct {
	UserName	string `json:"user_name" form:"user_name"`
}

func (req UpdateUserRequest) Validate() error {
	return nil
}

type SearchUserRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}

