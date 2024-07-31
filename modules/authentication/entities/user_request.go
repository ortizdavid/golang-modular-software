package entities

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
type ChangePasswordRequest struct {
	NewPassword		string `json:"new_password" form:"new_password"`
	PasswordConf	string `json:"password_conf" form:"password_conf"`
}

func (req ChangePasswordRequest) Validate() error {
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

