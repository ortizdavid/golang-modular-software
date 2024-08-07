package entities

// --------- CREATE-------------------------
type CreatePermissionRequest struct {
	Code string `json:"code"`
	PermissionName string `json:"permission_name"`
	Description string `json:"description"`
}

func (req CreatePermissionRequest) Validate() error {
	return nil
}

// --------- UPDATE---------------------------
type UpdatePermissionRequest struct {
	PermissionName string `json:"permission_name"`
	Description string `json:"description"`
}

func (req UpdatePermissionRequest) Validate() error {
	return nil
}