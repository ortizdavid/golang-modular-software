package entities

// -- Login
type LoginRequest struct {
	UserName	string `json:"user_name"`
	Password	string `json:"password"`
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

