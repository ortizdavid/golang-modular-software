package entities

// -- Login
type LoginRequest struct {
	UserName 	string `json:"user_name" form:"user_name"` // can be UserName, Email or Other
	Password	string `json:"password" form:"password"`
}	

func (req LoginRequest) Validate() error {
	return nil
}

// --RecoverPassword
type RecoverPasswordRequest struct {
	Password		string `json:"password"`
	PasswordConfiration	string `json:"password_confirmation"`
}

func (req RecoverPasswordRequest) Validate() error {
	return nil
}

// GetRecoverLink
type GetRecoverLinkRequest struct {
	Email		string `json:"email"`
}

func (req GetRecoverLinkRequest) Validate() error {
	return nil
}

