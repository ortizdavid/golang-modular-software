package entities

// --------- CREATE-------------------------
type CreateRoleRequest struct {
	RoleName string `json:"role_name" form:"role_name"`
	Code string `json:"code" form:"code"`
	Description string `json:"description" form:"description"`
}

func (req CreateRoleRequest) Validate() error {
	return nil
}

// --------- UPDATE---------------------------
type UpdateRoleRequest struct {
	RoleName string `json:"role_name" form:"role_name"`
	Code string `json:"code" form:"code"`
	Description string `json:"description" form:"description"`
}

func (req UpdateRoleRequest) Validate() error {
	return nil
}