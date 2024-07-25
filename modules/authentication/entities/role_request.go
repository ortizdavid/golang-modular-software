package entities

// --------- CREATE-------------------------
type CreateRoleRequest struct {
	Code string `json:"code"`
	RoleName string `json:"role_name"`
	Description string `json:"description"`
}

func (req CreateRoleRequest) Validate() error {
	return nil
}

// --------- UPDATE---------------------------
type UpdateRoleRequest struct {
	RoleName string `json:"role_name"`
	Description string `json:"description"`
}

func (req UpdateRoleRequest) Validate() error {
	return nil
}