package entities

// -- Change Password
type CreateUserRequest struct {
	UserName		string `json:"user_name"`
	RoleId			int `json:"role_id"`
	Email			string `json:"email"`
	Password		string `json:"password"`
}

func (req CreateUserRequest) Validate() error {
	return nil
}

// -- Change Password
type ChangePasswordRequest struct {
	NewPassword		string `json:"new_password"`
	ConfPassword	string `json:"conf_password"`
}

func (req ChangePasswordRequest) Validate() error {
	return nil
}

// -- Change Profile Image
type ChangeImageRequest struct {
	UserImage		string `json:"user_image"`
}

func (req ChangeImageRequest) Validate() error {
	return nil
}

// -- Add User Role
type AssignUserRoleRequest struct {
	RoleId	int `json:"role_id"`
}

func (req AssignUserRoleRequest) Validate() error {
	return nil
}

// -- Login
type LoginRequest struct {
	UserName	string `json:"user_name"`
	Password	string `json:"password"`
}	

func (req LoginRequest) Validate() error {
	return nil
}
