package entities

// --------- CREATE-------------------------
type CreatePermissionRequest struct {
	Code string `json:"code" form:"code"`
	PermissionName string `json:"permission_name" form:"permission_name"`
	Description string `json:"description" form:"description"`
}

func (req CreatePermissionRequest) Validate() error {
	return nil
}

// --------- UPDATE---------------------------
type UpdatePermissionRequest struct {
	PermissionName string `json:"permission_name" form:"permission_name"`
	Description string `json:"description" form:"permission_name"`
}

func (req UpdatePermissionRequest) Validate() error {
	return nil
}