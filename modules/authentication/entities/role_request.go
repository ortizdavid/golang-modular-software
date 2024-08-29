package entities

// --------- CREATE-------------------------
type CreateRoleRequest struct {
	RoleName string `json:"role_name" form:"role_name"`
	Code string `json:"code" form:"code"`
	Description string `json:"description" form:"description"`
	Status string `json:"status" form:"status"`
}

func (req CreateRoleRequest) Validate() error {
	return nil
}

// --------- UPDATE---------------------------
type UpdateRoleRequest struct {
	RoleName string `json:"role_name" form:"role_name"`
	Code string `json:"code" form:"code"`
	Description string `json:"description" form:"description"`
	Status string `json:"status" form:"status"`
}

func (req UpdateRoleRequest) Validate() error {
	return nil
}

type SearchRoleRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
