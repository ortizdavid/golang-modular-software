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

// -- Add Role Permission
type AssignRolePermissionRequest struct {
	PermissionId	int `json:"permission_id" form:"permission_id"`
}

func (req AssignRolePermissionRequest) Validate() error {
	return nil
}


type SearchPermissionRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
